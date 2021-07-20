package main

import (
	"log"
	"sync"
	"time"
)

/**
优势是并发流程上的协同，chan的主要任务是传递数据。另外cond是更低层次的工具，效率更高一些，但是肯定没有chan方便。

如果都用同一个互斥量的话，操作双方就无法独立行事，这就是完全串行的操作了，效率上会大打折扣。

进一步说，本来就是一个发一个收，理应一个用写锁一个用读锁，这样效率高，之后扩展起来也方便。因为读之间不用互斥。

for mailbox == 1 {
sendCond.Wait()  // 有情报，那么我就回家去等蓝帽子小孩儿了
}
有可能会碰到“假唤醒”的情况。而且，如果存在“有多个wait但只需唤醒一个”的情况，也需要用for语句。
在for语句里，唤醒后可以再次检查状态，如果状态符合就开始后续工作，如果不符合就再次wait。用if语句就办不到

 */

func main() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	var lock sync.RWMutex

	// sendCond 代表专用于发信的条件变量。 蓝帽子小孩
	// TODO 注意： lock变量的Lock方法和Unlock方法分别用于对其中写锁的锁定和解锁，它们与sendCond变量的含义是对应的
	sendCond := sync.NewCond(&lock)
	// recvCond 代表专用于收信的条件变量。 红帽子小孩
	//TODO  注意： 通过lock.RLocker()得来的值就是lock变量中的读锁，这个值所拥有的Lock方法和Unlock方法，
	// 在其内部会分别调用lock变量的RLock方法和RUnlock方法。也就是说，前两个方法仅仅是后两个方法的代理而已。
	// TODO 和demo62.go 不一样
	recvCond := sync.NewCond(lock.RLocker())

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 5
	go func(max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			// 这个Lock方法在这里意味的是：持有信箱上的锁，并且有打开信箱的权利，而不是锁上这个锁
			lock.Lock()
			for mailbox == 1 {  //  for 会多检查一次，看上面
				sendCond.Wait()  // 有情报，那么我就回家去等蓝帽子小孩儿了
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			// 离开之后我还要做一件事，那就是让红帽子小孩儿准时去你家楼下路过。也就是说，我会及时地通知你“信箱里已经有新情报
			recvCond.Signal()
		}
	}(max)

	go func(max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait() // 需要回家去等红帽子小孩儿
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been received.", j)
			lock.RUnlock()
			// 你还需要让蓝帽子小孩儿准时去我家楼下路过。这样我就知道信箱中的情报已经被你获取了
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
/**

2021/03/05 15:46:00 sender [1]: the mailbox is empty.
2021/03/05 15:46:00 sender [1]: the letter has been sent.
2021/03/05 15:46:00 receiver [1]: the mailbox is full.
2021/03/05 15:46:00 receiver [1]: the letter has been received.
2021/03/05 15:46:00 sender [2]: the mailbox is empty.
2021/03/05 15:46:00 sender [2]: the letter has been sent.
2021/03/05 15:46:00 receiver [2]: the mailbox is full.
2021/03/05 15:46:00 receiver [2]: the letter has been received.
2021/03/05 15:46:01 sender [3]: the mailbox is empty.
2021/03/05 15:46:01 sender [3]: the letter has been sent.
2021/03/05 15:46:01 receiver [3]: the mailbox is full.
2021/03/05 15:46:01 receiver [3]: the letter has been received.
2021/03/05 15:46:01 sender [4]: the mailbox is empty.
2021/03/05 15:46:01 sender [4]: the letter has been sent.
2021/03/05 15:46:01 receiver [4]: the mailbox is full.
2021/03/05 15:46:01 receiver [4]: the letter has been received.
2021/03/05 15:46:02 sender [5]: the mailbox is empty.
2021/03/05 15:46:02 sender [5]: the letter has been sent.
2021/03/05 15:46:02 receiver [5]: the mailbox is full.
2021/03/05 15:46:02 receiver [5]: the letter has been received.

 */