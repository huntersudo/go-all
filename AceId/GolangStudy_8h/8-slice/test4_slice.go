package main

import "fmt"

// 容量
func main() {
	  //len cap
	var numbers = make([]int, 3, 5)
     // todo len = 3, cap = 5, slice = [0 0 0]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//todo 向numbers切片追加一个元素1, numbers len = 4， [0,0,0,1], cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//todo 向numbers切片追加一个元素2, numbers len = 5， [0,0,0,1,2], cap = 5
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//todo 向一个容量cap已经满的slice 追加元素， len = 6, cap = 10, slice = [0 0 0 1 2 3] 扩容  cap*2
	numbers = append(numbers, 3)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("-=-------")

	var numbers2 = make([]int, 3)
	// len = 3, cap = 3, slice = [0 0 0]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	// len = 4, cap = 6, slice = [0 0 0 1]   todo cap*2
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
