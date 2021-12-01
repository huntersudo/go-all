package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one_new", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])

     //  类型断言”表达式
	value, ok := interface{}(container).([]string)
	if ok {
		fmt.Println(value)
	}
	/**
	在 Go 语言中，interface{}代表空接口，任何类型都是它的实现类型。
	我在下个模块，会再讲接口及其实现类型的问题。现在你只要知道，任何类型的值都可以很方便地被转换成空接口的值就行了。
	这里的具体语法是interface{}(x)，例如前面展示的interface{}(container)。你可能
	 */
}

/**

The element is "one_new".
 */