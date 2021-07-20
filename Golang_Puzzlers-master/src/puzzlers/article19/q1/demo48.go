package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	caller1()
	fmt.Println("Exit function main.")
}

func caller1() {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Exit function caller1.")
}

func caller2() {
	fmt.Println("Enter function caller2.")
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
	fmt.Println("Exit function caller2.")
}
/**
Enter function main.
Enter function caller1.
Enter function caller2.
panic: runtime error: index out of range [5] with length 5

goroutine 1 [running]:
main.caller2()
	D:/workspaces/Golang_Puzzlers/src/puzzlers/article19/q1/demo48.go:22 +0x89
main.caller1()
	D:/workspaces/Golang_Puzzlers/src/puzzlers/article19/q1/demo48.go:15 +0x8a
main.main()
	D:/workspaces/Golang_Puzzlers/src/puzzlers/article19/q1/demo48.go:9 +0x8a
 */