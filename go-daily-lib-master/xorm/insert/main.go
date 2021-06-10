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
	user := &User{
		Name: "lzy",
		Age: 50,
	}

	affected, _ := engine.Insert(user)
	fmt.Printf("%d records inserted, user.id:%d\n", affected, user.Id)

	users := make([]*User, 2)
	users[0] = &User{Name: "xhq", Age: 41}
	users[1] = &User{Name: "lhy", Age: 12}

	affected, _ = engine.Insert(&users)
	fmt.Printf("%d records inserted, id1:%d, id2:%d", affected, users[0].Id, users[1].Id)

	//1 records inserted, user.id:5
	//2 records inserted, id1:0, id2:0
	// todo 批量插⼊时，每个对象的 Id 字段不会被⾃动赋值，所以上⾯最后⼀⾏输 出 id1 和 id2 均为 0。
	//mysql> select * from user;
	//	+----+-------+------+------+--------+---------------------+---------------------+
	//	| id | name  | salt | age  | passwd | created             | updated             |
	//		+----+-------+------+------+--------+---------------------+---------------------+
	//	|  1 | pipi  |      |   18 |        | 2021-06-10 14:39:48 | 2021-06-10 14:39:48 |
	//	|  2 | pipi  |      |   18 |        | 2021-06-10 14:47:47 | 2021-06-10 14:47:47 |
	//	|  3 | pipi2 |      |   20 |        | 2021-06-10 14:48:04 | 2021-06-10 14:48:04 |
	//	|  4 | pi2   |      |   20 |        | 2021-06-10 14:48:20 | 2021-06-10 14:48:20 |
	//	|  5 | lzy   |      |   50 |        | 2021-06-10 15:01:57 | 2021-06-10 15:01:57 |
	//	|  6 | xhq   |      |   41 |        | 2021-06-10 15:01:57 | 2021-06-10 15:01:57 |
	//	|  7 | lhy   |      |   12 |        | 2021-06-10 15:01:57 | 2021-06-10 15:01:57 |
	//		+----+-------+------+------+--------+---------------------+---------------------+

	}
