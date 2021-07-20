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
	engine.Insert(&User{Name:"pi2", Age:20})
	slcUsers:= make([]User, 1)

	// Get() ⽅法只能返回单条记录，其⽣成的 SQL 语句总是有 LIMIT 1 。 Find() ⽅法返回所有符合条件的记录。 Find() 需要传⼊对象切⽚的指针 或 map 的指针：
    // map 的键为主键，所以如果表为复合主键就不能使⽤这种⽅式了。
	engine.Where("age > ? and age < ?", 12, 30).Find(&slcUsers)
	fmt.Println("users whose age between [12,30]:", slcUsers)

	mapUsers := make(map[int64]User)
	engine.Where("length(name) = ?", 3).Find(&mapUsers)
	fmt.Println("users whose has name of length 3:", mapUsers)
	//users whose age between [12,30]: [{0   0  0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} {1 pipi  18  2021-06-10 14:39:48 +0800 CST 2021-06-10 14:39:48 +0800 CST} {2 pipi  18  2021-06-10 14:47:47 +0800 CST 2021-06-10 14:47:47 +0800 CST} {3 pipi2  20  2021-06-10 14:48:04 +0800 CST 2021-06-10 14:48:04 +0800 CST} {4 pi2  20  2021-06-10 14:48:20 +0800 CST 2021-06-10 14:48:20 +0800 CST}]
	//users whose has name of length 3: map[4:{4 pi2  20  2021-06-10 14:48:20 +0800 CST 2021-06-10 14:48:20 +0800 CST}]

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
