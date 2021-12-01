package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 示例1。
	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap())
	fmt.Println()
// The capacity of new buffer with contents "ab": 8

	unreadBytes := buffer1.Bytes()  //todo 基于内容容器的切片返回给调用方，存在内容泄露的风险
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Println()
	// The unread bytes of the buffer: [97 98]

	contents = "cdefg"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()
/**
  Write contents "cdefg" ...
  The capacity of buffer: 8
 */

	// 只要扩充一下之前拿到的未读字节切片unreadBytes，
	//todo  就可以用它来读取甚至修改buffer中的后续内容。
	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Println()
// The unread bytes of the buffer: [97 98 99 100 101 102 103 0]

	value := byte('X')
	fmt.Printf("Set a byte in the unread bytes to %v ...\n", value)
	unreadBytes[len(unreadBytes)-2] = value
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer1.Bytes())
	fmt.Println()
/**
  Set a byte in the unread bytes to 88 ...
  The unread bytes of the buffer: [97 98 99 100 101 102 88]
 */

	// 不过，在buffer的内容容器真正扩容之后就无法这么做了。
	contents = "hijklmn"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer1.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())
	fmt.Println()
/**
  Write contents "hijklmn" ...
  The capacity of buffer: 23
 */

	unreadBytes = unreadBytes[:cap(unreadBytes)]
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes)
	fmt.Print("\n\n")
/**
  The unread bytes of the buffer: [97 98 99 100 101 102 88 0]
 */
	// 示例2。
	// Next方法返回的后续字节切片也存在相同的问题。
	contents = "12"
	buffer2 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer2.Cap())
	fmt.Println()
	/**
	The capacity of new buffer with contents "12": 8
	 */

	nextBytes := buffer2.Next(2)
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
	fmt.Println()
/**
  The next bytes of the buffer: [49 50]
 */
	contents = "34567"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer2.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer2.Cap())
	fmt.Println()
/**
  Write contents "34567" ...
  The capacity of buffer: 8
 */
	// 只要扩充一下之前拿到的后续字节切片nextBytes，
	// 就可以用它来读取甚至修改buffer中的后续内容。
	nextBytes = nextBytes[:cap(nextBytes)]
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
	fmt.Println()
/**
  The next bytes of the buffer: [49 50 51 52 53 54 55 0]
 */
	value = byte('X')
	fmt.Printf("Set a byte in the next bytes to %v ...\n", value)
	nextBytes[len(nextBytes)-2] = value
	fmt.Printf("The unread bytes of the buffer: %v\n", buffer2.Bytes())
	fmt.Println()
/**
  Set a byte in the next bytes to 88 ...
  The unread bytes of the buffer: [51 52 53 54 88]
 */
	// 不过，在buffer的内容容器真正扩容之后就无法这么做了。
	contents = "89101112"
	fmt.Printf("Write contents %q ...\n", contents)
	buffer2.WriteString(contents)
	fmt.Printf("The capacity of buffer: %d\n", buffer2.Cap())
	fmt.Println()
/**
  Write contents "89101112" ...
  The capacity of buffer: 24
 */
	nextBytes = nextBytes[:cap(nextBytes)]
	fmt.Printf("The next bytes of the buffer: %v\n", nextBytes)
	/**
	The next bytes of the buffer: [49 50 51 52 53 54 88 0]
	 */
}
