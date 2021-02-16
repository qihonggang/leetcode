package main

import (
	"fmt"
	"reflect"
)

func Transform(slice, function interface{}) interface{} {
	return transform(slice, function, false)
}

func TransformInPlace(slice, function interface{}) interface{} {
	return transform(slice, function, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {
	// check the `slice` type is Slice
	sliceIntype := reflect.ValueOf(slice)
	if sliceIntype.Kind() != reflect.Slice {
		panic("tracsform: not slice")
	}

	// check the function signature
	fn := reflect.ValueOf(function)
	elemType := sliceIntype.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceIntype.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceIntype
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)),sliceIntype.Len(), sliceIntype.Len())
	}
	for i := 0; i < sliceIntype.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceIntype.Index(i)})[0])
	}
	return sliceOutType.Interface()
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
	// 字符串数组
	listString := []string{"1", "2", "3", "4", "5", "6"}
	resultString := Transform(listString, func(a string) string{
		return a + a + a
	})
	fmt.Printf("strings: %v\n", resultString)

	// 整形数组
	listInt := []int{1,2,3,4,5,6,7,8,9}
	resultInt := TransformInPlace(listInt, func(a int) int{
		return a * 3
	})
	fmt.Printf("Ints: %v\n", resultInt)

	// 结构体
	var listStruct = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
	}
	resultStruct := TransformInPlace(listStruct, func(e Employee) Employee{
		e.Salary += 1000
		e.Age += 1
		return e
	})
	fmt.Printf("Struct: %v\n", resultStruct)
}

type Employee struct {
	Name string
	Age int
	Vacation int
	Salary int
}