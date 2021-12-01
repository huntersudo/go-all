package main

import "errors"

type IntSet struct {
	data map[int]bool
}
func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}
func (set *IntSet) Add(x int) {
	set.data[x] = true
}
func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}
func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

//todo 现在，我们想实现一个 Undo 的功能。我们可以再包装一下  IntSet  ，变成 UndoableIntSet  ，代码如下所示：


type UndoableIntSet struct { // Poor style
	IntSet    // Embedding (delegation)
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

func (set *UndoableIntSet) Add(x int) { // Override
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Delete(x int) { // Override
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // For garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}
// todo 我来解释下这段代码。
//我们在 UndoableIntSet 中嵌入了IntSet ，然后 Override 了 它的 Add()和 Delete() 方法；
//Contains() 方法没有 Override，所以，就被带到 UndoableInSet 中来了。
//在 Override 的 Add()中，记录 Delete 操作；
//在 Override 的 Delete() 中，记录 Add 操作；
//在新加入的 Undo() 中进行 Undo 操作。

