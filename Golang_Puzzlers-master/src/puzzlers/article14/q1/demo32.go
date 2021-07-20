package main

import (
	"fmt"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1。
	dog := Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	// pet  持有dog的副本
	var pet Pet = dog
	// dog的SetName方法是指针方法 ,操作之后，dog的名字会更新
	dog.SetName("monster")
	// pet的名字不变
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	fmt.Println()
/**
  The dog's name is "little pig".
  The dog's name is "monster".
  This pet is a dog, the name is "little pig".
 */


	// 示例2。
	dog1 := Dog{"little pig"}
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	fmt.Println()
/**
  The name of first dog is "little pig".
  The name of second dog is "little pig".
  The name of first dog is "monster".
  The name of second dog is "little pig".
 */

	// 示例3。
	dog = Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	/**
	The dog's name is "little pig".
	The dog's name is "monster".
	This pet is a dog, the name is "monster".
	 */
}
