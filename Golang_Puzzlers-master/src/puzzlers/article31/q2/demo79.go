package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 示例1。
	var contents string
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The length of new buffer with contents %q: %d\n",
		contents, buffer1.Len())
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap())
	fmt.Println()
	/**
	The length of new buffer with contents "": 0
	The capacity of new buffer with contents "": 32
	 */

	contents = "12345"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()
	/**
	Write contents "12345" ...
	The length of buffer: 5
	The capacity of buffer: 32
	 */
	contents = "67"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()
/**
  Write contents "67" ...
  The length of buffer: 7
  The capacity of buffer: 32
 */
	contents = "89"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer1.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Print("\n\n")
/**
  Write contents "89" ...
  The length of buffer: 9
  The capacity of buffer: 32
 */




	// 示例2。
	contents = "abcdefghijk"
	buffer2 := bytes.NewBufferString(contents)
	fmt.Printf("The length of new buffer with contents %q: %d\n",
		contents, buffer2.Len())
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer2.Cap())
	fmt.Println()
/**
  The length of new buffer with contents "abcdefghijk": 11
  The capacity of new buffer with contents "abcdefghijk": 32
 */
	n := 10
	fmt.Printf("Grow the buffer with %d ...\n", n)
	buffer2.Grow(n)
	//todo 如果当前内容容器的容量的一半，仍然大于或等于其现有长度再加上另需的字节数的和，
	// 扩容代码就会复用现有的内容容器，并把容器中的未读内容拷贝到它的头部位置。
	fmt.Printf("The length of buffer: %d\n", buffer2.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer2.Cap())
	fmt.Print("\n\n")
/**
  Grow the buffer with 10 ...
  The length of buffer: 11
  The capacity of buffer: 32

 */
	// 示例3。
	var buffer3 bytes.Buffer
	fmt.Printf("The length of new buffer: %d\n", buffer3.Len())
	fmt.Printf("The capacity of new buffer: %d\n", buffer3.Cap())
	fmt.Println()
/**
  The length of new buffer: 0
  The capacity of new buffer: 0
 */
	contents = "xyz"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer3.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", buffer3.Len())
	fmt.Printf("The capacity of buffer: %d\n", buffer3.Cap())  //todo 64
	//todo 对于处在零值状态的Buffer值来说，如果第一次扩容时的另需字节数不大于64，那么该值就会基于一个预先定义好的、长度为64的字节数组来创建内容容器
	// 在这种情况下，这个内容容器的容量就是64。这样做的目的是为了让Buffer值在刚被真正使用的时候就可以快速地做好准备。
	/**
	Write contents "xyz" ...
	The length of buffer: 3
	The capacity of buffer: 64
	 */
}
