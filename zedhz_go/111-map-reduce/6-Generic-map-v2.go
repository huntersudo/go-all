package main

import (
	"fmt"
	"reflect"
)

// todo 所以，如果要写一个健壮的程序，对于这种用interface{} 的“过度泛型”，就需要我们自己来做类型检查。来看一个有类型检查的 Map 代码：

func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {

	//check the `slice` type is Slice
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	//check the function signature
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}

	return sliceOutType.Interface()

}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {

	//Check it is a funciton
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}
// todo 我来列一下代码中的几个要点。
//代码中没有使用 Map 函数，因为和数据结构有含义冲突的问题，所以使用Transform，这个来源于 C++ STL 库中的命名。
//有两个版本的函数，一个是返回一个全新的数组  Transform()，一个是“就地完成” TransformInPlace()。
//在主函数中，用 Kind() 方法检查了数据类型是不是 Slice，函数类型是不是 Func。
//检查函数的参数和返回类型是通过 verifyFuncSignature() 来完成的：NumIn()用来检查函数的“入参”；NumOut()  ：用来检查函数的“返回值”。
//如果需要新生成一个 Slice，会使用 reflect.MakeSlice() 来完成。

func uV21(){
	list := []string{"1", "2", "3", "4", "5", "6"}
	result := Transform(list, func(a string) string{
		return a +a +a
	})
	fmt.Println(result)
	//{"111","222","333","444","555","666"}
}

func uV22(){
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	TransformInPlace(list, func (a int) int {
		return a*3
	})
	fmt.Println(list)
	//{3, 6, 9, 12, 15, 18, 21, 24, 27}
}

func uV23(){

	var list = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
	}

	result := TransformInPlace(list, func(e Employee) Employee {
		e.Salary += 1000
		e.Age += 1
		return e
	})
	fmt.Println(result)

}




