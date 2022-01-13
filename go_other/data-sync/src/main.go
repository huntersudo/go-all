package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"

	"bou.ke/monkey"
	"github.com/astaxie/beego/logs"
	"github.com/spf13/viper"
)

var wg sync.WaitGroup

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/")

	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.user", "root")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.dbname", "scmbd")
	viper.SetDefault("tb.schemes", []string{"scmbd_user", "scmbd_org", "scmbd_menu", "scmbd_role", "scmbd_role_menu", "scmbd_user_role"})
	viper.SetDefault("tb.scmbd_user.SQL", "select id, aes_decrypt(name, 'des%lu#v234@3czp') as name, aes_decrypt(username, 'des%lu#v234@3czp') as username, create_time, update_time, user_status, org_id from scmbd_user")
	viper.SetDefault("tb.scmbd_org.SQL", "select id, name, parent_id, create_time, update_time from scmbd_org")
	viper.SetDefault("tb.scmbd_menu.SQL", "select id, board_id, identification, name, parent_id, leaf_type, url, url_type, create_time, update_time from scmbd_menu")
	viper.SetDefault("tb.scmbd_role.SQL", "select id, name, create_time, update_time from scmbd_role")
	viper.SetDefault("tb.scmbd_role_menu.SQL", "select id, menu_id, role_id from scmbd_role_menu")
	viper.SetDefault("tb.scmbd_user_role.SQL", "select id, user_id, role_id from scmbd_user_role")

	viper.SetDefault("stor.storageDir", ".")

	viper.SetDefault("remote.mode", "sftp")
	viper.SetDefault("remote.host", "127.0.0.1")
	viper.SetDefault("remote.port", 22)
	viper.SetDefault("remote.user", "root")
	viper.SetDefault("remote.password", "")
	viper.SetDefault("remote.workingDir", "/tmp/")
	viper.SetDefault("log", map[string]interface{}{
		"filename": "/tmp/scmbd_sync.log",
		"level":    7,
		"maxlines": 0,
		"maxsize":  0,
		"daily":    true,
		"maxdays":  365,
		"color":    true,
	})

	if err := viper.ReadInConfig(); err != nil {
		logs.Warn("read config file error, error: %v, Use default settings!", err)
	}

	beelogConfig, err := json.Marshal(viper.GetStringMap("log"))
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	logs.SetLogFuncCallDepth(3)
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, string(beelogConfig))

}

// go build -gcflags '-N -l'

func main() {
	var guard *monkey.PatchGuard
	aesKeyHex := "af5c5d660ca7a22cdbedf09465c191492a4ea99695bec34f136d8a15bd3616d2"
	aesNounceHex := "fa8b0eb0aac5ad2331449bbc"
	//using monkey patch to change behaviors of  viper.GetString
	guard = monkey.Patch(viper.GetString, func(key string) string {
		guard.Unpatch()
		defer guard.Restore()
		tmpVal := viper.GetString(key)
		if strings.HasPrefix(tmpVal, "ENC(") && strings.HasSuffix(tmpVal, ")") {
			rxp := regexp.MustCompile(`^ENC\((.*)\)$`)
			cipherHex := rxp.FindStringSubmatch(tmpVal)[1]
			key, _ := hex.DecodeString(aesKeyHex)
			nonce, _ := hex.DecodeString(aesNounceHex)
			cipherText, _ := hex.DecodeString(cipherHex)
			content, err := AesGCMDecrypt(key, nonce, cipherText)
			if err != nil {
				logs.Error(err)
				panic(err)
			}
			return string(content)
		}
		return tmpVal
	})

	// mysql conn uri should add "parseTime=true" or else will panic while parse time values
	connURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&&charset=utf8mb4", viper.GetString("db.user"),
		viper.GetString("db.password"), viper.GetString("db.host"), viper.GetInt("db.port"),
		viper.GetString("db.dbname"))

	localPaths := make(chan string)
	wg.Add(1)
	go func(localPaths chan<- string) {
		defer wg.Done()
		defer close(localPaths)
		for _, schemeName := range viper.GetStringSlice("tb.schemes") {
			logs.Info("begin dump table %s", schemeName)
			//outputPath
			localPaths <- Dump(connURI, viper.GetString(fmt.Sprintf("tb.%s.SQL", schemeName)), viper.GetString("stor.storageDir"), schemeName)
			logs.Info("succeeded dump table %s", schemeName)
		}

	}(localPaths)

	wg.Add(1)

	go func(localPaths <-chan string) {
		defer wg.Done()
		// @TODO init remote conn
		remoteMode := viper.GetString("remote.mode")
		if remoteMode != "ftp" && remoteMode != "sftp" {
			logs.Error("remote storage type only can be ftp or sftp. pleases check your settings!")
			return
		}
		r := &RemoteStor{
			ConnURI:  fmt.Sprintf("%s:%d", viper.GetString("remote.host"), viper.GetInt("remote.port")),
			User:     viper.GetString("remote.user"),
			Password: viper.GetString("remote.password"),
			WorkDir:  viper.GetString("remote.workingDir"),
			Timeout:  time.Duration(viper.GetInt("remote.timeout")) * time.Second,
			ConnType: viper.GetString("remote.mode"),
		}
		logs.Info("start init remote %s connection to %s", r.ConnType, r.ConnURI)
		if err := r.Init(); err != nil {
			logs.Error(err)
			return
		}
		logs.Info("succeeded init remote %s connection to %s", r.ConnType, r.ConnURI)
		defer r.Close()
		for p := range localPaths {
			xlsxPath := path.Join(path.Dir(p), strings.Split(path.Base(p), ".")[0]+".xlsx")
			if err := GenXLSXFromCSV(p, xlsxPath, ","); err != nil {
				logs.Error(err)
				return
			}
			err := os.Remove(p)
			if err != nil {
				logs.Warn(err)
			}
			if err := r.Stor(path.Base(xlsxPath), xlsxPath); err != nil {
				logs.Error("store file %s to remote server %s error: %v", xlsxPath, r.ConnURI, err)
				continue
			}
			logs.Info("succeeded store file %s to remote server %s", xlsxPath, r.ConnURI)
		}
		logs.Info("finished store files to remote server: %s", r.ConnURI)
	}(localPaths)

	wg.Wait()

	logs.Info("end jobs.")
}
