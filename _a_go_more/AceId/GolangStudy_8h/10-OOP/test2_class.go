package main

import "fmt"

// Hero OOP
//todo 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	//todo 如果说类的属性首字母大写, 表示该属性是对外能够访问的，否则的话只能够类的内部访问
	Name  string
	Ad    int
	level int
}

/*func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.level)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string) {
	//todo this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}*/

func (this *Hero) Show1() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//todo this 是调用该方法的对象的指针
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100}

	hero.Show1()
// Name =  zhang3
	//Ad =  100
	//Level =  0
	hero.SetName("li4")

	hero.Show1()
	// Name =  li4
	//Ad =  100
	//Level =  0
}
