package main

func main() {
	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
}
/**
panic: runtime error: index out of range [5] with length 5

goroutine 1 [running]:
main.main()
	D:/workspaces/Golang_Puzzlers/src/puzzlers/article19/q0/demo47.go:5 +0x1b

Process finished with exit code 2

 */