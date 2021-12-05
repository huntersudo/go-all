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

func (u User) RefCallArgs( age int, name string) error {
	return nil
}

func (u User) RefCallNoArgs() {
	fmt.Println("hello world")
}

func (u User) RefCallPoint(name string, age *int) {
	fmt.Println(" name: ", name, ", age:", *age)
}

func (u *User) RefPointMethod() {
	fmt.Println("RefPointMethod:hello world")
}

func (u *User) PointMethodReturn(name string, age int)(string, int) {
	return name,age
}


func main(){
	getType:=reflect.TypeOf(&User{})
	for i:=0;i<getType.NumMethod();i++{
		//m:=getType.Method(i)
		//fmt.Printf("%s: %v\n",m.Name,m.Type)
	}
	/**
	PointMethodReturn: func(*main.User, string, int) (string, int)
	RefCallArgs: func(*main.User, int, string) error
	RefCallNoArgs: func(*main.User)
	RefCallPoint: func(*main.User, string, *int)
	RefPointMethod: func(*main.User)
	 */


	// 常用 MethodByName
	user := User{1, "jonson", 25}
	ref := reflect.ValueOf(user)
	tf:= ref.MethodByName("RefCallArgs").Type()
	// 获取函数的参数个数，返回值个数，结构体方法个数
	fmt.Printf("numIn:%d,numOut:%d,numMethod:%d\n",tf.NumIn(), tf.NumOut(),ref.NumMethod()) //2  1  3

	// TODO 在运行时调用该方法
	m := ref.MethodByName("RefCallNoArgs")
	args := make([]reflect.Value, 0)
	m.Call(args)  // 打印 hell oworld

	m = ref.MethodByName("RefCallArgs")
	args  = []reflect.Value{reflect.ValueOf(18),reflect.ValueOf("json")}
	m.Call(args)

	//TODO 指针参数
	m = ref.MethodByName("RefCallPoint")
	age := 19
	args = []reflect.Value{reflect.ValueOf("jonson"), reflect.ValueOf(&age)}
	m.Call(args) // name:  jonson , age: 19

	// TODO 方法接收者是指针
	//refnew := reflect.ValueOf(user)  //panic: reflect: call of reflect.Value.Call on zero Value
	refnew := reflect.ValueOf(&user)  //
	m = refnew.MethodByName("RefPointMethod")
	args  = make([]reflect.Value, 0)
	m.Call(args) //RefPointMethod:hello world

	// TODO 有返回值
	m  = refnew.MethodByName("PointMethodReturn")
	args = []reflect.Value{reflect.ValueOf("jonson"), reflect.ValueOf(30)}
	res:= m.Call(args)
	fmt.Println("return name:",res[0].Interface()) //json
	fmt.Println("return age:",res[1].Interface())  //30
}