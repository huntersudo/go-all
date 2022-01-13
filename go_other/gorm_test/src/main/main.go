package main

/**
jdbc:mysql://10.154.2.147:3306/scas
*/
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type GormUser struct {
	Id    int
	Name  string
	Age   int
	Sex   byte
	Phone string
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "scas:123456@tcp(10.154.2.147:3306)/scas?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//设置全局表名禁用复数
	// // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	db.SingularTable(true)

	db.LogMode(true)


	////可以通过定义 DefaultTableNameHandler 对默认表名应用任何规则。
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string {
	//	return "prefix_" + defaultTableName;
	//}

}

func main() {
	//gormUser := GormUser{Id: 2,Name:"XIAOHONG",Age: 10,Sex: 1,Phone: "18896725679"}
	//gormUser := GormUser{Id: 1,Name:"xiaoming"}
	//gormUser.Insert()
	update2()
}

func (user *GormUser) Insert() {
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table("gorm_user").Create(user)
}

func update() {
	//注意，Model方法必须要和Update方法一起使用
	//使用效果相当于Model中设置更新的主键key（如果没有where指定，那么默认更新的key为id），Update中设置更新的值
	//如果Model中没有指定id值，且也没有指定where条件，那么将更新全表
	//log: UPDATE `gorm_user` SET `id` = 1, `name` = 'wanGang'  WHERE `gorm_user`.`id` = 1
	user := GormUser{Id: 1, Name: "wanGang"}
	db.Model(&user).Update(user)
}

func update1() {
	//注意到上面Update中使用了一个Struct，你也可以使用map对象。
	//需要注意的是：使用Struct的时候，只会更新Struct中这些非空的字段。
	//对于string类型字段的""，int类型字段0，bool类型字段的false都被认为是空白值，不会去更新表
	//下面这个更新操作只使用了where条件没有在Model中指定id
	//2句：
	//UPDATE `gorm_user` SET `name` = 'xiaohong'  WHERE (id = 1)
	//UPDATE `gorm_user` SET `age` = 10  WHERE (id = 1)
	db.Model(&GormUser{}).Where("id = ?", 1).Update("name", "xiaohong").Update("age", 10)

}

func update2()  {

	//如果你想手动将某个字段set为空值, 可以使用单独选定某些字段的方式来更新：
	gormUserModel := GormUser{Id: 2}
	// UPDATE `gorm_user` SET `name` = ''  WHERE `gorm_user`.`id` = 1
	//db.Model(&gormUserModel).Select("name").Update(map[string]interface{}{"name": "","age":0})

	// UPDATE `gorm_user` SET `age` = 0, `name` = ''  WHERE `gorm_user`.`id` = 1
	//db.Model(&gormUserModel).Select("name","age").Update(map[string]interface{}{"name": "","age":0})

	//忽略掉某些字段：
	//当你的更新的参数为结构体，而结构体中某些字段你又不想去更新，那么可以使用Omit方法过滤掉这些不想update到库的字段：

	gormUser2 := GormUser{Id: 3,Name:"xioaming",Age:12}
	db.Model(&gormUserModel).Omit("name").Update(&gormUser2)
    //UPDATE `gorm_user` SET `age` = 12, `id` = 1  WHERE `gorm_user`.`id` = 1
    //UPDATE `gorm_user` SET `age` = 12, `id` = 3  WHERE `gorm_user`.`id` = 3
}
/**


CREATE TABLE `gorm_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `age` int(3) NOT NULL DEFAULT '0',
  `sex` tinyint(3) NOT NULL DEFAULT '0',
  `phone` varchar(40) NOT NULL DEFAULT '',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4


*/
