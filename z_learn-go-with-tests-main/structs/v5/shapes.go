package main

import "math"

// Shape is implemented by anything that can tell us its Area.
type Shape interface {
	Area() float64
}
// 加了这个代码后测试运行通过了。
// go语言的接口 是隐式实现的, todo 实现了，只做增加，不做修改的想法

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
