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

const ConStr ="sml_user:1qaz@WSX@tcp(192.168.88.132:3306)/sml_db?charset=utf8"
func main() {
	// Count() ⽅法统计满⾜条件的记录数量：
	engine, _ := xorm.NewEngine("mysql", ConStr)

	num, _ := engine.Where("age >= ?", 10).Count(&User{})
	fmt.Printf("there are %d users whose age >= 10", num)
}
