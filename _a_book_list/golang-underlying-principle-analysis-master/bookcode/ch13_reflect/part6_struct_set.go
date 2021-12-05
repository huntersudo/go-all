package main

import "reflect"

var s  struct {
	X int // an exported field
	y float64 // a non-exported field
}


func main(){
	//错误 TODO reflec.ValueOf() 参数为空接口,如果将值类型赋值到空接口，会产生一次复制，所以必须传递指针
/*	vs := reflect.ValueOf(s)
	vx:= vs.Field(0)
	vb := reflect.ValueOf(123)
	vx.Set(vb) //panic: reflect: reflect.Value.Set using unaddressable value
	*/

	//TODO 因为field方法的接收器是结构体, 所以必须用Elem 获取指针指向的结构体值类型
	// vs := reflect.ValueOf(&s)  //panic: reflect: call of reflect.Value.Field on ptr Value
	vs := reflect.ValueOf(&s).Elem()
	vx:= vs.Field(0) // 这里必须是结构体调用
	vb := reflect.ValueOf(123)
	vx.Set(vb)
}