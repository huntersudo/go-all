package main

import (
	"fmt"
	"reflect"
)


// todo valueOf  typeOf  打印的是 pair的值
func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))  // type :  float64
	fmt.Println("value : ", reflect.ValueOf(arg)) // value :  1.2345
}

func main() {
	var num float64 = 1.2345

	reflectNum(num)
}
