###

我在讲 go 语句的时候说过，这里的 P 是 processor 的缩写，每个 processor 都是一个可以承载若干个 G，且能够使这些 G 适时地与 M 进行对接并得到真正运行的中介。     
正是由于 P 的存在，G 和 M 才可以呈现出多对多的关系，并能够及时、灵活地进行组合和分离。   
这里的 G 就是 goroutine 的缩写，可以被理解为 Go 语言自己实现的用户级线程。M 即为 machine 的缩写，代表着系统级线程，或者说操作系统内核级别的线程。    

G(go自己实现的用户级别线程) - P(逻辑CPU) -M (系统级线程)

P 的数量意味着 Go 程序背后的运行时系统中，会有多少个用于承载可运行的 G 的队列存在。

最大 P 数量就代表着 Go 语言运行时系统同时运行 goroutine 的能力，也可以被视为其中逻辑 CPU 的最大个数。    
而go test命令的-cpu标记正是用于设置这个最大个数的。