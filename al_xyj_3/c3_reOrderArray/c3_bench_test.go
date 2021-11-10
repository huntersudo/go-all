package main

import (
"testing"
)

func BenchmarkReOrderArrayV1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		reOrderArrayV1(arr) // run reOrderArrayV1(arr) b.N times
	}
}
func BenchmarkReOrderArrayV2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		reOrderArrayV2(arr, isEven) // run reOrderArrayV2(arr, isEven) b.N times
	}
}