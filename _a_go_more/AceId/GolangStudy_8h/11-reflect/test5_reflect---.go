package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is called ..")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Aceld", 18}

	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//todo 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name()) // inputType is : User

	//todo  获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue) // inputValue is: {1 Aceld 18}

	//通过type 获取里面的字段 todo  NumField
	//1. 获取interface的reflect.Type，通过Type得到NumField ,进行遍历
	//2. 得到每个field，数据类型
	//3. 通过field有一个Interface()方法等到 对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	// Id: int = 1
	//Name: string = Aceld
	//Age: int = 18

	//通过type 获取里面的方法,调用 todo NumMethod
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
	// Call: func(main.User)

}
