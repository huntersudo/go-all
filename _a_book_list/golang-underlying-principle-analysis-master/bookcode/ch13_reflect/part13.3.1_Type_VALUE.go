package main

import "reflect"

func main() {
	/**
	TODO 原理是 将传递过来的接口变量，转变为 底层的实际空接口 ,Type实际是typ字段
	// TypeOf returns the reflection Type that represents the dynamic type of i.
	// If i is a nil interface value, TypeOf returns nil.
	func TypeOf(i interface{}) Type {
		eface := *(*emptyInterface)(unsafe.Pointer(&i))
		return toType(eface.typ)
	}


	// emptyInterface is the header for an interface{} value.
	type emptyInterface struct {
		typ  *rtype
		word unsafe.Pointer
	}


	 */
	reflect.TypeOf()
	reflect.ValueOf()  // todo P208-209
	/**

	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i. ValueOf(nil) returns the zero Value.
	func ValueOf(i interface{}) Value {
		if i == nil {
			return Value{}
		}

		// TODO: Maybe allow contents of a Value to live on the stack.
		// For now we make the contents always escape to the heap. It
		// makes life easier in a few places (see chanrecv/mapassign
		// comment below).
		escapes(i)

		return unpackEface(i)
	}

	// unpackEface converts the empty interface i to a Value.
	func unpackEface(i interface{}) Value {
		e := (*emptyInterface)(unsafe.Pointer(&i))
		// NOTE: don't read e.word until we know whether it is really a pointer or not.
		t := e.typ
		if t == nil {
			return Value{}
		}
		f := flag(t.Kind())
		if ifaceIndir(t) {
			f |= flagIndir  // TODO flagIndir 间接 P209
		}
		return Value{t, e.word, f}
	}

	 */
}
