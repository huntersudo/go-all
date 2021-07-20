package q3

import "testing"
/**
$ go test -bench=. -run=^$ puzzlers/src/puzzlers/article20/q3
goos: windows
goarch: amd64
pkg: puzzlers/src/puzzlers/article20/q3
BenchmarkGetPrimes-12             595237              1887 ns/op
PASS
ok      puzzlers/src/puzzlers/article20/q3      1.564s

如果运行go test命令的时候不加-run标记，那么就会使它执行被测代码包中的所有功能测试函数。

加入标记-cpu来设置一个最大 P 数量的列表

$ go test -bench=. -run=^$ -cpu=6 puzzlers/src/puzzlers/article20/q3
goos: windows
goarch: amd64
pkg: puzzlers/src/puzzlers/article20/q3
BenchmarkGetPrimes-6      563094              1881 ns/op
PASS
ok      puzzlers/src/puzzlers/article20/q3      1.461s

*/

/**
go test命令在执行性能测试函数的时候会给它一个正整数，若该测试函数的唯一参数的名称为b，则该正整数就由b.N代表。
一个会迭代b.N次的循环中调用了GetPrimes函数，并给予它参数值1000。go test命令会先尝试把b.N设置为1，然后执行测试函数。
如果测试函数的执行时间没有超过上限，此上限默认为 1 秒，那么命令就会改大b.N的值，然后再次执行测试函数，如此往复，直到这个时间大于或等于上限为止.
我们就说这是命令此次对该测试函数的最后一次执行。这时的b.N的值就会被包含在测试结果中，也就是上述测试结果中的595237

它指的是被测函数的执行次数[GetPrimes]，而不是性能测试函数[BenchmarkGetPrimes]的执行次数。

 */
func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}


/**
如果是性能测试的话，可以在执行 go test 命令时追加 -benchtime 标记。这实际上是下一篇文章的思考题。

另外，-timeout 标记可用于设定当次测试运行的总时长，一旦超过这个时长就panic。

具体用法可以参考：https://golang.org/cmd/go/#hdr-Testing_flags 。

 */

/**
go如何写出可测试的代码
1. 隐藏不该暴露的。测试代码相当于使用你程序的一方，不要让它能够影响甚至破坏你的程序。
2. 暴露出来的API应该对使用者友好。换句话会所，如果在写测试代码的时候，你发现API调用起来不那么顺畅，那么就需要考虑程序的API是否得当了。
3. 保持API在风格上的一致性。风格一致的程序有助于使用方快速运用和理解程序。

这3点是基本的原则了，其他的细节上的东西说到底都是为了做到这3点的。测试代码就是对生产代码的前期验证

 */