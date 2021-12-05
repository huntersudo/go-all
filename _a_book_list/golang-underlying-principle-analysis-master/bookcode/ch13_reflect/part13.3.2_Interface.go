package main

import "reflect"

func main() {

	reflect.ValueOf().Interface()  // todo P209-210
	/**
	// packEface converts v to the empty interface.
	func packEface(v Value) interface{} {
		t := v.typ
		var i interface{}
		e := (*emptyInterface)(unsafe.Pointer(&i))
		// First, fill in the data portion of the interface.
		switch {
		case ifaceIndir(t):
			if v.flag&flagIndir == 0 {
				panic("bad indir")
			}
	       //  TODO 间接
			// Value is indirect, and so is the interface we're making.
			ptr := v.ptr
			if v.flag&flagAddr != 0 {
				// TODO: pass safe boolean from valueInterface so
				// we don't need to copy if safe==true?
				c := unsafe_New(t)
				typedmemmove(t, c, ptr)
				ptr = c
			}
			e.word = ptr
		case v.flag&flagIndir != 0:
			// Value is indirect, but interface is direct. We need
			// to load the data at v.ptr into the interface data word.
			e.word = *(*unsafe.Pointer)(v.ptr)
		default:
	        // todo 直接
			// Value is direct, and so is the interface.
			e.word = v.ptr
		}
		// Now, fill in the type portion. We're very careful here not
		// to have any operation between the e.word and e.typ assignments
		// that would let the garbage collector observe the partially-built
		// interface value.
		e.typ = t
		return i
	}

	*/
}
