package main

import "fmt"

// todo 通过刚刚的一些示例，你现在应该有点明白了，Map、Reduce、Filter 只是一种控制逻辑，真正的业务逻辑是以传给它们的数据和函数来定义的。
//   是的，这是一个很经典的“业务逻辑”和“控制逻辑”分离解耦的编程模式。
//   接下来，我们来看一个有业务意义的代码，来进一步帮助你理解什么叫“控制逻辑”与“业务逻辑”分离。


type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

var list = []Employee{
	{"Hao", 44, 0, 8000},
	{"Bob", 34, 10, 5000},
	{"Alice", 23, 5, 9000},
	{"Jack", 26, 0, 4000},
	{"Tom", 48, 9, 7500},
	{"Marry", 29, 0, 6000},
	{"Mike", 32, 8, 4000},
}
func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i, _ := range list {
		if fn(&list[i]) {
			count += 1
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i, _ := range list {
		if fn(&list[i]) {
			newList = append(newList, list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i, _ := range list {
		sum += fn(&list[i])
	}
	return sum
}
// 1. 统计有多少员工大于 40 岁
func useAgeMore40(){
	old := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people: %d\n", old)
	//old people: 2
}
// 统计有多少员工的薪水大于 6000
func userSalaryMoreThan6000(){
	high_pay := EmployeeCountIf(list, func(e *Employee) bool {
		return e.Salary > 6000
	})
	fmt.Printf("High Salary people: %d\n", high_pay)
	//High Salary people: 4
}
//3. 列出有没有休假的员工
func useNoVocation(){
	no_vacation := EmployeeFilterIn(list, func(e *Employee) bool {
		return e.Vacation == 0
	})
	fmt.Printf("People no vacation: %v\n", no_vacation)
	//People no vacation: [{Hao 44 0 8000} {Jack 26 0 4000} {Marry 29 0 6000}]
}
// 4. 统计所有员工的薪资总和
func sumSalary(){

	total_pay := EmployeeSumIf(list, func(e *Employee) int {
		return e.Salary
	})

	fmt.Printf("Total Salary: %d\n", total_pay)
	//Total Salary: 43500
}
// 5. 统计 30 岁以下员工的薪资总和
func sumSalaryFilterAgeLess30(){
	younger_pay := EmployeeSumIf(list, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	fmt.Printf("younger_pay: %d\n", younger_pay)
}
// todo 泛型 Map-Reduce刚刚的 Map-Reduce 都因为要处理数据的类型不同，而需要写出不同版本的 Map-Reduce，
//  虽然它们的代码看上去是很类似的。所以，这里就要提到泛型编程了
