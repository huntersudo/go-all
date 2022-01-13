package concurrency

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			//，当我们迭代 urls 时，不是直接写入 map ，而是使用 send statement 将每个调用 wc 的 result 结构体发送到 resultChannel 。
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// 将想要加快速度的那部分代码并行化，同时确保不能并发的部分仍然是线性处理
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		// channel to map
		results[r.string] = r.bool
	}

	return results
}
//todo
//这个比在 TDD 上的寻常练习轻松一些。某种程度说，我们已经参与了 CheckWebsites 函数的一 个长期重构；
//输入和输出从未改变，它只是变得更快了。但是我们所做的测试以及我们编写的基准测试 允许我们重构 CheckWebsites ，
//让我们有信心保证软件仍然可以工作，同时也证明它确实变得更快 了
//goroutines 是 Go 的基本并发单元，它让我们可以同时检查多个网站。
//anonymous functions（匿名函数），我们用它来启动每个检查网站的并发进程。
//channels，用来组织和控制不同进程之间的交流，使我们能够避免 race condition（竞争条 件） 的问题。
//the race detector（竞争探测器） 帮助我们调试并发代码的问题
