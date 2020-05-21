package learn

import (
	"strconv"
	"strings"
	"testing"
)

func TestCondition(t *testing.T) {
	a := "b"
	if a == "b" {
		println("a = b")
	}

	if a = "c"; a == "c" {
		println("a = c")
	}

	if b := "1"; b == "2" {
		println("b = 2")
	} else if b == "3" {
		println("b = 3")
	} else {
		println("b = 1")
	}
}

func TestFor(t *testing.T) {
	// var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for i := 0; i < len(a); i++ {
	// 	t.Log(a[i])
	// }
	b := 0

	for {
		if b == 9 {
			return
		}
		t.Log(b)
		b++
	}
}

func TestSwicth(t *testing.T) {
	switch a := "a"; a {
	case "a":
		t.Log("a")
	case "b":
		t.Log("b")
	case "c":
		t.Log("c")
	default:
		t.Log("default")
	}

	b := "b"
	switch {
	case b != "b":
		t.Log("b != b")
	case b != "c":
		t.Log("b != c")
	case b == "d":
		t.Log("b = d")
	}
}

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, v := range a {
		t.Log(i, v)
	}
}

func TestMap(t *testing.T) {
	m := map[string]string{}
	m["1"] = "1"
	m["2"] = "2"
	delete(m, "2")
	for k, v := range m {
		t.Log(k, v)
	}
}

func TestFactory(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is exsisting", n)
	} else {
		t.Logf("%d is not exsisting", n)
	}
	mySet[3] = true
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is exsisting", n)
	} else {
		t.Logf("%d is not exsisting", n)
	}
	t.Log(len(mySet))
}

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(s)
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf-8 %s", s)
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	t.Log(len(parts))
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
