package main

import "reflect"

//健壮版的 Generic Filter


func Filter2(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

func FilterInPlace2(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type, " +
			"not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}

var boolType = reflect.ValueOf(true).Type()

func filter(slice, function interface{}, inPlace bool) (interface{}, int) {

	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, boolType) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}

	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	out := sliceInType

	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}

	return out.Interface(), len(which)
}

//todo 后记
// 使用反射来做这些东西会有一个问题，那就是代码的性能会很差。所以，上面的代码不能用在需要高性能的地方。怎么解决这个问题，我会在下节课给你介绍下
// 这节课中的代码大量地参考了 Rob Pike 的版本，你可以点击这个链接查看： https://github.com/robpike/filter。