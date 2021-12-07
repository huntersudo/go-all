package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

//  S1-词法分析  P3
//   token化，将代码转换为 go语言中定义的符号
//  "go/token" "go/scanner"  用于扫描源代码


// 模拟对源文件的扫描
func main() {
	// src is the input that we want to tokenize.
	src := []byte("cos(x) + 2i*sin(x) // Euler")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil , scanner.ScanComments)
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
// 打印出 token的位置、符号、及其字符串字面量
// 每个标识符和运算符都被特定的token的代替
/**
1:1	   IDENT	  "cos"
1:4	   (	      ""
1:5	   IDENT     "x"
1:6     )	      ""
1:8  	+	      ""
1:10	IMAG	  "2i"
1:12	*	      ""
1:13	IDENT	  "sin"
1:16	(	      ""
1:17	IDENT 	  "x"
1:18	)	      ""
1:20	;	      "\n"
1:20	COMMENT  	"// Euler"
 */