package main

import "reflect"

type  children struct {
	Age int
}
type Nested struct {
	X int
	Child children
}

func main(){
	// 嵌套结构体
	vs := reflect.ValueOf(&Nested{}).Elem()
	vz := vs.Field(1)
	vz.Set(reflect.ValueOf(children{ Age:19 }))
}