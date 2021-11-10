package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// 本来，Visitor 是⾯ 向对象设计模英中⼀个很重要的设计模款（参看Wikipedia Visitor Pattern词条），这个模式是⼀ 种将算法与操作对象的结构分离的⼀种⽅法。
// todo 这种分离的实际结果是能够在不修改结构的情况下 向现有对象结构添加新操作，是遵循开放/封闭原则的⼀种⽅法。


type Visitor func(shape Shape)

type Shape interface {
	accept(Visitor)
}

// 我们的实例的对象 Circle和 Rectangle实现了 Shape 接口的 accept() 方法，这个方法就是等外面给我们传递一个 Visitor。

type Circle struct {
	Radius int
}
func (c Circle) accept(v Visitor) {
	v(c)
}

type Rectangle struct {
	Width, Heigh int
}
func (r Rectangle) accept(v Visitor) {
	v(r)
}

// 然后，我们实现两个 Visitor：一个是用来做 JSON 序列化的；另一个是用来做 XML 序列化的。

func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

//下面是使用 Visitor 这个模式的代码：

func main() {
	c := Circle{10}
	r :=  Rectangle{100, 200}
	shapes := []Shape{c, r}

	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}
}

// todo 其实，这段代码的目的就是想解耦数据结构和算法。虽然使用 Strategy 模式也是可以完成的，而且会比较干净，
//   但是在有些情况下，多个 Visitor 是来访问一个数据结构的不同部分，这种情况下，数据结构有点像一个数据库，
//   而各个 Visitor 会成为一个个的小应用。 kubectl就是这种情况。
