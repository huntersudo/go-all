package main

import "fmt"

// https://geekr.dev/posts/go-re-order-array

func reOrderArrayV1(arr []int) []int {
	var oddArr, evenArr []int
	for _, value := range arr {
		if value % 2 == 0 {
			evenArr = append(evenArr, value)
		} else {
			oddArr = append(oddArr, value)
		}
	}
	return append(oddArr, evenArr...)
}
// 非常简单，先声明两个数组切片，分别用于存储奇数和偶数，
//然后遍历待排序的数组切片，根据是否可以被 2 整除将切片数据分发到偶数和奇数切片，
//最后将偶数切片数据追加到奇数切片之后作为新的切片返回。
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("排序前:", arr)
	fmt.Println("排序后:", reOrderArrayV1(arr))
}
// 排序前: [0 1 2 3 4 5 6 7 8 9]
//排序后: [1 3 5 7 9 0 2 4 6 8]