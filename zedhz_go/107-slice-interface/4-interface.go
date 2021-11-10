package main

import "fmt"

// 接口编程
// 下面，我们来看段代码，其中是两个方法，它们都是要输出一个结构体，其中一个使用一个函数，另一个使用一个“成员函数”。
type Person struct {
	Name string
	Sexual string
	Age  int
}

func PrintPerson(p *Person) {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func (p *Person) Print() {
	fmt.Printf("Name=%s, Sexual=%s, Age=%d\n",
		p.Name, p.Sexual, p.Age)
}

func main() {
	var p = Person{
		Name: "Hao Chen",
		Sexual: "Male",
		Age: 44,
	}

	PrintPerson(&p)
	p.Print()
}
// 你更喜欢哪种方式呢？
// todo 在 Go 语言中，使用“成员函数”的方式叫“Receiver”，这种方式是一种封装，
// 因为 PrintPerson()本来就是和 Person强耦合的，所以理应放在一起。更重要的是，这种方式可以进行接口编程，
// 对于接口编程来说，也就是一种抽象，主要是用在“多态”，这个技术，我在《Go 语言简介（上）：接口与多态》中讲过，你可以点击链接查看。

