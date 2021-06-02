package main

import (
	"fmt"
	"io"
)

//type IntSet struct { /* ... */
//}
//
//func (*IntSet) String() string {}
//
//var _ = IntSet{}.String() //
//
//var s IntSet
//var _ = s.String()
//
//var _ fmt.Stringer = &s // OK
//var _ fmt.Stringer = s  // compile error: IntSet lacks String method
//
//
//os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
//os.Stdout.Close() // OK: *os.File has Close method
//
//var w io.Writer
//w = os.Stdout
//w.Write([]byte("hello")) // OK: io.Writer has Write method w.Close()
//
//var any interface{}
//any = true
//any = 12.34
//any = "hello"
//any = map[string]int{"one": 1}
//any = new(bytes.Buffer)



//var w io.Writer
//w = os.Stdout
//w = new(bytes.Buffer)
//w = nil

var x interface{} = []int{1, 2, 3}
fmt.Println(x == x)


