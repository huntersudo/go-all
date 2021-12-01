package main

import "fmt"

/*type Country struct {
	Name string
}
type City struct {
	Name string
}
type Printable interface {
	PrintStr()
}
func (c Country) PrintStr() {
	fmt.Println(c.Name)
}
func (c City) PrintStr() {
	fmt.Println(c.Name)
}
func main() {
	c1 := Country {"China"}
	c2 := City {"Beijing"}
	c1.PrintStr()
	c2.PrintStr()
}*/
// 可以看到，这段代码中使用了一个 Printable 的接口，
// 而 Country 和 City 都实现了接口方法 PrintStr() 把自己输出。然而，这些代码都是一样的，能不能省掉呢？

/*type WithName struct {
	Name string
}
type Country struct {
	WithName
}
type City struct {
	WithName
}
type Printable interface {
	PrintStr()
}
func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}
func main() {
	c1 := Country {WithName{ "China"}}
	c2 := City { WithName{"Beijing"}}
	c1.PrintStr()
	c2.PrintStr()
}*/
// todo 我们可以使用“结构体嵌入”的方式来完成这个事
//  引入一个叫 WithName的结构体，但是这会带来一个问题：在初始化的时候变得有点乱。那么，有没有更好的方法呢？再来看另外一个解

type Country struct {
	Name string
}
type City struct {
	Name string
}
type Stringable interface {
	ToString() string
}
func (c Country) ToString() string {
	return "Country = " + c.Name
}
func (c City) ToString() string{
	return "City = " + c.Name
}
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}
func main() {
	d1 := Country {"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}
// todo 引入接口
//  在这段代码中，我们可以看到，我们使用了一个叫Stringable 的接口，
//  我们用这个接口把“业务类型” Country 和 City 和“控制逻辑” Print() 给解耦了。
//  于是，只要实现了Stringable 接口，都可以传给 PrintStr() 来使用。

// todo 这种编程模式在 Go 的标准库有很多的示例，最著名的就是 io.Read 和 ioutil.ReadAll 的玩法，
//  其中 io.Read 是一个接口，你需要实现它的一个 Read(p []byte) (n int, err error) 接口方法，
//  只要满足这个规则，就可以被 ioutil.ReadAll这个方法所使用。
//  这就是面向对象编程方法的黄金法则——“Program to an interface not an implementation

