package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		// 这里的 f 对应 float64 ， .2 表示输出 2 位小数。
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
