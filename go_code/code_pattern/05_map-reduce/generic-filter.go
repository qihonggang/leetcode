package main

import (
	"fmt"
	"reflect"
)

func Filter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

func FilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() != reflect.Ptr {
		panic("FilterInPlace: wrong type, " +
			"not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}

var boolType = reflect.ValueOf(true).Type()

func filter(slice, function interface{}, inPlace bool) (interface{}, int) {

	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, boolType) {
		panic("filter: function must be of type func(" + elemType.String() + ") bool")
	}

	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	out := sliceInType

	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}
	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}

	return out.Interface(), len(which)
}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	// check it is a function
	if fn.Kind() != reflect.Func {
		return false
	}

	// NumIn() - return a function type's input parameter count.
	// NumOut() - return a function type's output parameter count.
	if (fn.Type().NumIn() != len(types) - 1) || (fn.Type().NumOut() != 1) {
		return false
	}

	// In() - return the type of a function type's i'th input parameter.
	for i := 0; i < len(types) - 1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}

	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types) - 1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

func main() {
	var intset = []int{1,2,3,4,5,6,7,8,9,10}
	out := Filter(intset, func(n int) bool {
		return n % 2 == 1
	})
	fmt.Printf("%v\n", out)

	out = Filter(intset, func(n int) bool {
		return n > 5
	})
	fmt.Printf("%v\n", out)
}
