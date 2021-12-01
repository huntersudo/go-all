package main

import "fmt"

func main() {
	// 示例1。
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1))
	// a1: [1 2 3 4 5 6 7] (len: 7, cap: 7)
	s9 := a1[1:4]
	//s9[0] = 1
	fmt.Printf("s9: %v (len: %d, cap: %d)\n",
		s9, len(s9), cap(s9))
	// s9: [2 3 4] (len: 3, cap: 6)
	// 6=7-1,底层数组的长度 7 - 切片表达式的起始索引 1

	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)
		fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n",
			i, s9, len(s9), cap(s9))
	}
	/**
	s9(1): [2 3 4 1] (len: 4, cap: 6)
	s9(2): [2 3 4 1 2] (len: 5, cap: 6)
	s9(3): [2 3 4 1 2 3] (len: 6, cap: 6)
	s9(4): [2 3 4 1 2 3 4] (len: 7, cap: 12)
	s9(5): [2 3 4 1 2 3 4 5] (len: 8, cap: 12)
	 */

	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1))
	// a1: [1 2 3 4 1 2 3] (len: 7, cap: 7)
	// 一个切片的底层数组永远不会被替换
	// 扩容的时候，会生成新的底层数组、新的切片
	fmt.Println()

}
