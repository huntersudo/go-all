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
	engine, _ := xorm.NewEngine("mysql", ConStr)

	// Rows() ⽅法与 Iterate() 类似，不过返回⼀个 Rows 对象由我们⾃⼰迭代，更加灵活：
	rows, _ := engine.Where("age > ? and age < ?", 12, 30).Rows(&User{})
	defer rows.Close()

	u := &User{}
	for rows.Next() {
		rows.Scan(u)

		fmt.Println(u)
	}
	//&{1 pipi  18  2021-06-10 14:39:48 +0800 CST 2021-06-10 14:39:48 +0800 CST}
	//&{2 pipi  18  2021-06-10 14:47:47 +0800 CST 2021-06-10 14:47:47 +0800 CST}
	//&{3 pipi2  20  2021-06-10 14:48:04 +0800 CST 2021-06-10 14:48:04 +0800 CST}
	//&{4 pi2  20  2021-06-10 14:48:20 +0800 CST 2021-06-10 14:48:20 +0800 CST}
}
