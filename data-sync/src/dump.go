package main

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

func initDB(connURI string) *sql.DB {
	db, err := sql.Open("mysql", connURI)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	return db
}

// ScmbdUser struct of table scmbd_user
type ScmbdUser struct {
	ID         int64
	Name       string
	Username   string
	CreateTime sql.NullTime
	UpdateTime sql.NullTime
	UserStatus int
	OrgID      int64
}

func parseUserColumn(u *ScmbdUser) string {
	id := strconv.FormatInt(u.ID, 10)
	createTime := ""
	if u.CreateTime.Valid {
		//createTime = u.CreateTime.Time.Format("20060102150405") // @TODO time layout should be clarified
		createTime = strconv.FormatInt(u.CreateTime.Time.UnixNano()/1000000, 10)
	}
	updateTime := ""
	if u.UpdateTime.Valid {
		// updateTime = u.UpdateTime.Time.Format("20060102150405") // @TODO
		updateTime = strconv.FormatInt(u.UpdateTime.Time.UnixNano()/1000000, 10)
	}
	userStatus := strconv.Itoa(u.UserStatus)
	orgID := strconv.FormatInt(u.OrgID, 10)
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s", id, u.Name, u.Username, createTime, updateTime, userStatus, orgID)
}

// ScmbdOrg struct of table scmbd_org
type ScmbdOrg struct {
	ID         int64
	Name       string
	ParentID   sql.NullInt64
	CreateTime sql.NullTime
	UpdateTime sql.NullTime
}

func parseOrgColumn(o *ScmbdOrg) string {
	id := strconv.FormatInt(o.ID, 10)
	parentID := ""
	if o.ParentID.Valid {
		parentID = strconv.FormatInt(o.ParentID.Int64, 10)
	}

	createTime := ""
	if o.CreateTime.Valid {
		//createTime = o.CreateTime.Time.Format("20060102150405")
		createTime = strconv.FormatInt(o.CreateTime.Time.UnixNano()/1000000, 10)
	}

	updateTime := ""
	if o.UpdateTime.Valid {
		//updateTime = o.UpdateTime.Time.Format("20060102150405")
		updateTime = strconv.FormatInt(o.UpdateTime.Time.UnixNano()/1000000, 10)
	}

	return fmt.Sprintf("%s,%s,%s,%s, %s", id, o.Name, parentID, createTime, updateTime)

}

// ScmbdMenu struct for table scmbd_menu
type ScmbdMenu struct {
	ID             int64
	BoardID        sql.NullInt64
	Identification sql.NullString
	Name           string
	ParentID       sql.NullInt64
	LeafType       int
	URL            sql.NullString
	URLType        sql.NullInt64
	CreateTime     sql.NullTime
	UpdateTime     sql.NullTime
}

func parseMenuColumn(m *ScmbdMenu) string {
	id := strconv.FormatInt(m.ID, 10)
	boardID := ""
	if m.BoardID.Valid {
		boardID = strconv.FormatInt(m.BoardID.Int64, 10)
	}
	identification := ""
	if m.Identification.Valid {
		identification = m.Identification.String
	}
	parentID := ""
	if m.ParentID.Valid {
		parentID = strconv.FormatInt(m.ParentID.Int64, 10)
	}
	leafType := strconv.Itoa(m.LeafType)
	url := ""
	if m.URL.Valid {
		url = m.URL.String
	}
	urlType := ""
	if m.URLType.Valid {
		urlType = strconv.FormatInt(m.URLType.Int64, 10)
	}
	createTime := ""
	if m.CreateTime.Valid {
		// createTime = m.CreateTime.Time.Format("20060102150405")
		createTime = strconv.FormatInt(m.CreateTime.Time.UnixNano()/1000000, 10)
	}
	updateTime := ""
	if m.UpdateTime.Valid {
		// updateTime = m.UpdateTime.Time.Format("20060102150405")
		updateTime = strconv.FormatInt(m.UpdateTime.Time.UnixNano()/1000000, 10)
	}
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s", id, boardID, identification, m.Name, parentID, leafType, url, urlType, createTime, updateTime)
}

// ScmbdRole structs for table scmbd_role
type ScmbdRole struct {
	ID         int64
	Name       string
	CreateTime sql.NullTime
	UpdateTime sql.NullTime
}

func parseRoleColumn(r *ScmbdRole) string {
	id := strconv.FormatInt(r.ID, 10)
	createTime := ""
	if r.CreateTime.Valid {
		// createTime = r.CreateTime.Time.Format("20060102150405")
		createTime = strconv.FormatInt(r.CreateTime.Time.UnixNano()/1000000, 10)
	}
	updateTime := ""
	if r.UpdateTime.Valid {
		// updateTime = r.UpdateTime.Time.Format("20060102150405")
		updateTime = strconv.FormatInt(r.UpdateTime.Time.UnixNano()/1000000, 10)
	}
	return fmt.Sprintf("%s,%s,%s,%s", id, r.Name, createTime, updateTime)
}

// ScmbdRoleMenu struct for table scmbd_role_menu
type ScmbdRoleMenu struct {
	ID     int64
	MenuID int64
	RoleID int64
}

func parseRoleMenuColumn(rm *ScmbdRoleMenu) string {
	return fmt.Sprintf("%d,%d,%d", rm.ID, rm.MenuID, rm.RoleID)
}

// ScmbdUserRole struct for table scmbd_user_role
type ScmbdUserRole struct {
	ID     int64
	UserID int64
	RoleID int64
}

func parseUserRleColumn(ur *ScmbdUserRole) string {
	return fmt.Sprintf("%d,%d,%d", ur.ID, ur.UserID, ur.RoleID)
}

// Dump dumps table scmbd_org scmbd_menu...
func Dump(connURI, querySQL, outputDIR, schema string) (outputPath string) {
	db := initDB(connURI)
	defer db.Close()

	results, err := db.Query(querySQL)
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	curTimeStamp := time.Now().Format("20060102") // using date precision
	outputPath = path.Join(outputDIR, schema+"_"+curTimeStamp+".csv")
	f, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	defer f.Close()
	for results.Next() {
		contentLine := ""
		switch schema {
		case "scmbd_user":
			var u ScmbdUser
			err = results.Scan(&u.ID, &u.Name, &u.Username, &u.CreateTime, &u.UpdateTime, &u.UserStatus, &u.OrgID)
			if err != nil {
				logs.Error(err)
				panic(err)
			}
			contentLine = parseUserColumn(&u)
		case "scmbd_org":
			var o ScmbdOrg
			err = results.Scan(&o.ID, &o.Name, &o.ParentID, &o.CreateTime, &o.UpdateTime)
			if err != nil {
				logs.Error(err)
				panic(err)
			}
			contentLine = parseOrgColumn(&o)
		case "scmbd_menu":
			var m ScmbdMenu
			err = results.Scan(&m.ID, &m.BoardID, &m.Identification, &m.Name, &m.ParentID, &m.LeafType, &m.URL, &m.URLType, &m.CreateTime, &m.UpdateTime)
			if err != nil {
				logs.Error(err)
				panic(err)
			}
			contentLine = parseMenuColumn(&m)
		case "scmbd_role":
			var r ScmbdRole
			err = results.Scan(&r.ID, &r.Name, &r.CreateTime, &r.UpdateTime)
			if err != nil {
				logs.Error(err)
				panic(err)
			}
			contentLine = parseRoleColumn(&r)
		case "scmbd_role_menu":
			var rm ScmbdRoleMenu
			err = results.Scan(&rm.ID, &rm.MenuID, &rm.RoleID)
			if err != nil {
				logs.Error(err)
				panic(err)
			}
			contentLine = parseRoleMenuColumn(&rm)
		case "scmbd_user_role":
			var ur ScmbdUserRole
			err = results.Scan(&ur.ID, &ur.UserID, &ur.RoleID)
			if err != nil {
				logs.Error(err)
				panic(err)
			}
			contentLine = parseUserRleColumn(&ur)
		default:
			logs.Error("can not find table scheme %s", schema)
			panic("can not find table " + schema)
		}
		_, err := f.WriteString(contentLine + "\n")
		if err != nil {
			logs.Error(err)
			panic(err)
		}
	}
	return
}
