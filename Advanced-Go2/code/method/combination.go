package main

import "sync"

type Cache struct {
	m map[string]string
	sync.Mutex
}

func (p *Cache) Lookup(key string) string{
	p.Lock()
	defer p.Unlock()

	return p.m[key]
}

// todo 编译器绑定
// Cache 结构体类型通过嵌入一个匿名的 sync.Mutex 来继承它的 Lock 和 Unlock 方法.
//但是 在调用 p.Lock() 和 p.Unlock() 时, p 并不是 Lock 和 Unlock 方法的真正接收者, 而是 会将它们展开为 p.Mutex.Lock() 和 p.Mutex.Unlock() 调用. 这种展开是编译期完成的, 并没 有运行时代价.

//todo
// 在传统的面向对象语言(eg.C++或Java)的继承中，子类的方法是在运行时动态绑定到对象的，因此基 类实现的某些方法看到的 this 可能不是基类类型对应的对象，这个特性会导致基类方法运行的不确 定性。
// 而在Go语言通过嵌入匿名的成员来“继承”的基类方法， this 就是实现该方法的类型的对象， Go语言中方法是编译时静态绑定的。如果需要虚函数的多态特性，我们需要借助Go语言接口来实现。