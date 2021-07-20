package main

import "fmt"

func main() {
	// 示例1。
	s6 := make([]int, 0)
	fmt.Printf("The capacity of s6: %d\n", cap(s6))
	for i := 1; i <= 5; i++ {
		s6 = append(s6, i)
		fmt.Printf("s6(%d): len: %d, cap: %d\n", i, len(s6), cap(s6))
	}
	fmt.Println()

	/**
	The capacity of s6: 0
	s6(1): len: 1, cap: 1
	s6(2): len: 2, cap: 2
	s6(3): len: 3, cap: 4
	s6(4): len: 4, cap: 4
	s6(5): len: 5, cap: 8
	 新切片的容量（以下简称新容量）将会是原切片容量（以下简称原容量）的 2
	*/

	// 示例2。
	s7 := make([]int, 1024)
	fmt.Printf("The capacity of s7: %d\n", cap(s7))
	s7e1 := append(s7, make([]int, 200)...)
	fmt.Printf("s7e1: len: %d, cap: %d\n", len(s7e1), cap(s7e1))
	s7e2 := append(s7, make([]int, 400)...)
	fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2))
	s7e3 := append(s7, make([]int, 600)...)
	fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3))
	fmt.Println()
/**
  The capacity of s7: 1024
  s7e1: len: 1224, cap: 1280
  s7e2: len: 1424, cap: 1696
  s7e3: len: 1624, cap: 2048
  当原切片的长度（以下简称原长度）大于或等于1024时，Go 语言将会以原容量的1.25倍作为新容量的基准（以下新容量基准）
 */
	// 示例3。
	s8 := make([]int, 10)
	fmt.Printf("The capacity of s8: %d\n", cap(s8))
	s8a := append(s8, make([]int, 11)...)
	fmt.Printf("s8a: len: %d, cap: %d\n", len(s8a), cap(s8a))
	s8b := append(s8a, make([]int, 23)...)
	fmt.Printf("s8b: len: %d, cap: %d\n", len(s8b), cap(s8b))
	s8c := append(s8b, make([]int, 45)...)
	fmt.Printf("s8c: len: %d, cap: %d\n", len(s8c), cap(s8c))

	/**
	The capacity of s8: 10
	s8a: len: 21, cap: 22
	s8b: len: 44, cap: 44
	s8c: len: 89, cap: 96
	如果我们一次追加的元素过多，以至于使新长度比原容量的 2 倍还要大，那么新容量就会以新长度为基准
	 */
}
