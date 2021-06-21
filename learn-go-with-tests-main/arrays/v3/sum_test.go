package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		// mySlice := []int{1,2,3} 而不是 mySlice := [3]int{1,2,3}
		// 用 切片类型，它可以接收不同大小的切片集合。语法上和数组非常相似，只是在声明的时候 不指定长度：
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
