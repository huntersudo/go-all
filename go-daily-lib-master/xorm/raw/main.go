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

// 执⾏原始的 SQL

func main() {
	engine, _ := xorm.NewEngine("mysql", ConStr)

	// Query() ⽅法返回 []map[string][]byte ，切⽚中的每个元素都代表⼀条记录， map 的键对应列名， []byte 为值。
	//还有 QueryInterface() ⽅法返回 []map[string]interface{} ，
	// QueryString() ⽅法返回 []map[string]interface{} 。
	querySql := "select * from user limit 1"
	reuslts, _ := engine.Query(querySql)
	for _, record := range reuslts {
		for key, val := range record {
			fmt.Println(key+":", string(val))
		}
	}

	updateSql := "update `user` set name=? where id=?"
	res, _ := engine.Exec(updateSql, "dj_update", 1)
	fmt.Println(res.RowsAffected())
   // deleted:
	//id: 1
	//name: dj
	//salt:
	//age: 38
	//passwd:
	//created: 2021-06-10 14:39:48
	//updated: 2021-06-10 15:15:58

	//1 <nil>

	// mysql> select * from user;
	//+----+-----------+------+------+--------+---------------------+---------------------+---------------------+
	//| id | name      | salt | age  | passwd | created             | updated             | deleted             |
	//+----+-----------+------+------+--------+---------------------+---------------------+---------------------+
	//|  1 | dj_update |      |   38 |        | 2021-06-10 14:39:48 | 2021-06-10 15:15:58 | NULL                |
	//|  2 | pipi      |      |   18 |        | 2021-06-10 14:47:47 | 2021-06-10 14:47:47 | NULL                |
	//|  3 | pipi2     |      |   20 |        | 2021-06-10 14:48:04 | 2021-06-10 14:48:04 | NULL                |
	//|  4 | pi2       |      |   20 |        | 2021-06-10 14:48:20 | 2021-06-10 14:48:20 | 2021-06-10 15:37:37 |
	//|  6 | xhq       |      |   41 |        | 2021-06-10 15:01:57 | 2021-06-10 15:01:57 | NULL                |
	//|  7 | lhy       |      |   12 |        | 2021-06-10 15:01:57 | 2021-06-10 15:01:57 | NULL                |
	//|  9 | xhq       |      |   41 |        | 2021-06-10 15:12:26 | 2021-06-10 15:12:26 | NULL                |
	//| 10 | lhy       |      |   12 |        | 2021-06-10 15:12:26 | 2021-06-10 15:12:26 | NULL                |
	//+----+-----------+------+------+--------+---------------------+---------------------+---------------------+
	//8 rows in set (0.00 sec)


}
