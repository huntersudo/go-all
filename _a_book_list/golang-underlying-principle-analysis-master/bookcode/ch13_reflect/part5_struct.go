package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("jonson ReflectCallFunc")
}

// 结构体与反射
func main() {
	user := User{1, "jonson", 25}
	getType := reflect.TypeOf(user)
	fmt.Println("get Type is :", getType.Name()) //User
	getValue := reflect.ValueOf(user)
	fmt.Println("get all Fields is:", getValue) // {1 jonson 25}

	// 遍历结构体字段
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v =%v\n", field.Name, field.Type, value)
	}
	// Id: int =1
	//Name: string =jonson
	//Age: int =25
	/**
	// A StructField describes a single field in a struct.
	type StructField struct {
		// Name is the field name.
		Name string
		// PkgPath is the package path that qualifies a lower case (unexported)
		// field name. It is empty for upper case (exported) field names.
		// See https://golang.org/ref/spec#Uniqueness_of_identifiers
		PkgPath string

		Type      Type      // field type
		Tag       StructTag // field tag string
		Offset    uintptr   // offset within struct, in bytes
		Index     []int     // index sequence for Type.FieldByIndex
		Anonymous bool      // is an embedded field
	}

	// structType represents a struct type.
	type structType struct {
		rtype
		pkgPath name
		fields  []structField // sorted by offset
	}


	 */
}
