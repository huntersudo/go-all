package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func warpError(resp http.Response ,url string) (string,error){
	if _,err:=html.Parse(resp.Body);err !=nil{
		return nil,fmt.Errorf("parse %s as Html: %v",url,err)
	}

}

func lll()  {
	f,err:=os.Open("file.txt")
	if err!=nil{
		// 失败的情形，马上返回错误
	}
	// 正常的处理流程
// todo Go语言中大部分函数的代码结构几乎相同，首先是一系列的初始检查，用于防止错误发生，之后是函数 的实际逻辑。

}