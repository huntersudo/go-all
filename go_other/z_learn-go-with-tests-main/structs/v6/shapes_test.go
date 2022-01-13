package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
// 列表驱动测试
func TestArea(t *testing.T) {
// 创建了一个匿名的结构体
	// 开发人员能方便的引入一个新的几何形状，只需实现 Area 方法并把新的类型加到测试用例 中。
	//另外发现 Area 方法有错误，我们可以在修复这个错误之前非常容易的添加新的测试用例
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
