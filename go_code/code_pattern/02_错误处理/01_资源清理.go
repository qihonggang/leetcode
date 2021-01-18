package main

import (
	"io"
	"log"
	"os"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r, err := os.Open("a")
	if err != nil {
		log.Fatalf("err opening 'a'\n")
	}
	defer Close(r)

	r1, err := os.Open("b")
	if err != nil {
		log.Fatalf("error opening 'b'\n")
	}
	defer Close(r1)
}


