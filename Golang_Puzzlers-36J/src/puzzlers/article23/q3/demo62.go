package main

import (
	"log"
	"sync"
	"time"
)


/**
为什么先要锁定条件变量基于的互斥锁，才能调用它的Wait方法？

条件变量的Wait方法主要做了四件事。
1、把调用它的 goroutine（也就是当前的 goroutine）加入到当前条件变量的通知队列中。
2、解锁当前的条件变量基于的那个互斥锁。
3、让当前的 goroutine 处于等待状态，等到通知到来时再决定是否唤醒它。此时，这个 goroutine 就会阻塞在调用这个Wait方法的那行代码上。
4、如果通知到来并且决定唤醒这个 goroutine，那么就在唤醒它之后重新锁定当前条件变量基于的互斥锁。自此之后，当前的 goroutine 就会继续执行后面的代码了。

因为条件变量的Wait方法在阻塞当前的 goroutine 之前，会解锁它基于的互斥锁，所以在调用该Wait方法之前，我们必须先锁定那个互斥锁
为什么条件变量的Wait方法要这么做呢？你可以想象一下，如果Wait方法在互斥锁已经锁定的情况下，阻塞了当前的 goroutine，那么又由谁来解锁呢？别的 goroutine 吗？
---必须由当前的goroutine自己解锁。

====
虽然等待的 goroutine 有多个，但每次成功的 goroutine 却只可能有一个。别忘了，条件变量的Wait方法会在当前的 goroutine 醒来后先重新锁定那个互斥锁。

多个goroutine都被唤醒，但是只有一个获得lock，然后其他goroutine先后会进入临界区，然后发现共享资源的状态不对，所以，需要继续for循环，多次检查。

 */

/**
if语句只会对共享资源的状态检查一次，而for语句却可以做多次检查，直到这个状态改变为止。那为什么要做多次检查呢

这主要是为了保险起见。如果一个 goroutine 因收到通知而被唤醒，但却发现共享资源的状态，依然不符合它的要求，
那么就应该再次调用条件变量的Wait方法，并继续等待下次通知的到来。

 */


func main() {
	// mailbox 代表信箱。
	// 0代表信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock 代表信箱上的锁。
	var lock sync.Mutex
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)
	// recvCond 代表专用于收信的条件变量。
	// TODO 和demo61.go 不一样
	recvCond := sync.NewCond(&lock)

	// send 代表用于发信的函数。
	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		log.Printf("sender [%d-%d]: the mailbox is empty.",
			id, index)
		mailbox = 1
		log.Printf("sender [%d-%d]: the letter has been sent.",
			id, index)
		lock.Unlock()
		// 却会唤醒所有为此等待的 goroutine。
		// TODO 和demo61.go 不一样
		recvCond.Broadcast()
	}

	// recv 代表用于收信的函数。
	recv := func(id, index int) {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		log.Printf("receiver [%d-%d]: the mailbox is full.",
			id, index)
		mailbox = 0
		log.Printf("receiver [%d-%d]: the letter has been received.",
			id, index)
		lock.Unlock()
		sendCond.Signal() // 确定只会有一个发信的goroutine。
	}


	// TODO 和demo61.go 不一样

	// sign 用于传递演示完成的信号。
	sign := make(chan struct{}, 3)
	max := 6
	go func(id, max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			send(id, i)
		}
	}(0, max)
	go func(id, max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, j)
		}
	}(1, max/2)
	go func(id, max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= max; k++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, k)
		}
	}(2, max/2)

	<-sign
	<-sign
	<-sign
}
/**
https://time.geekbang.org/column/article/41717?utm_source=java3y&utm_medium=toutiao&utm_term=pc_interstitial_878
对条件变量这个工具本身还有疑问的读者可以去看我写的《Go 并发编程实战》第二版。这本书从并发程序的基本概念讲起，用一定篇幅的图文内容详细地讲解了条件变量的用法，
同时还有一个贯穿了互斥锁和条件变量的示例。由于这本书的版权在出版社那里，所以我不能把其中的内容搬到这里。

我在这里只对大家共同的疑问进行简要说明：

1. 条件变量适合保护那些可执行两个对立操作的共享资源。比如，一个既可读又可写的共享文件。又比如，既有生产者又有消费者的产品池。

2. 对于有着对立操作的共享资源（比如一个共享文件），我们通常需要基于同一个锁的两个条件变量（比如 rcond 和 wcond）分别保护读操作和写操作（比如 rcond 保护读，wcond 保护写）。
而且，读操作和写操作都需要同时持有这两个条件变量。因为，读操作在操作完成后还要向 wcond 发通知；写操作在操作完成后还要向 rcond 发通知。如此一来，读写操作才能在较少的锁争用的情况下交替进行。

3. 对于同一个条件变量，我们在调用它的 Signal 方法和 Broadcast 方法的时候不应该处在其包含的那个锁的保护下。也就是说，我们应该先撤掉保护屏障，再向 Wait 方法的调用方发出通知。
否则，Wait 方法的调用方就有可能会错过通知。这也是我更推荐使用 Broadcast 方法的原因。所有等待方都错过通知的概率要小很多。

4. 相对应的，我们在调用条件变量的 Wait 方法的时候，应该处在其中的锁的保护之下。因为有同一个锁保护，所以不可能有多个 goroutine 同时执行到这个 Wait 方法调用，
也就不可能存在针对其中锁的重复解锁。

5. 再强调一下。对于同一个锁，多个 goroutine 对它重复锁定时只会有一个成功，其余的会阻塞；多个 goroutine 对它重复解锁时也只会有一个成功，但其余的会抛 panic


老师，我有一个疑问，对于cond来说，每次只唤醒一个goruntine，如果这么goruntine发现消息不是自己想要的就会从新阻塞在wait函数中，
那么真正需要这个消息的goruntine还会被唤醒吗？
作者回复: 不会，所以我才说应该优先用 broadcast。



 */
