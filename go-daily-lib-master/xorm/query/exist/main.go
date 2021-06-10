package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

const ConStr ="sml_user:1qaz@WSX@tcp(192.168.88.132:3306)/sml_db?charset=utf8"

func main() {
	engine, _ := xorm.NewEngine("mysql", ConStr)

	//f, err := os.Create("D:\\workspace\\ws1\\src\\github.com\\go-all\\go-daily-lib-master\\xorm\\query\\debug\\log\\sql_sml.log")
	//if err != nil {
	//	panic(err)
	//}
	//
	//engine.SetLogger(log.NewSimpleLogger(f))
	//注意 log.NewSimpleLogger(f) 是 xorm 的⼦包 xorm.io/xorm/log 提供的简单⽇志封装， ⽽⾮标准库 log
	//engine.ShowSQL(true)

	user1 := &User{}
	has, _ := engine.ID(1).Exist(user1)
	if has {
		fmt.Println("user with id=1 exist")
	} else {
		fmt.Println("user with id=1 not exist")
	}

	user2 := &User{}
	has, _ = engine.Where("name=?", "dj2").Get(user2)
	if has {
		fmt.Println("user with name=dj2 exist")
	} else {
		fmt.Println("user with name=dj2 not exist")
	}
// sql
//[xorm] [info]  2021/06/10 14:41:44.835046 [SQL] SELECT `id`, `name`, `salt`, `age`, `passwd`, `created`, `updated` FROM `user` WHERE `id`=? LIMIT 1 [1] - 10.9439ms
//[xorm] [info]  2021/06/10 14:41:44.858027 [SQL] SELECT `id`, `name`, `salt`, `age`, `passwd`, `created`, `updated` FROM `user` WHERE (name=?) LIMIT 1 [dj2] - 1.0123ms



}
