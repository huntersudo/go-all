package main

import (
	"fmt"
	"reflect"
)

func main(){
	// reflect.ValueOf中的Interface方法可以以空接口的形式返回值
	// 接口值.(类型)  断言的语法对接口进行转换
	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(convertPointer)  //0xc00000a098
	fmt.Println(convertValue) // 	1.2345

	// 这里是reflect.ValueOf提供的一些转换到具体类型的方法

	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)//type:int64 value:56

	x1 := reflect.ValueOf(&a).Int()
	fmt.Printf("type:%T value:%v\n", x1, x1)
	// panic(&ValueError{"reflect.Value.Int", v.kind()})
	// return "reflect: call of " + e.Method + " on " + e.Kind.String() + " Value"
	// panic: reflect: call of reflect.Value.Int on ptr Value

	b := "json"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)//type:string value:json
	c := 12.5
	z := reflect.ValueOf(c).Float()
	fmt.Printf("type:%T value:%v\n", z, z)//type:float64 value:12.5
}

/* 这是一个设计方式，错误的需求是  在某件事情上  错误是
type ValueError struct {
	Method string
	Kind   Kind
}

func (e *ValueError) Error() string {
	if e.Kind == 0 {
		return "reflect: call of " + e.Method + " on zero Value"
	}
	return "reflect: call of " + e.Method + " on " + e.Kind.String() + " Value"
}
 */