package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Player struct {
	Id int64
	Name string
	Age int
	Level int
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

const ConStr ="sml_user:1qaz@WSX@tcp(192.168.88.132:3306)/sml_db?charset=utf8"

func main() {
	// 使⽤ xorm 来操作数据库，⾸先需要使⽤ xorm.NewEngine() 创建⼀个引擎

	engine, err := xorm.NewEngine("mysql", ConStr)
     if err!=nil {
     	fmt.Errorf("connect db error",err)
	 }
	r,e:=engine.Exec(" select version()")
	fmt.Println(r)
	fmt.Println(e)
	fmt.Println("ddd")
	// 表结构同步： 调⽤ Sync2() ⽅法会根据 User 的结构⾃动创建⼀个 user
	// 如果表 user 已经存在， Sync() ⽅法会对⽐ User 结构与表结构的不同，对表做相应的修改。我们给 User 结构添加⼀个 Level 字段
	// *此修改只限于添加字段。
	engine.Sync2(&Player{})
	engine.Insert(&Player{Name:"dj", Age:18})

	//
	p1 := &Player{}
	engine.Where("name = ?", "dj").Get(p1)
	fmt.Println("after insert:", p1)
	time.Sleep(5 * time.Second)
	//
	//engine.Table(&Player{}).ID(p1.Id).Update(map[string]interface{}{"age":30})
	////
	//p2 := &Player{}
	//engine.Where("name = ?", "dj").Get(p2)
	//fmt.Println("after update:", p2)
	//time.Sleep(5 * time.Second)
	////
	//engine.ID(p1.Id).Delete(&Player{})
	//
	//p3 := &Player{}
	//engine.Where("name = ?", "dj").Unscoped().Get(p3)
	//fmt.Println("after delete:", p3)
}