package main

import "fmt"

//todo  kubectl 的代码⽐较复杂，不过，其本原理简单来说，它从命令⾏和yaml⽂件中获取信息，通 过Builder模式并把其转成⼀系列的资源，最后⽤ Visitor 模式模式来迭代处理这些Reources。
// https://github.com/kubernetes/kubernetes/blob/cea1d4e20b4a7886d8ff65f34c6d4f95efcb4742/staging/src/k8s.io/cli-runtime/pkg/resource/visitor.go
//  kubernetes/vendor/k8s.io/cli-runtime/pkg/resource/visitor.go

//Visitor 模式定义首先，kubectl 主要是用来处理 Info结构体，下面是相关的定义：

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

// 最后，为 Info 实现 Visitor 接⼝中的 Visit() ⽅法，实现就是直接调⽤传进来的⽅法

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

// 有一个 VisitorFunc 的函数类型的定义；

//Name Visitor这个 Visitor 主要是用来访问 Info 结构中的 Name 和 NameSpace 成员：

type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

// 可以看到，在这段代码中：声明了一个 NameVisitor 的结构体，这个结构体里有一个 Visitor 接口成员，这里意味着多态；
// 在实现 Visit() 方法时，调用了自己结构体内的那个 Visitor的 Visitor() 方法，这其实是一种修饰器的模式，用另一个 Visitor 修饰了自己（关于修饰器模式，可以复习下第 113 讲）。

//Other Visitor这个 Visitor 主要用来访问 Info 结构中的 OtherThings 成员：
type OtherThingsVisitor struct {
	visitor Visitor
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Println("==> LogVisitor")
		}
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

//使用方代码现在，我们看看使用上面的代码：
// todo

func main() {
	info := Info{}
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)

	fmt.Println()
	fmt.Println("###########################")
	fmt.Println()

	// 装饰器模式
	info1 := Info{}
	var v1 Visitor = &info1
	// Cannot use 'OtherThingsVisitor' (type OtherThingsVisitor struct { visitor Visitor }) as the type VisitorFunc
	//v1 = NewDecoratedVisitor(v1, NameVisitor, OtherThingsVisitor) // todo  不太对

	v1 = NewDecoratedVisitor(v1, NameVisitorFun, OtherThingsVisitorFun)

	v1.Visit(loadFile)
}

//可以看到，
//todo Visitor 们一层套一层；
//     我用 loadFile 假装从文件中读取数据；
//todo 最后执行  v.Visit(loadfile)  ，这样，我们上面的代码就全部开始激活工作了。
//这段代码输出如下的信息，你可以看到代码是怎么执行起来的：

/**
LogVisitor() before call function
NameVisitor() before call function
OtherThingsVisitor() before call function
==> OtherThings=We are running as remote team.
OtherThingsVisitor() after call function
==> Name=Hao Chen, NameSpace=MegaEase
NameVisitor() after call function
==> LogVisitor
LogVisitor() after call function

*/

//上面的代码有以下几种功效：
//解耦了数据和程序；
//使用了修饰器模式；
//还做出了 Pipeline 的模式。所以，其实我们可以重构一下上面的代码。

// todo refactor
//   Visitor 修饰器我们用修饰器模式来重构一下上面的代码。

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}
func NameVisitorFun(info *Info, err error) error {
	fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
	return nil
}
func OtherThingsVisitorFun(info *Info, err error) error {
	fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
	return nil
}

//这段代码并不复杂，
// todo 我来解释下。
//   用一个 DecoratedVisitor 的结构来存放所有的VistorFunc函数；
//   NewDecoratedVisitor 可以把所有的 VisitorFunc转给它，构造 DecoratedVisitor 对象；
//   DecoratedVisitor实现了 Visit() 方法，里面就是来做一个 for-loop，顺着调用所有的 VisitorFunc。
