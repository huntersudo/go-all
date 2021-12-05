package main

import (
	"sync"
)

var count int64  = 0
var m sync.Mutex  // 互斥锁，比atomic的自旋锁要好，
func add() {
	m.Lock()
	count++
	m.Unlock()
}

func main() {
	go add()
	go add()
}