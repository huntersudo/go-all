package main

// 错误返回

// todo 在处理错误返回值的时候，没有错误的返回值最好直接写为 nil 。
func returnsError() error {
	 if bad() {
	 	return (*MyError)(err)
	 }
	 return nil
}