package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}
	return result
}

func main(){
	square := func(x int) int {
		return x * x
	}
	nums := []int{1,2,3,4}

	squared_arr := Map(nums, square)
	fmt.Println(squared_arr)

	upcase := func(s string) string{
		return strings.ToUpper(s)
	}
	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := Map(strs, upcase)
	fmt.Println(upstrs)
}