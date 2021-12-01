package main

import (
	"fmt"
	"io"
	"os"
)

// Go语言中，对于基础类型（非接口类型）不支持隐式的转换
// Go语言对基础类型的类型一致性要求可谓是非常的严格
// 但是Go语言对于接口类型的转换则非常的 灵活
func main() {
	fmt.Println("Interface Implicit Conversion")

	f, _ := os.Open("./conversion.go")
	fmt.Printf("(*os.File) type: %T\n", f)  // (*os.File) type: *os.File

	// 隐式转换, *os.File 满足 io.ReadCloser 接口
	var readCloser io.ReadCloser = f
	fmt.Printf("readCloser type: %T\n", readCloser)  // readCloser type: *os.File

	// 隐式转换, io.ReadCloser 满足 io.Reader 接
	var reader io.Reader = readCloser
	fmt.Printf("reader type: %T\n", reader)  // reader type: *os.File

	// 隐式转换, io.ReadCloser 满足 io.Closer 接
	var closer io.Closer = readCloser
	fmt.Printf("closer type: %T\n", closer)  // closer type: *os.File

	//todo  // 显式转换, io.Closer 不满足 io.Reader 接口
	// Use Type Assertion to access to underlying concrete value
	var reader2 io.Reader = closer.(io.Reader)
	fmt.Printf("reader2 type: %T\n", reader2)   // 	fmt.Printf("reader2 type: %T\n", reader2)

	f.Close()

	numInt := 1
	fmt.Printf("numInt type: %T\n", numInt)  // numInt type: int

	var numInt32 int32 = int32(numInt)
	fmt.Printf("type: %T\n", numInt32)  // type: int32

	var numInt64 int64 = int64(numInt)
	fmt.Printf("type: %T\n", numInt64)  // type: int64

}
