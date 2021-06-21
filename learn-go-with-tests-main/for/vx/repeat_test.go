package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

//基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
//代码运行的次数应该不影响你，框架将决定什么是「好」的值，以便让你获得一些得体的结果。
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
// goos: windows
//goarch: amd64
//pkg: github.com/quii/learn-go-with-tests/for/vx
//BenchmarkRepeat-12    	 9705102	       123 ns/op
//PASS
