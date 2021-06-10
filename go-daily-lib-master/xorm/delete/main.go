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
	Deleted time.Time `xorm:"deleted"`
}
const ConStr ="sml_user:1qaz@WSX@tcp(192.168.88.132:3306)/sml_db?charset=utf8"

func main() {

	//如果我们为 time.Time/int/int64 这些类型的字段设置 xorm:"created" 标签，插⼊数据时，该字段会⾃动更新为当前时间；
	//如果我们为 tiem.Time/int/int64 这些类型的字段设置 xorm:"updated" 标签，插⼊和更新数据时，该字段会⾃动更新为当前时间；
	//如果我们为 time.Time 类型的字段设置了 xorm:"deleted" 标签，删除数据时，只是设置删除时间，并不真正删除记录。
  //todo 已删除的记录必须使⽤ Unscoped() ⽅法查询，如果要真正 删除某条记录，也 可以使⽤ Unscoped() 。

	engine, _ := xorm.NewEngine("mysql", ConStr)
    engine.Sync2(&User{})
	affected, _ := engine.Where("name = ?", "pi2").Delete(&User{})
	fmt.Printf("%d records deleted", affected)
	fmt.Println()
	u:=&User{}
	engine.Where("name = ?","pi2").Unscoped().Get(u)
	fmt.Println("after delete:",u)

	// 0 records deleted
	//after delete: &{4 pi2  20  2021-06-10 14:48:20 +0800 CST 2021-06-10 14:48:20 +0800 CST 2021-06-10 15:37:37 +0800 CST}


	//mysql> select * from user;
	//	+----+-------+------+------+--------+---------------------+---------------------+---------------------+
	//	| id | name  | salt | age  | passwd | created             | updated             | deleted             |
	//		+----+-------+------+------+--------+---------------------+---------------------+---------------------+
	//	|  1 | dj    |      |   38 |        | 2021-06-10 14:39:48 | 2021-06-10 15:15:58 | NULL                |
	//	|  2 | pipi  |      |   18 |        | 2021-06-10 14:47:47 | 2021-06-10 14:47:47 | NULL                |
	//	|  3 | pipi2 |      |   20 |        | 2021-06-10 14:48:04 | 2021-06-10 14:48:04 | NULL                |
	//	|  4 | pi2   |      |   20 |        | 2021-06-10 14:48:20 | 2021-06-10 14:48:20 | 2021-06-10 15:37:37 |
	//	|  6 | xhq   |      |   41 |        | 2021-06-10 15:01:57 | 2021-06-10 15:01:57 | NULL                |
	//	|  7 | lhy   |      |   12 |        | 2021-06-10 15:01:57 | 2021-06-10 15:01:57 | NULL                |
	//	|  9 | xhq   |      |   41 |        | 2021-06-10 15:12:26 | 2021-06-10 15:12:26 | NULL                |
	//	| 10 | lhy   |      |   12 |        | 2021-06-10 15:12:26 | 2021-06-10 15:12:26 | NULL                |
	//		+----+-------+------+------+--------+---------------------+---------------------+---------------------+
	//		8 rows in set (0.00 sec)
	//

	}
