package q1

import (
	"testing"
)

var expectedPrimes = []int{
	2, 3, 5, 7, 11, 13, 17, 19,
	23, 29, 31, 37, 41, 43, 47, 53,
	59, 61, 67, 71, 73, 79, 83, 89,
	97, 101, 103, 107, 109, 113, 127, 131,
	137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223,
	227, 229, 233, 239, 241, 251, 257, 263,
	269, 271, 277, 281, 283, 293, 307, 311,
	313, 317, 331, 337, 347, 349, 353, 359,
	367, 373, 379, 383, 389, 397, 401, 409,
	419, 421, 431, 433, 439, 443, 449, 457,
	461, 463, 467, 479, 487, 491, 499, 503,
	509, 521, 523, 541, 547, 557, 563, 569,
	571, 577, 587, 593, 599, 601, 607, 613,
	617, 619, 631, 641, 643, 647, 653, 659,
	661, 673, 677, 683, 691, 701, 709, 719,
	727, 733, 739, 743, 751, 757, 761, 769,
	773, 787, 797, 809, 811, 821, 823, 827,
	829, 839, 853, 857, 859, 863, 877, 881,
	883, 887, 907, 911, 919, 929, 937, 941,
	947, 953, 967, 971, 977, 983, 991, 997,
}

func TestGetPrimesWith1000(t *testing.T) {
	max := 1000
	primes := GetPrimes(max)
	for i, prime := range primes {
		expectedPrime := expectedPrimes[i]
		if prime != expectedPrime {
			t.Errorf("%dth prime number %d is not the expected value %d",
				i, prime, expectedPrime)
		}
	}
	if t.Failed() == false {
		t.Logf("The primes less than %d are all correct.", max)
	}
}

func BenchmarkGetPrimesWith100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(100)
	}
}

func BenchmarkGetPrimesWith10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(10000)
	}
}

func BenchmarkGetPrimesWith1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000000)
	}
}

/**

$ go test -bench=. -run=^$ -cpu=1,2,4 puzzlers/src/puzzlers/article20/q3
goos: windows
goarch: amd64
pkg: puzzlers/src/puzzlers/article20/q3
BenchmarkGetPrimes        621372              1948 ns/op
BenchmarkGetPrimes-2      685741              1985 ns/op
BenchmarkGetPrimes-4      600038              2007 ns/op
PASS
ok      puzzlers/src/puzzlers/article20/q3      4.329s

概括来讲，go test命令每一次对性能测试函数的执行，都是一个探索的过程。它会在测试函数的执行时间上限不变的前提下，尝试找到被测程序的最大执行次数。

在这个过程中，性能测试函数可能会被执行多次。为了以后描述方便，我们把这样一个探索的过程称为：对性能测试函数的一次探索式执行，
这其中包含了对该函数的若干次执行，当然，肯定也包括了对被测程序更多次的执行。

-count 5，那么对于每一个测试函数，命令都会在预设的不同条件下（比如不同的最大 P 数量下）分别重复执行五次

$ go test -bench=. -run=^$ -cpu=1,2,4  -count=5 puzzlers/src/puzzlers/article20/q3
goos: windows
goarch: amd64
pkg: puzzlers/src/puzzlers/article20/q3
BenchmarkGetPrimes        640077              1721 ns/op
BenchmarkGetPrimes        685843              1746 ns/op
BenchmarkGetPrimes        712242              1752 ns/op
BenchmarkGetPrimes        685760              1723 ns/op
BenchmarkGetPrimes        686828              1737 ns/op
BenchmarkGetPrimes-2      600855              1897 ns/op
BenchmarkGetPrimes-2      662068              1911 ns/op
BenchmarkGetPrimes-2      620343              1913 ns/op
BenchmarkGetPrimes-2      619414              1924 ns/op
BenchmarkGetPrimes-2      663129              1931 ns/op
BenchmarkGetPrimes-4      641799              1926 ns/op
BenchmarkGetPrimes-4      619258              1904 ns/op
BenchmarkGetPrimes-4      600102              1912 ns/op
BenchmarkGetPrimes-4      600949              1906 ns/op
BenchmarkGetPrimes-4      619645              1935 ns/op
PASS
ok      puzzlers/src/puzzlers/article20/q3      18.512s

 */

/**


性能测试函数的执行次数 = `-cpu`标记的值中正整数的个数 x `-count`标记的值 x 探索式执行中测试函数的实际执行次数

功能测试函数的执行次数 = `-cpu`标记的值中正整数的个数 x `-count`标记的值

追加标记-parallel，该标记的作用是：设置同一个被测代码包中的功能测试函数的最大并发执行数

M 和 P 是多对多的，或者说是动态结合的。一个 P 在不同时刻可能会对接不同的 M，反过来也是如此。
而且 M 和 P 在数量上没有直接的关系。P 和 G 表面上说是一对多的，但其实一个 G 在它的生命周期中也可能会在不同 P 的可运行队列之间游走。
更详细的东西，可以看我写的那本书。
 */