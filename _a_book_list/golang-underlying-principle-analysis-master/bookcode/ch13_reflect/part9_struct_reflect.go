package main

import (
	"fmt"
	"reflect"
)

func MakeStruct(vals ...interface{}) reflect.Value {
	var sfs []reflect.StructField
	for k, v := range vals {
		t := reflect.TypeOf(v)
		sf := reflect.StructField{
			Name: fmt.Sprintf("F%d", (k + 1)),
			Type: t,
		}
		sfs = append(sfs, sf)
	}
	st := reflect.StructOf(sfs)  // todo 生成反射 结构体，参数为 StructField的切片
	so := reflect.New(st)  // todo  reflect.New 生成指定类型的反射指针对象
	return so
}

func main(){
/**
	struct {
	int
	string
	[]int
	}
 */
	sr := MakeStruct(0, "", []int{})

	fmt.Println(sr.Elem().Field(0).Interface()) // 0
	sr.Elem().Field(0).SetInt(20)
	fmt.Println(sr.Elem().Field(0).Interface())  //20

	fmt.Println(sr.Elem().Field(1).Interface())  // 空行
	sr.Elem().Field(1).SetString("reflect me")
	fmt.Println(sr.Elem().Field(1).Interface())  // reflect me

	fmt.Println(sr.Elem().Field(2).Interface()) // []
	v := []int{1, 2, 3}
	rv := reflect.ValueOf(v)
	sr.Elem().Field(2).Set(rv)
	fmt.Println(sr.Elem().Field(2).Interface())  //[1 2 3]
}
/**

0
20

reflect me
[]
[1 2 3]

 */