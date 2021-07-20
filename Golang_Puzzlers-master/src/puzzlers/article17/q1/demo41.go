package main

import "fmt"

func main() {
	// 示例1。
	//切片
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	// 只有一个迭代变量的情况意味着什么呢？这意味着，该迭代变量只会代表当次迭代对应的元素值的索引值。
	// 当i的值等于3的时候，与之对应的是切片中的第 4 个元素值4
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
	fmt.Println()
// [1 2 3 7 5 6]


	// 示例2。
	// 数组
	// range表达式只会在for语句开始执行时被求值一次，无论后边会有多少次迭代；
	// range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的副本而不是原值。
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	// 5
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
			// =1+6
		} else {
			numbers2[i+1] += e
			// =1+2
			// =2+3
			// =3+4
			// =4+5
			// =5+6
		}
	}
	fmt.Println(numbers2)
	fmt.Println()
// [7 3 5 7 9 11]

	// 示例3。TODO 和数组比较着来看
	// range 切片的副本 ，但是实际操作的还是 切片的底层数组，也就是会修改其值
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers3) - 1
	// 5
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
			// =21+1 => {22,3,6,10,15,21}
		} else {
			numbers3[i+1] += e
			// =1+2  ,=>{1,3,3,4,5,6}
			// =3+3  ,=>{1,3,6,4,5,6}
			// =6+4  ,=>{1,3,6,10,5,6}
			// =10+5 ,=>{1,3,6,10,15,6}
			// =15+6 ,=>{1,3,6,10,15,21}
		}
	}
	fmt.Println(numbers3)

	// [22 3 6 10 15 21]
}
