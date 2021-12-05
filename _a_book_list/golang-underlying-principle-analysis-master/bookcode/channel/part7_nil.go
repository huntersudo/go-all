package main

import "fmt"

func main(){
	a := make(chan int)
	b := make(chan int)
	// select 语句对case对nill通道进行操作时，case分支将永远得不到执行
	go func() {
		for i := 0; i < 2; i++ {
			select {
			case a<-1:
				a = nil
			case b<-2:
				b = nil
			}
		}
	}()
	fmt.Println(<-a)
	fmt.Println(<-b)
}