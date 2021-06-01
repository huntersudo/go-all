package main

import "fmt"

func main() {

	sliceA := make([]int, 5, 10) //length = 5; capacity = 10
	sliceB := sliceA[0:5]        //length = 5; capacity = 10
	sliceC := sliceA[0:5:5]      //length = 5; capacity = 5

	fmt.Println("len(sliceA) = ", len(sliceA))
	fmt.Println("cap(sliceA) = ", cap(sliceA))
	// 5
	fmt.Println("len(sliceB) = ", len(sliceB))
	// 10
	fmt.Println("cap(sliceB) = ", cap(sliceB))
     // 5
	fmt.Println("len(sliceC) = ", len(sliceC))
	// 5
	fmt.Println("cap(sliceC) = ", cap(sliceC))
}