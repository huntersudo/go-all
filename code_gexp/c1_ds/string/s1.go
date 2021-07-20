package main

// src/builtin/builtin.go
// string is the set of all strings of 8-bit bytes, conventionally but not
// necessarily representing UTF-8-encoded text. A string may be empty, but
//not nil. Values of string type are immutable.
type string1 string

// string可以为空（长度为0），但不会是nil；
// string对象不可以修改。
//字符串构建过程是先跟据字符串构建stringStruct，再转换成string。
//string在runtime包中就是stringStruct，对外呈现叫做string。

/**
源码包 src/runtime/string.go:stringStruct 定义了string的数据结构：

 type stringStruct struct {
  str unsafe.Pointer
  len int
 }

 */

// byte切片可以很方便的转换成string，需要注意的是这种转换需要一次内存拷贝
// string转换成byte切片，也需要一次内存拷贝，

