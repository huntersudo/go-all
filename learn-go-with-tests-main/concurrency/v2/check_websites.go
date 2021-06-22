package concurrency

import (
	"time"
)

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		//。唯一的区别是循环的每次迭代都会启动一个新的 goroutine，与当前进程（ WebsiteChecker 函数）同时发生，
		// 这里url被重复用于每次迭代了，所以函数里要用 u ，固定为url的副本，无法更改
		go func(u string) {
			// go test -race todo 解决map并发问题
			// 利用channel中转
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}
