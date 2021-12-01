package main

import (
	"fmt"
	"reflect"
)

// 不过，对于 Go 的修饰器模式，还有一个小问题，那就是好像无法做到泛型。
// todo 比如上面那个计算时间的函数，其代码耦合了需要被修饰的函数的接口类型，无法做到非常通用。如果这个问题解决不了，那么，这个修饰器模式还是有点不好用的
//  因为 Go 语言不像 Python 和 Java，Python 是动态语言，而 Java 有语言虚拟机，所以它们可以实现一些比较“变态”的事。
//  但是，Go 语言是一个静态的语言，这就意味着类型需要在编译时就搞定，否则无法编译。
//  不过，Go 语言支持的最大的泛型是 interface{}  ，还有比较简单的 Reflection 机制，在上面做做文章，应该还是可以搞定的。


func Decorator(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}
// todo 段代码动用了 reflect.MakeFunc() 函数，创造了一个新的函数，其中的 targetFunc.Call(in) 调用了被修饰的函数。
// 关于 Go 语言的反射机制，你可以阅读下官方文章The Laws of Reflection，我就不多说了。
// 这个 Decorator() 需要两个参数：
// 第一个是出参 decoPtr ，就是完成修饰后的函数；
// 第二个是入参 fn ，就是需要修饰的函数。
// todo 这样写是不是有些“傻”？的确是的。不过，这是我个人在 Go 语言里所能写出来的最好的代码了。如果你知道更多优雅的写法，请你要一定告诉我！


func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}



func use(){
	type MyFoo func(int, int, int) int
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1, 2, 3)

// 你会发现，使用 Decorator() 时，还需要先声明一个函数签名，感觉好傻啊，一点都不泛型，不是吗？如果你不想声明函数签名，就可以这样：
	mybar := bar
	Decorator(&mybar, bar)
	mybar("hello,", "world!")
}
//todo  好吧，看上去不是那么漂亮，但是 it works。看样子 Go 语言目前本身的特性无法做成像 Java 或 Python 那样，对此，我们只能期待 Go 语言多放“糖”了！

