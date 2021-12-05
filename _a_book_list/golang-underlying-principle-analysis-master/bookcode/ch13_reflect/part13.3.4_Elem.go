package main

import "reflect"

func main() {

	// todo P211-212
	//  elem 返回接口内部包含的 或者指针指向的数据值
	vs := reflect.ValueOf(&s).Elem()
	vx:= vs.Field(0) // 这里必须是结构体调用
	vb := reflect.ValueOf(123)
	vx.Set(vb)

	/**
	// Set assigns x to the value v.
	// It panics if CanSet returns false.
	// As in Go, x's value must be assignable to v's type.
	func (v Value) Set(x Value) {
		v.mustBeAssignable()
		x.mustBeExported() // do not let unexported x leak
		var target unsafe.Pointer
		if v.kind() == Interface {
			target = v.ptr
		}
		x = x.assignTo("reflect.Set", v.typ, target)
		if x.flag&flagIndir != 0 {
			if x.ptr == unsafe.Pointer(&zeroVal[0]) {
				typedmemclr(v.typ, v.ptr)
			} else {
				typedmemmove(v.typ, v.ptr, x.ptr)
			}
		} else {
			*(*unsafe.Pointer)(v.ptr) = x.ptr
		}
	}

	*/
}
