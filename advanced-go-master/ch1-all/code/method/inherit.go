package main

import (
	"fmt"
	"image/color"
)

type point struct {
	X, Y float64
}

func (p *point) setX(x float64) {
	p.X = x
}

func (p *point) setY(y float64) {
	p.Y = y
}

func (p point) print() {
	fmt.Printf("p.X = %f, p.Y = %f\n", p.X, p.Y)
}
// 组合 实现  继承
type coloredPoint struct {
	point
	clr color.RGBA
}
// 我们不仅可以继承匿名成员的内部成员，而且可以继承匿名成员类型所对应的方法
// todo 。不过这种方式继承的 方法并不能实现C++中虚函数的多态特性。所有继承来的方法的接收者参数依然是那个匿名成员本身， 而不是当前的变量。
func (clr coloredPoint) print() {
	fmt.Printf("I'm coloredPoint::print\n")
	fmt.Printf("clr.X = %f, clr.Y = %f\n", clr.X, clr.Y)
	clr.point.print()
}

func main() {
	fmt.Println("Method Inherit")
	var cp coloredPoint
	cp.setX(1.01)
	cp.setY(2.02)
	cp.print()
}
