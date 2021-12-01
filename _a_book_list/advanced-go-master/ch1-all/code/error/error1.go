package main

import (
	"encoding/json"
	"fmt"
	_range "github.com/rainbowmango/goexpertprogrammingsourcecode/range"
	"log"
	"syscall"
)

func mapFind() string {

	m :=make(map[string]string,5)
	if v,ok:=m["key"];ok{
		return v
	}
	return ""
}

func Errno()  {
	err := syscall.Chmod(":invalid path:", 0666)
	if err != nil {
		// 将 err 强制断 言为 syscall.Errno 错误类型来
		log.Fatal(err.(syscall.Errno))
	}
}

func main() {

}

// recover 用法
func ParseJson(input string) (s *Syntax,err error){
	defer func(){
		if p:=recover();p!=nil{
			err =fmt.Errorf("JSON: internal error: %v ",p)
		}
	}()
	// ..parse ...

}













