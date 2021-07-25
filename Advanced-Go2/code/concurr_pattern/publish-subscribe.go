package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// todo 在传统生产者和消费者模型中，是将消息发送到一个队列中，而发布订阅模型则是将消息发布 给一个主题。
// todo  抽象程序，为啥自己总是想不到 ，不能很快理解
type (
	subscriberChannel chan string                // 订阅者为一个管道
	topicFilter       func(contents string) bool // 主题为一个过滤器
	//topicFilter func(v interface{}) bool  // 主题为一个过滤器
)

// 发布者对象
type publisher struct {
	m           sync.RWMutex                      // 读写锁
	bufferSize  int                               // 订阅队列的缓存大小
	timeout     time.Duration                     // 发布超时时间
	subscribers map[subscriberChannel]topicFilter // 订阅者信息
}

//  构建一个发布者对象, 可以设置发布超时时间和缓存队列的长度
func NewPublisher(timeout time.Duration, bufferSize int) *publisher {
	return &publisher{
		bufferSize:  bufferSize,
		timeout:     timeout,
		subscribers: make(map[subscriberChannel]topicFilter),
	}
}

// 添加一个新的订阅者，订阅全部主题
func (p *publisher) subscribeAll() subscriberChannel {
	return p.subscribeTopicFilter(nil)
}

// 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *publisher) subscribeTopicFilter(filter topicFilter) subscriberChannel {
	ch := make(chan string, p.bufferSize)
	p.m.Lock()
	defer p.m.Unlock()

	p.subscribers[ch] = filter
	return ch
}

// 退出订阅
func (p *publisher) evict(subscriber subscriberChannel) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, subscriber) // map delete key
	close(subscriber)
}

// 发布一个主题
func (p *publisher) publish(contents string) {
	p.m.RLock()
	defer p.m.RUnlock()

	var waitGroup sync.WaitGroup
	for subscriber, topicFilter := range p.subscribers {
		waitGroup.Add(1)
		go p.sendTopic(subscriber, topicFilter, contents, &waitGroup)
	}
	waitGroup.Wait()
}

// 关闭发布者对象，同时关闭所有的订阅者管道。
func (p *publisher) close() {
	p.m.Lock()
	defer p.m.Unlock()

	for subscriber := range p.subscribers {
		delete(p.subscribers, subscriber)
		close(subscriber)
	}
}

// todo 发送主题，可以容忍一定的超时
func (p *publisher) sendTopic(subscriber subscriberChannel, filter topicFilter, contents string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	if filter != nil && !filter(contents) {
		return
	}

	select {
	case subscriber <- contents:
	case <-time.After(p.timeout):
	}
}

// 下面的例子中，有两个订阅者分别订阅了全部主题和含有”golang”的主题：
func main() {
	fmt.Println("Publish and Subscribe Pattern")

	p := NewPublisher(100*time.Millisecond, 10)
	defer p.close()

	// 全部主题
	subscriberForAll := p.subscribeAll()
	// 含有”golang”的主题：
	subscriberForGolang := p.subscribeTopicFilter(func(contents string) bool {
		return strings.Contains(contents, "golang")
	})

	//// for
	//subscriberForGolang_2 := p.subscribeTopicFilter(func(v interface{}) bool {
	//	if s,ok:=v.(string);ok{
	//		return strings.Contains(s,"golang_2")
	//	}
	//	return  false
	//	//return strings.Contains(contents, "golang_2")
	//})

	p.publish("hello, world")
	p.publish("hello, golang")
	p.publish("hello, golang_2")

	go func() {
		for msg := range subscriberForAll {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range subscriberForGolang {
			fmt.Println("golang:", msg)
		}
	}()

	//go func() {
	//	for msg := range subscriberForGolang_2 {
	//		fmt.Println("golang_2:", msg)
	//	}
	//}()
	time.Sleep(3 * time.Second)
}
