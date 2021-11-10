package main

import "errors"

//先声明一种函数接口，表示我们的 Undo 控制可以接受的函数签名是什么样的

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}
//看到这里，你不必觉得奇怪， Undo 本来就是一个类型，不必是一个结构体，是一个函数数组也没有什么问题。
// todo 然后，我们在 IntSet 里嵌入 Undo，接着在 Add() 和 Delete() 里使用刚刚的方法，就可以完成功能了


type IntSetUndo struct {
	data map[int]bool
	undo Undo
}

func NewIntSetUndo() IntSetUndo {
	return IntSetUndo{data: make(map[int]bool)}
}

func (set *IntSetUndo) Undo() error {
	return set.undo.Undo()
}

func (set *IntSetUndo) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSetUndo) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSetUndo) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

//todo 这个就是控制反转，不是由控制逻辑 Undo 来依赖业务逻辑 IntSet，而是由业务逻辑 IntSet 依赖 Undo 。
// 这里依赖的是其实是一个协议，这个协议是一个没有参数的函数数组。可以看到，这样一来，我们 Undo 的代码就可以复用了