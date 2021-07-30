package main

import "fmt"

func func1()  {
	fmt.Println("func1()")
}

func func2()  {
	fmt.Println("func2()")
}

func func3()  {
	fmt.Println("func3()")
}

func main() {
	defer func1()
	defer func2()
	defer func3()
}
//func3()
//func2()
//func1()