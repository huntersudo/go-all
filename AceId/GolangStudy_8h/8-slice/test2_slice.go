package main

import "fmt"

func printArray(myArray []int) {
	//todo 引用传递
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main() {
	myArray := []int{1,2,3,4} // 动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)
     // todo myArray type is []int 相比较 array

	printArray(myArray)   // todo 引用传递

	fmt.Println(" ==== ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}