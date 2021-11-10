
package main

import "fmt"

/**
我们动用了一个高阶函数 decorator()，在调用的时候，先把 Hello() 函数传进去，
然后会返回一个匿名函数。这个匿名函数中除了运行了自己的代码，也调用了被传入的 Hello() 函数
 */
func decorator(f func(s string)) func(s string) {

	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

func main() {
	decorator(Hello)("Hello, World!")

    // 可读性更好
	hello := decorator(Hello)
	hello("Hello")
}
//Started
//Hello, World!
//Done
