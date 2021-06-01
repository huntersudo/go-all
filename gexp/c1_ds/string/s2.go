package main

//字符串拼接

// str := "Str1" + "Str2" + "Str3"
// 即便有非常多的字符串需要拼接，性能上也有比较好的保证，因为新字符串的内存空间是一次分配完成的，所以性能 消耗主要在拷贝数据上。
//一个拼接语句的字符串编译时都会被存放到一个切片中，拼接过程需要遍历两次切片，第一次遍历获取总的字符串长 度，据此申请内存，第二次遍历会把字符串逐个拷贝过去。

/*
func concatstrings(a []string) string { // 字符串拼接
	length := 0 // 拼接后总的字符串长度
	for _, str := range a {
		length += length(str)
	}

	s, b := rawstring(length) // 生成指定大小的字符串，返回一个string和切片，二者共享内存空间
	for _, str := range a {
		copy(b, str) // string无法修改，只能通过切片修改
	b = b[len(str):]
	}
	return s
}
*/
// 因为string是无法直接修改的，所以这里使用rawstring()方法初始化一个指定大小的string，同时返回一个切 片，二者共享同一块内存空间，
// 后面向切片中拷贝数据，也就间接修改了string。

/*
func rawstring(size int) (s string, b []byte) { // 生成一个新的string，返回的string和切片共享相同的空间
	 p := mallocgc(uintptr(size), nil, false)
	 stringStructOf(&s).str = p
	 stringStructOf(&s).len = size
	 *(*slice)(unsafe.Pointer(&b)) = slice{p, size, size}
	 return
}
*/