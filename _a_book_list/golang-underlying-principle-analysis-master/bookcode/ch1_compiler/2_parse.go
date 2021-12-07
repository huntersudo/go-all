package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// S2-语法解析, P5
// 对符号化的go文件进行解析，自上而下的递归下降算法，完成无需回溯的语法扫描
// syntax/nodes.go  syntax/parser.go 中
// 源文件中每一种声明都有一种语法，在go语言规范中都有定义，
// 参考： [2] https://golang.org/ref/spec#add_op
// 每一种声明语法或者表达式都有对应的结构体  ast/ast.go
// AssignStmt struct {
//		Lhs    []Expr
//		TokPos token.Pos   // position of Tok
//		Tok    token.Token // assignment token, DEFINE
//		Rhs    []Expr
//	}
//语法解析： 将 语义存储到 对应的结构体中
//


func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src1, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
		return
	}

	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

const src = `package pkgname
func main() {
	a := b + c(12)
}
`
const src1 = `
package main
import "fmt"
const word="GO"
type STR string
var hello STR ="i love" + word
func main() {
	fmt.Println(hello)
}
`

/**

   0  *ast.BlockStmt {
   1  .  Lbrace: 29
   2  .  List: []ast.Stmt (len = 1) {
   3  .  .  0: *ast.AssignStmt {
   4  .  .  .  Lhs: []ast.Expr (len = 1) {
   5  .  .  .  .  0: *ast.Ident {
   6  .  .  .  .  .  NamePos: 32
   7  .  .  .  .  .  Name: "a"
   8  .  .  .  .  .  Obj: *ast.Object {
   9  .  .  .  .  .  .  Kind: var
  10  .  .  .  .  .  .  Name: "a"
  11  .  .  .  .  .  .  Decl: *(obj @ 3)
  12  .  .  .  .  .  }
  13  .  .  .  .  }
  14  .  .  .  }
  15  .  .  .  TokPos: 34
  16  .  .  .  Tok: :=
  17  .  .  .  Rhs: []ast.Expr (len = 1) {
  18  .  .  .  .  0: *ast.BinaryExpr {
  19  .  .  .  .  .  X: *ast.Ident {
  20  .  .  .  .  .  .  NamePos: 37
  21  .  .  .  .  .  .  Name: "b"
  22  .  .  .  .  .  }
  23  .  .  .  .  .  OpPos: 39
  24  .  .  .  .  .  Op: +
  25  .  .  .  .  .  Y: *ast.CallExpr {
  26  .  .  .  .  .  .  Fun: *ast.Ident {
  27  .  .  .  .  .  .  .  NamePos: 41
  28  .  .  .  .  .  .  .  Name: "c"
  29  .  .  .  .  .  .  }
  30  .  .  .  .  .  .  Lparen: 42
  31  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
  32  .  .  .  .  .  .  .  0: *ast.BasicLit {
  33  .  .  .  .  .  .  .  .  ValuePos: 43
  34  .  .  .  .  .  .  .  .  Kind: INT
  35  .  .  .  .  .  .  .  .  Value: "12"
  36  .  .  .  .  .  .  .  }
  37  .  .  .  .  .  .  }
  38  .  .  .  .  .  .  Ellipsis: 0
  39  .  .  .  .  .  .  Rparen: 45
  40  .  .  .  .  .  }
  41  .  .  .  .  }
  42  .  .  .  }
  43  .  .  }
  44  .  }
  45  .  Rbrace: 47
  46  }
 */