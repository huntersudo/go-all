package main

import "fmt"

/**
该段程序源自select的实现代码，
程序中定义一个长度为10的切片order，pollorder和lockorder分 别是对order切片做了order[low:high:max]操作生成的切片，
最后程序分别打印pollorder和lockorder的容 量和长度。
 */
func main() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]
	/**
	order[low:high:max]操作意思是对order进行切片，新切片范围是[low, high),新切片容量是max
	order长度为2倍的orderLen，
	pollorder切片指的是order的前半部分切片，lockorder指的是order的 后半部分切片，即原order分成了两段。
	所以，pollorder和lockerorder的长度和容量都是orderLen，即5。

	 */
	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))

	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}
