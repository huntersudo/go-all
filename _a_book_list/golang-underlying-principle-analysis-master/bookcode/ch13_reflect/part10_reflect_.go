package main

import (
	"fmt"
	"reflect"
)

func main() {

	ta:=reflect.ArrayOf(5,reflect.TypeOf(123))
	fmt.Printf("%v\n",ta) //[5]int

	tc:=reflect.ChanOf(reflect.SendDir,ta)
	fmt.Printf("%v\n",tc) //chan<- [5]int

	tp:=reflect.PtrTo(ta)
	fmt.Printf("%v\n",tp) //*[5]int
	tp1:=reflect.PtrTo(tc)
	fmt.Printf("%v\n",tp1) //*chan<- [5]int

    ts:=reflect.SliceOf(tp)
	fmt.Printf("%v\n",ts) //[]*[5]int

	tm:=reflect.MapOf(ta,tc)
	fmt.Printf("%v\n",tm) //map[[5]int]chan<- [5]int

	//FuncOf(in, out []Type, variadic bool)
	tf:=reflect.FuncOf([]reflect.Type{ta},[]reflect.Type{tp,tc},false)
	fmt.Printf("%v\n",tf) //func([5]int) (*[5]int, chan<- [5]int)


	tt:=reflect.StructOf([]reflect.StructField{
		{Name:"Age",Type:reflect.TypeOf("abc")},
	})
	fmt.Printf("%v\n",tt) //struct { Age string }

	//MakeChan(typ Type, buffer int)  todo 第一个参数 reflect.Type
	//reflect.MakeMap()
	//reflect.MakeFunc()
	//reflect.MakeChan()





}
