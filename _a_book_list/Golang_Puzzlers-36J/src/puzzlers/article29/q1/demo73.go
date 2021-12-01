package main

import (
	"fmt"
)

func main() {
	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str)) //  runes(char): ['G' 'o' '爱' '好' '者']
	fmt.Printf("  => runes(hex): %x\n", []rune(str)) //  runes(hex): [47 6f 7231 597d 8005]
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str)) // bytes(hex): [47 6f e7 88 b1 e5 a5 bd e8 80 85]
}
