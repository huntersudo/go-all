package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// S3-抽象语法树构建  p6
// ast 是一种常见的树状结构的中间态
// import\type\const\func 都是一个根节点 ，在根节点下包含当前声明的子节点
// 核心  gc/noder.go

func main(){
	src := `
package main
func main() {
	a := b + c(12)
}
`
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}