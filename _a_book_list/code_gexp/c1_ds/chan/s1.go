package main

/**

 func makechan(t *chantype, size int) *hchan {
    var c *hchan
    c = new(hchan)
    c.buf = malloc(元素类型大小*size)
    c.elemsize = 元素类型大小
    c.elemtype = 元素类型
    c.dataqsiz = size
    return c
 }

 */

/**
向channel 写数据
向一个channel中写数据简单过程如下：
1. 如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G,并把数据 写入，最后把该G唤醒，结束发送过程；
2. 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程；
3. 如果缓冲区中没有空余位置，将待发送数据写入G，将当前G加入sendq，进入睡眠，等待被读goroutine唤 醒；

 */

/**
从channel 读数据
从一个channel读数据简单过程如下：
1. 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束 读取过程；
2. 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾 部，把G唤醒，结束读取过程；
3. 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程； 4. 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒；

 */

/**
关闭channel

关闭channel时会把recvq中的G全部唤醒，本该写入G的数据位置为nil。
把sendq中的G全部唤醒，但这些G会 panic。
除此之外，panic出现的常见场景还有：
1. 关闭值为nil的channel
2. 关闭已经被关闭的channel
3. 向已经关闭的channel写数据

 */