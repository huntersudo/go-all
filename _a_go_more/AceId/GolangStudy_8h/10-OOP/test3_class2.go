package main

import "fmt"
/**
继承
 */
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

//=================

type SuperMan struct {
	Human //todo SuperMan类继承了Human类的方法

	level int
}

//todo 重定义父类的方法Eat()  重写父类方法
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

func main() {
	h := Human{"zhang3", "female"}

	h.Eat() // Human.Eat()...
	h.Walk() // Human.Walk()...

	//定义一个子类对象
	//s := SuperMan{Human{"li4", "female"}, 88}
	var s SuperMan
	s.name = "li4"
	s.sex = "male"
	s.level = 88

	s.Walk() //父类的方法   Human.Walk()...
	s.Eat()  //子类的方法  重写了 SuperMan.Eat()...
	s.Fly()  //子类的方法  SuperMan.Fly()...

	s.Print()
}
