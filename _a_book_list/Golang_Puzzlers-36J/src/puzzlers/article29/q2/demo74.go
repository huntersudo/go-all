package main

import "fmt"

func main() {
	str := "Go爱好者"
	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}
	/**
	0: 'G' [47]
	1: 'o' [6f]
	2: '爱' [e7 88 b1]
	5: '好' [e5 a5 bd]
	8: '者' [e8 80 85]
	 */
}
