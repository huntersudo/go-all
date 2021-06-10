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

	// 与 Find() ⼀样， Iterate() 也是找到满⾜条件的所有记录，只不过传⼊了⼀个回调去处理每条记录：
	// 如果回调返回⼀个⾮ nil 的错误，后⾯的记录就不会再处理了。

	engine.Where("age > ? and age < ?", 12, 30).Iterate(&User{}, func(i int, bean interface{}) error {
		fmt.Printf("user%d:%v\n", i, bean.(*User))
		return nil
	})

//user0:&{1 pipi  18  2021-06-10 14:39:48 +0800 CST 2021-06-10 14:39:48 +0800 CST}
//user1:&{2 pipi  18  2021-06-10 14:47:47 +0800 CST 2021-06-10 14:47:47 +0800 CST}
//user2:&{3 pipi2  20  2021-06-10 14:48:04 +0800 CST 2021-06-10 14:48:04 +0800 CST}
//user3:&{4 pi2  20  2021-06-10 14:48:20 +0800 CST 2021-06-10 14:48:20 +0800 CST}


	// mysql> select * from user;
	//+----+-------+------+------+--------+---------------------+---------------------+
	//| id | name  | salt | age  | passwd | created             | updated             |
	//+----+-------+------+------+--------+---------------------+---------------------+
	//|  1 | pipi  |      |   18 |        | 2021-06-10 14:39:48 | 2021-06-10 14:39:48 |
	//|  2 | pipi  |      |   18 |        | 2021-06-10 14:47:47 | 2021-06-10 14:47:47 |
	//|  3 | pipi2 |      |   20 |        | 2021-06-10 14:48:04 | 2021-06-10 14:48:04 |
	//|  4 | pi2   |      |   20 |        | 2021-06-10 14:48:20 | 2021-06-10 14:48:20 |
	//+----+-------+------+------+--------+---------------------+---------------------+
	//4 rows in set (0.00 sec)
}
