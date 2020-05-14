package learn

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wedensday
)

const (
	Readable = 1 << iota
	Writable
	Exectable
)

func TestContant(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstant1(t *testing.T) {
	a := 1 //0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Exectable == Exectable)
}
