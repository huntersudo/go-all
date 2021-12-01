package main

import "fmt"
// todo  多态
//todo 本质是一个指针  具体源码: 一个指向具体类型，一个指向函数列表
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

// 具体的类  todo 实现接口的方法即可 有⼦类(实现了⽗类的全部接⼝⽅法)
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// todo 触发多态  ⽗类类型的变量(指针) 指向(引⽤) ⼦类的具体数据变量
func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {

	var animal AnimalIF //接口的数据类型， 父类指针
	animal = &Cat{"Green"}

	animal.Sleep() //调用的就是Cat的Sleep()方法 , 多态的现象

	animal = &Dog{"Yellow"}

	animal.Sleep() // 调用Dog的Sleep方法，多态的现象


	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	// Cat is Sleep
	// color =  Green
	// kind =  Cat
	showAnimal(&dog)
	// Dog is Sleep
	// color =  Yellow
	// kind =  Dog
}
