package main

import (
	"fmt"
	"reflect"
	"strings"
)

//todo 下面，我们来看一下，一个非常简单的、不做任何类型检查的泛型的 Map 函数怎么写。

func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}
// 我来简单解释下这段代码。
//首先，我们通过 reflect.ValueOf() 获得 interface{} 的值，其中一个是数据 vdata，另一个是函数 vfn。
//todo 然后，通过 vfn.Call() 方法调用函数，通过 []refelct.Value{vdata.Index(i)}获得数据。

//Go 语言中的反射的语法有点令人费解，不过，简单看一下手册，还是能够读懂的。反射不是这节课的重点，我就不讲了。
//如果你还不太懂这些基础知识，课下可以学习下相关的教程。于是，我们就可以有下面的代码——不同类型的数据可以使用相同逻辑的Map()代码。

// 于是，我们就可以有下面的代码——不同类型的数据可以使用相同逻辑的Map()代码。

func useG1(){

	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}

	squared_arr := Map(nums,square)
	fmt.Println(squared_arr)
	//[1 4 9 16]

	upcase := func(s string) string {
		return strings.ToUpper(s)
	}
	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := Map(strs, upcase);
	fmt.Println(upstrs)
	//[HAO CHEN MEGAEASE]
}

/**
但是，因为反射是运行时的事，所以，如果类型出问题的话，就会有运行时的错误。比如：

x := Map(5, 5)
fmt.Println(x)
代码可以很轻松地编译通过，但是在运行时却出问题了，而且还是 panic 错误……

panic: reflect: call of reflect.Value.Len on int Value

goroutine 1 [running]:
reflect.Value.Len(0x10b5240, 0x10eeb58, 0x82, 0x10716bc)
        /usr/local/Cellar/go/1.15.3/libexec/src/reflect/value.go:1162 +0x185
main.Map(0x10b5240, 0x10eeb58, 0x10b5240, 0x10eeb60, 0x1, 0x14, 0x0)
        /Users/chenhao/.../map.go:12 +0x16b
main.main()
        /Users/chenhao/.../map.go:42 +0x465
exit status 2

 */