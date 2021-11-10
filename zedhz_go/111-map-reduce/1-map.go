package main

import (
	"fmt"
	"strings"
)

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}
// 整个 Map 函数的运行逻辑都很相似，函数体都是在遍历第一个参数的数组，然后，调用第二个参数的函数，把它的值组合成另一个数组返回。

func use1(){

	var list = []string{"Hao", "Chen", "MegaEase"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x)
	//["HAO", "CHEN", "MEGAEASE"]

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y)
	//[3, 4, 8]
}
