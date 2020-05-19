package learn

import "testing"

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
