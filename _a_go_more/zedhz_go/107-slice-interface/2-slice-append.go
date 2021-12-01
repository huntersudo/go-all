package main

import (
	"bytes"
	"fmt"
)

// slice append

func main1() {

	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)
	a[2] = 42
	fmt.Println(a)
	fmt.Println(b)
// [0 0 42 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1] 可以发现，修改a，不影响b
// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// 在这段代码中，把 a[1:16] 的切片赋给 b ，此时，a 和 b 的内存空间是共享的，
// 然后，对 a 做了一个 append()的操作，这个操作会让 a 重新分配内存，这就会导致 a 和 b 不再共享，如下图所示：

// 从图中，我们可以看到，append()操作让 a 的容量变成了 64，而长度是 33。
//todo 这里你需要重点注意一下，append()这个函数在 cap 不够用的时候，就会重新分配内存以扩大容量，如果够用，就不会重新分配内存了！
}


func main() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path,'/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1,"suffix"...)

	fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB
}
//在这个例子中，dir1 和 dir2 共享内存，
//todo  虽然 dir1 有一个 append() 操作，但是因为 cap 足够，于是数据扩展到了dir2 的空间。
// 如果要解决这个问题，我们只需要修改一行代码。我们要把代码
//  dir1 := path[:sepIndex] 修改未： dir1 := path[:sepIndex:sepIndex]
// 新的代码使用了 Full Slice Expression，最后一个参数叫“Limited Capacity”，于是，后续的 append() 操作会导致重新分配内存。