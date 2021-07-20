package main

import (
	"log"
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

const ConStr ="sml_user:1qaz@WSX@tcp(192.168.88.128:3306)/sml_db?charset=utf8"
func main() {
	engine, err := xorm.NewEngine("mysql",ConStr)
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}
}
