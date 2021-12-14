package main

import "fmt"

var g = 100


func f() (r int) {
	defer func() {
		g = 200
	}()

	fmt.Printf("f: g = %d\n", g)

	return g  // 先把g保存在栈上--执行defer---函数返回
}

func main() {
	i := f()
	fmt.Printf("main: i = %d, g = %d\n", i, g)

	i = f2()
	fmt.Printf("main g2: i = %d, g = %d\n", i, g)

}
//f: g = 100
//main: i = 100, g = 200

//f: g = 200
//main g2: i = 300, g = 200

func f2() (r int) {
	defer func() {
		g = 200
	}()

	fmt.Printf("f: g = %d\n", g)

	return g2()  //todo 先把g2保存在栈上（g2是函数则先执行）--执行defer---函数返回
}

func  g2() int  {
	 g=300
	 return g
}

