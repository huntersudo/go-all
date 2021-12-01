package main

import "fmt"

/**
defer语句用于延迟函数的调用，每次defer都会把一个函数压入栈中，函数返回前再把延迟的函数取出并执行。

为了方便描述，我们把创建defer的函数称为主函数，defer语句后面的函数称为延迟函数。

延迟函数可能有输入参数，这些参数可能来源于定义defer的函数，延迟函数也可能引用主函数用于返回的变量，
也 就是说延迟函数可能会影响主函数的一些行为，这些场景下，如果不了解defer的规则很容易出错。
其实官方说明的defer的三个原则很清楚，本节试图汇总defer的使用场景并做简单说明。
*/

/**
题目1
*/
func deferFuncParameter() {

	var aInt = 1
	defer fmt.Println(aInt) //1
	// 延迟函数fmt.Println(aInt)的参数在defer语句出现时就已经确定了，所以无论后面如何修改 aInt变量都不会影响延迟函数
	aInt = 2
	return
}

func main() {
	//deferFuncParameter()
	//deferFuncParameter2()
	fmt.Println(deferFuncReturn())
}

/**
题目二

*/
func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}
func deferFuncParameter2() {
	var aArray = [3]int{1, 2, 3}

	defer printArray(&aArray) // 10 2 3
    // 延迟函数printArray()的参数在defer语句出现时就已经确定了，即数组的地址
    // 所以对数组的最终修改值会被打印出来
	aArray[0] = 10
	return
}

/**
题目三
函数拥有一个具名返回值result，函数内部声明一个变量i，defer指定一个延迟函数，最后返回变量i。 延迟函数中递增result。
// 返回 2
函数的return语句并不是原子的，实际执行分为设置返回值—>ret，defer语句实际执行在 返回前，
即拥有defer的函数返回过程是：设置返回值—>执行defer—>ret。
所以return语句先把result设置为i 的值，即1，defer语句中又把result递增1，所以最终返回2。
 */
func deferFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}