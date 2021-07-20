package main

import (
	"fmt"
	"time"
)

/**
使用select可以监控多channel，比如监控多个channel，当其中某一个channel有数据时，就从其读出数据
*/

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

/**
程序中创建两个channel： chan1和chan2。
函数addNumberToChan()函数会向两个channel中周期性写入数 据。
通过select可以监控两个channel，任意一个可读时就从其中读出数据。
 */
func main() {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)
	go addNumberToChan(chan1)
	go addNumberToChan(chan2)
	for {
		/**
		事实上select语句的多个case执行顺序是随机的
		 */
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
		default:
			fmt.Printf("No element in chan1 and chan2.\n")
			time.Sleep(1 * time.Second)
		}
	}
}
/**
通过这个示例想说的是：select的case语句读channel不会阻塞，尽管channel中没有数据。
这是由于case语句编 译后调用读channel时会明确传入不阻塞的参数，
   此时读不到数据时不会将当前goroutine加入到等待队列，而是直 接返回。
 */

/**
No element in chan1 and chan2.
Get element from chan1: 1
Get element from chan2: 1
No element in chan1 and chan2.
Get element from chan2: 1
Get element from chan1: 1
Get element from chan2: 1
Get element from chan1: 1
No element in chan1 and chan2.
No element in chan1 and chan2.

 */