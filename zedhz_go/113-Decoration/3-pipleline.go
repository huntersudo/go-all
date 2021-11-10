package main

import "net/http"

//在使用上，需要对函数一层层地套起来，看上去好像不是很好看，如果需要修饰器比较多的话，代码就会比较难看了。不过，我们可以重构一下

//todo 重构时，我们需要先写一个工具函数，用来遍历并调用各个修饰器：

type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		h = d(h)
	}
	return h
}
// todo 这样的代码是不是更易读了一些？Pipeline 的功能也就出来了
func use1(){

	http.HandleFunc("/v4/hello", Handler(hello,WithServerHeader, WithBasicAuth, WithDebugLog))
}