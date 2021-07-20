package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// User
/**
 上⾯演⽰了 3 种使⽤ Get() 的⽅式：
 使⽤主键： engine.ID(1) 查询主键（即 id ）为 1 的⽤户；
 使⽤条件语句： engine.Where("name=?", "dj") 查询 name = "dj" 的⽤户；
 使⽤对象中的⾮空字段：
   user3 设置了 Id 字段为 5， engine.Get(user3) 查询 id = 5 的⽤户；
  user4 设置了字段 Name 为 "pipi" ， engine.Get(user4) 查询 name = "pipi" 的⽤户。
 */

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

const ConStr ="sml_user:1qaz@WSX@tcp(192.168.88.129:3306)/sml_db?charset=utf8"

func main() {
	engine, _ := xorm.NewEngine("mysql", ConStr)

	//engine.Insert(&User{Name:"dj", Age:18})

	user1 := &User{}
	has, _ := engine.ID(2).Get(user1)
	if has {
		fmt.Printf("user1:%v\n", user1)
	}

	user2 := &User{}
	has, _ = engine.Where("name=?", "dj").Get(user2)
	if has {
		fmt.Printf("user2:%v\n", user2)
	}

	user3 := &User{Id: 3}
	has, _ = engine.Get(user3)
	if has {
		fmt.Printf("user3:%v\n", user3)
	}

	engine.Insert(&User{Name:"pipi", Age:18})

	user4 := &User{Name: "pipi"}
	has, _ = engine.Get(user4)
	if has {
		fmt.Printf("user4:%v\n", user4)
	}
}
