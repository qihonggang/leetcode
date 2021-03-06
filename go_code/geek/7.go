package main

import "fmt"

type IceCreamMaker interface {
	Hello()
}

type Ben struct {
	//id int
	name string
}

func (b *Ben) Hello() {
	fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
}

//type Jerry struct {
//	name string
//}

type Jerry struct {
	field1 *[5]byte
	field2 int
}

func (j *Jerry) Hello() {
	fmt.Printf("Jerry says, \"Hello my name is %s\"\n", j.field1)
}

func main() {
	var ben = &Ben{name: "Ben"}
	var jerry = &Jerry{}
	var maker IceCreamMaker = ben

	var loop0, loop1 func()

	loop0 = func() {
		maker = ben
		go loop1()
	}

	loop1 = func() {
		maker = jerry
		go loop0()
	}

	go loop0()

	for  {
		maker.Hello()
	}
}