package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close可以关闭一个channel
		close(c)
	}()

	//todo  可以使用range来迭代不断操作channel  会自动判断 chan 是否close
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished..")
}
