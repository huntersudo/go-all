package main

import (
	"fmt"
)

type Animal interface {
	// ScientificName 用于获取动物的学名。
	ScientificName() string
	// Category 用于获取动物的基本分类。
	Category() string
}

type Named interface {
	// Name 用于获取名字。
	Name() string
}

type Pet interface {
	Animal
	Named
}

type PetTag struct {
	name  string
	owner string
}

func (pt PetTag) Name() string {
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog struct {
	PetTag
	scientificName string
}

func (dog Dog) ScientificName() string {
	return dog.scientificName
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	petTag := PetTag{name: "little pig"}
	//_, ok := interface{}(petTag).(Named)
	_, ok := (interface{}(petTag)).(Named)
 // 类型断⾔是⼀个使⽤在接⼝值上的操作。语法上它看起来像x.(T)被称为断⾔类型，
 // 这⾥x表示⼀个接⼝的类型和T表示⼀ 个类型。⼀个类型断⾔检查它操作对象的动态类型是否和断⾔的类型匹配。

	fmt.Printf("PetTag implements interface Named: %v\n", ok)
	// PetTag implements interface Named: true
	dog := Dog{
		PetTag:         petTag,
		scientificName: "Labrador Retriever",
	}
	_, ok = interface{}(dog).(Animal)
	fmt.Printf("Dog implements interface Animal: %v\n", ok)
	// Dog implements interface Animal: true
	_, ok = interface{}(dog).(Named)
	fmt.Printf("Dog implements interface Named: %v\n", ok)
	// Dog implements interface Named: true
	_, ok = interface{}(dog).(Pet)
	fmt.Printf("Dog implements interface Pet: %v\n", ok)
	// Dog implements interface Pet: true
}
