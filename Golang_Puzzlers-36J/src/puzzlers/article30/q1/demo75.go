package main

import (
	"fmt"
	"strings"
)

func main() {
	// 示例1。
	var builder1 strings.Builder
	builder1.WriteString("A Builder is used to efficiently build a string using Write methods.")
	fmt.Printf("The first output(%d):\n%q\n", builder1.Len(), builder1.String())
	fmt.Println()
	/**
	The first output(68):
	"A Builder is used to efficiently build a string using Write methods."

	*/
	builder1.WriteByte(' ')
	builder1.WriteString("It minimizes memory copying. The zero value is ready to use.")
	builder1.Write([]byte{'\n', '\n'})
	builder1.WriteString("Do not copy a non-zero Builder.")
	fmt.Printf("The second output(%d):\n\"%s\"\n", builder1.Len(), builder1.String())
	fmt.Println()
/**
  The second output(162):
  "A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use.

  Do not copy a non-zero Builder."
 */
	// 示例2。
	fmt.Println("Grow the builder ...")
	// TODO 主动扩容容量
	builder1.Grow(10)
	fmt.Printf("The length of contents in the builder is %d.\n", builder1.Len())
	fmt.Println()
/**
  Grow the builder ...
  The length of contents in the builder is 162.
 */
	// 示例3。
	fmt.Println("Reset the builder ...")
	builder1.Reset()
	fmt.Printf("The third output(%d):\n%q\n", builder1.Len(), builder1.String())
	/**
	Reset the builder ...
	The third output(0):
	""
	 */
}
