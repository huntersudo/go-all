package main

import (
	"fmt"
	"reflect"
)

func main(){
	//TODO Elem[ Interface||Ptr ]方法返回 指针或者接口指向的数据
	aa := 56
	xx := reflect.ValueOf(&aa).Elem().Int()
	fmt.Println( reflect.ValueOf(&aa)) //0xc000100028
	fmt.Printf("type:%T value:%v\n", xx, xx) //type:int64 value:56
	fmt.Println()

	// 例子：反射类型是一个空接口，而空接口中包含了int类型的指针
	var z=123
	var y=&z
	var x interface{}=y
	v:=reflect.ValueOf(&x)
	fmt.Println(v)  //0xc000038240
	vx:=v.Elem()
	fmt.Println(vx.Kind())  // interface
	vy:=vx.Elem()
    fmt.Println(vy.Kind())  //ptr || interface中的值
	vz:=vy.Elem()
	fmt.Println(vz.Kind()) // int || 指针的基本类型
    fmt.Println()

	//todo TypeOf的  elem ，返回的具体类型的值
	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())    // chan
	fmt.Println(tc.ChanDir()) // <-chan
	tm := tc.Elem()
	ta, tb := tm.Key(), tm.Elem()
	//           map      key的类型 array  value的类型 slice
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind()) // map array slice
	tx, ty := ta.Elem(), tb.Elem()
	// 数组的类型是 int16 ，value切片的类型是byte(int8)
	fmt.Println(tx.Kind(), ty.Kind()) // int16 uint8

	//TODO SET 修改反射的值
	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	//pointer.SetFloat(77) //panic: reflect: reflect.Value.SetFloat using unaddressable value
	// TODO 反射中存储的实际值是指针时才能赋值，在反射之前，实际值被转换为空接口，存储的值是一个副本，修改它会引起混淆
	fmt.Println("settability of pointer:", pointer.CanSet()) //false
	newValue := pointer.Elem()
	fmt.Println("settability of pointer:", newValue.CanSet()) //true
	// 修改反射的值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)
}



