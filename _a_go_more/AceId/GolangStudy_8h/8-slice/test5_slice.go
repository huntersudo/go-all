package main

import "fmt"

// 截取
func main() {
	s := []int{1, 2, 3} //len = 3, cap = 3, [1,2,3]

	//todo [0, 2) 左闭右开
	s1 := s[0:2] // [1, 2]

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s) // [100 2 3]
	fmt.Println(s1)  // [100 2] todo 底层数组是一样的

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2 = [0,0,0]

	//将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)  // [100 2 3]

}
