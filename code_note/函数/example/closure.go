//闭包函数：函数+引用

package main

import "fmt"

//函数test()   返回值 func (int, int) (int, int)
// 返回的是一个匿名函数
// 这个匿名函数引用了函数外的 i j
func test1() func(int, int) (int, int) {
	i := 1
	j := 2
	return func(x, y int) (int, int) {
		i += x
		j += y
		return i, j
	}

}

func main() {
	t := test1()

	fmt.Println(t(1, 2)) // 2 4
	//保存
	fmt.Println(t(1, 2)) //3 6

	fmt.Println(t(1, 2)) // 4 8
}
