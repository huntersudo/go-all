package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
//todo 指针类型的 Value 不能使用 NumField 方法，在执行此方法前需要调用 Elem() 提取底 层值。
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
