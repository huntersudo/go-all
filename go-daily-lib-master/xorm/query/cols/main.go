package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
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


// 除此之外 xorm ⽀持只返回指定的列（ xorm.Cols() ）或忽略特定的列（ xorm.Omit() ）：
const ConStr ="sml_user:1qaz@WSX@tcp(192.168.88.128:3306)/sml_db?charset=utf8"

func main() {
	engine, _ := xorm.NewEngine("mysql", ConStr)

	engine.ShowSQL(true)

	user1 := &User{}
	engine.ID(1).Cols("id", "name", "age").Get(user1)
	// SELECT `id`, `name`, `age` FROM `user` WHERE `id`=? LIMIT 1 [1]
	fmt.Printf("user1:%v\n", user1)

	user2 := &User{Name: "pipi"}
	engine.Omit("created", "updated").Get(user2)
	// SELECT `id`, `name`, `salt`, `age`, `passwd` FROM `user` WHERE `name`=? LIMIT 1 [pipi] - 969.2µs
	fmt.Printf("user2:%v\n", user2)

   //[xorm] [info]  2021/06/09 18:16:06.224451 [SQL] SELECT `id`, `name`, `age` FROM `user` WHERE `id`=? LIMIT 1 [1] - 12.9653ms
	//user1:&{1 dj  18  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}
	//[xorm] [info]  2021/06/09 18:16:06.247390 [SQL] SELECT `id`, `name`, `salt`, `age`, `passwd` FROM `user` WHERE `name`=? LIMIT 1 [pipi]
	//user2:&{4 pipi  18  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC}

}
