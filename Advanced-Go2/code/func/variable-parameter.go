package main

import "fmt"

// pack when call in, not unpack when call out
func foo1(parm ...int) {
	print(parm)
}

func bar1(parm []interface{}) {
	// not unpack when call out
	print(parm)
	// pack when call out
	print(parm...)
}

// pack when call in, unpack when call out.
func print(parm ...interface{}) {
	fmt.Println(parm...)
}

func main() {
	fmt.Println("Variable parmmeter Function")

	foo1(1, 2)
	fooParm := []int{1, 2}
	foo1(fooParm...)
	print(1, 2)
//[1 2]
//[1 2]
//1 2
	interfaceParm := []interface{}{123, "abc"}
	print(interfaceParm)
	print(interfaceParm...)

	bar1(interfaceParm)
//[123 abc]
//123 abc

//[123 abc]
//123 abc
}
