package main

import "fmt"

type Shape interface {
	Sides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Area() int {
	panic("implement me")
}

func (s *Square) Sides() int {
	return 4
}

var _ Shape = (*Square)(nil)

func main() {
	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}
