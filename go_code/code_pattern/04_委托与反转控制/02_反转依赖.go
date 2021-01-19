package main

import "errors"

// 控制对象
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
		functions[index] = nil
	}
	*undo = functions[:index]
	return nil
}

// 业务逻辑
type InSet struct {
	data map[int]bool
	undo Undo
}

func NewInSet() InSet {
	return InSet{data:make(map[int]bool)}
}

func (set *InSet) Undo() error {
	return set.undo.Undo()
}

func (set *InSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() {
			set.Delete(x)
		})
	} else {
		set.undo.Add(nil)
	}
}

func (set *InSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() {
			set.Add(x)
		})
	} else {
		set.undo.Add(nil)
	}

}

func (set *InSet) Contains(x int) bool {
	return set.data[x]
}
