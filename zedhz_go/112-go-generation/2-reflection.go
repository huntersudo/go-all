package main

import (
	"fmt"
	"reflect"
)

type ContainerV2 struct {
	s reflect.Value
}
func NewContainer(t reflect.Type, size int) *ContainerV2 {
	if size <=0  { size=64 }
	return &ContainerV2{
		s: reflect.MakeSlice(reflect.SliceOf(t), 0, size),
	}
}
func (c *ContainerV2) Put(val interface{})  error {
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		return fmt.Errorf("Put: cannot put a %T into a slice of %s", val, c.s.Type().Elem())
	}
	c.s = reflect.Append(c.s, reflect.ValueOf(val))
	return nil
}
func (c *ContainerV2) Get(refval interface{}) error {
	if reflect.ValueOf(refval).Kind() != reflect.Ptr ||
		reflect.ValueOf(refval).Elem().Type() != c.s.Type().Elem() {
		return fmt.Errorf("Get: needs *%s but got %T", c.s.Type().Elem(), refval)
	}
	reflect.ValueOf(refval).Elem().Set( c.s.Index(0) ) // todo 将slice的 第一个 set refval的值
	c.s = c.s.Slice(1, c.s.Len())
	return nil
}
//todo 这里的代码并不难懂，这是完全使用 Reflection 的玩法，我简单解释下。
//在 NewContainer()时，会根据参数的类型初始化一个 Slice。
//在 Put()时，会检查 val 是否和 Slice 的类型一致。
//在 Get()时，我们需要用一个入参的方式，因为我们没有办法返回 reflect.Value 或 interface{}，不然还要做 Type Assert。
//不过有类型检查，所以，必然会有检查不对的时候，因此，需要返回 error。

func use2(){

	f1 := 3.1415926
	f2 := 1.41421356237

	c := NewContainer(reflect.TypeOf(f1), 16)

	if err := c.Put(f1); err != nil {
		panic(err)
	}
	if err := c.Put(f2); err != nil {
		panic(err)
	}

	g := 0.0

	if err := c.Get(&g); err != nil {
		panic(err)
	}
	fmt.Printf("%v (%T)\n", g, g) //3.1415926 (float64)
	fmt.Println(c.s.Index(0)) //1.4142135623
}

// todo 可以看到，Type Assert 是不用了，但是用反射写出来的代码还是有点复杂的。那么，有没有什么好的方法？
// 对于泛型编程最牛的语言 C++ 来说，这类问题都是使用 Template 解决的。

/**

//用<class T>来描述泛型
template <class T>
T GetMax (T a, T b)  {
    T result;
    result = (a>b)? a : b;
    return (result);
}

int i=5, j=6, k;
//生成int类型的函数
k=GetMax<int>(i,j);

long l=10, m=5, n;
//生成long类型的函数
n=GetMax<long>(l,m);

todo C++ 的编译器会在编译时分析代码，根据不同的变量类型来自动化生成相关类型的函数或类，在 C++ 里，叫模板的具体化。
这个技术是编译时的问题，所以，我们不需要在运行时进行任何的类型识别，我们的程序也会变得比较干净。
那么，我们是否可以在 Go 中使用 C++ 的这种技术呢？答案是肯定的，只是 Go 的编译器不会帮你干，你需要自己动手。
*/