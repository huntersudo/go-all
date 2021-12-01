package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 示例1。
	var buffer1 bytes.Buffer
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()
/**
  Write contents "Simple byte buffer for marshaling data." ...
  The length of buffer: 39
  The capacity of buffer: 64
 */
	// 示例2。
	p1 := make([]byte, 7)
	n, _ := buffer1.Read(p1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	/**
	7 bytes were read. (call Read)
	todo ，Buffer值的长度是未读内容的长度，而不是已存内容的总长度。
	The length of buffer: 32
	The capacity of buffer: 64
	 */
}
