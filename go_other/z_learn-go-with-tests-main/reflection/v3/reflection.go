package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
//我们需要检查字段的类型是 string 。
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
