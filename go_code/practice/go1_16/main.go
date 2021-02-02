package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed helloworld/*
var f embed.FS

func main(){
	data1, _ := f.ReadFile("helloworld/hello1.txt")
	fmt.Printf(string(data1))

	data2, _ := f.ReadFile("helloworld/hello2.txt")
	fmt.Printf(string(data2))
}


