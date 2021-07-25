package main

import "fmt"

func printArray1(myArray [4]int) {
	//值拷贝

	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}

	myArray[0] = 111
}


func main() {
	// 固定长度的数组
	var myArray1 [10]int
	//for i := 0; i < 10; i++ {
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])	
	}

	myArray2 := [10]int{1,2,3,4}
	myArray3 := [4]int{11,22,33,44}


	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)
  // myArray1 types = [10]int   todo  数组传参 必须 包含 [x]类型
	//myArray2 types = [10]int
	//myArray3 types = [4]int

	printArray1(myArray3)  // 值拷贝
	//printArray1(myArray2)  // 值拷贝 todo  Cannot use 'myArray2' (type [10]int) as the type [4]int

	fmt.Println(" ------ ")
	for index, value := range myArray3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
}