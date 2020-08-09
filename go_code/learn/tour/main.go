package main

import (
	"github.com/go-programing-tour-book/tour/cmd"
	"log"
)

func main(){
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
