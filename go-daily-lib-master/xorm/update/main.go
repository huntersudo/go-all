package main

import (
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
	// 可以传⼊结构指针或 map[string]interface{} 。
	// 对于传⼊结构体指针的情况， xorm 只会更新⾮空的字段。
	// 如果⼀定要更新空字段，需要使⽤ Cols() ⽅法显⽰指定更新的列。
	// 使⽤ Cols() ⽅法指定列后，即使字段为空也会更新

	//mysql> select * from user;
	//	+----+-------+------+------+--------+---------------------+---------------------+
	//	| id | name  | salt | age  | passwd | created             | updated             |
	//		+----+-------+------+------+--------+---------------------+---------------------+
	//	|  1 | pipi  |      |   18 |        | 2021-06-10 14:39:48 | 2021-06-10 14:39:48 |

		engine.ID(1).Update(&User{Name: "ldj"})
	//mysql> select * from user;
	//	+----+-------+------+------+--------+---------------------+---------------------+
	//	| id | name  | salt | age  | passwd | created             | updated             |
	//		+----+-------+------+------+--------+---------------------+---------------------+
	//	|  1 | ldj   |      |   18 |        | 2021-06-10 14:39:48 | 2021-06-10 15:12:32 |


		engine.ID(1).Cols("name", "age").Update(&User{Name: "dj"})
	// 这里age 被更新了  0
	//+----+-------+------+------+--------+---------------------+---------------------+
	//| id | name  | salt | age  | passwd | created             | updated             |
	//	+----+-------+------+------+--------+---------------------+---------------------+
	//|  1 | dj    |      |    0 |        | 2021-06-10 14:39:48 | 2021-06-10 15:13:31 |


	// 由于使⽤ map[string]interface{} 类型的参数， xorm ⽆法推断表名，必须使⽤ Table() ⽅法指定。
	// 第⼀个 Update() ⽅法只会更新 name 字 段，其他空字段不更新。
	// 第⼆个 Update() ⽅法会更新 name 和 age 两个字段， age 被更新为 0
	engine.Table(&User{}).ID(1).Update(map[string]interface{}{"age": 38})
	//+----+-------+------+------+--------+---------------------+---------------------+
	//| id | name  | salt | age  | passwd | created             | updated             |
	//	+----+-------+------+------+--------+---------------------+---------------------+
	//|  1 | dj    |      |   38 |        | 2021-06-10 14:39:48 | 2021-06-10 15:15:58 |


}
