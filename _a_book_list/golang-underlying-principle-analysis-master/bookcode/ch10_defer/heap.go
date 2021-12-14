package main

import "fmt"

func main(){
	for i:=0;i<5;i++{
		defer fmt.Println("defer func",i)
	}
}
// defer func 4
//defer func 3
//defer func 2
//defer func 1
//defer func 0