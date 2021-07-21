package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(waitGroup *sync.WaitGroup, quit <-chan bool) {
	defer waitGroup.Done()

	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(100 * time.Millisecond)
		case <-quit:
			return
		}
	}
}

func main() {
	fmt.Println("Select to Quit a Goroutine")

	quit := make(chan bool)
	var waitGroup sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go worker1(&waitGroup, quit)
	}

	time.Sleep(time.Second)
	close(quit)

	waitGroup.Wait()
}
