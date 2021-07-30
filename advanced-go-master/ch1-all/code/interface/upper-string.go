package main

import (
	"fmt"
	"os"
	"strings"
)

/*
"Fprintf" function signature:
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

"a" should implement "Stringer" interface, which describe the type as a string.
type Stringer interface {
    String() string
}
*/

type upperString string
 // 如果满足了 fmt.Stringer 接口，则默认使用对象的 String 方法返回的结果打印：
func (s upperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Println("Upper String")
	fmt.Fprintln(os.Stdout, upperString("hello, world"))
}
