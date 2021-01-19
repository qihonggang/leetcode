package main

import "fmt"

var _ Painter = (*Label)(nil)
var _ Clicker = (*Button)(nil)
var _ Painter = (*Button)(nil)
var _ Clicker = (*ListBox)(nil)
var _ Painter = (*ListBox)(nil)

type Widget struct {
	X, Y int
}

type Label struct {
	Widget
	Text string
}

func (l Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &l, l.Text )
}

type Button struct {
	Label
}

func (b Button) Paint() {
	fmt.Printf("Button.Paint(%s)\n", b.Text)
}

func (b Button) Click() {
	fmt.Printf("Button.Paint(%s)\n", b.Text)
}

type ListBox struct {
	Widget
	Texts []string
	Index int
}

func (l ListBox) Paint() {
	fmt.Printf("ListBox.Paint(*q)\n", l.Texts)
}

func (l ListBox) Click() {
	fmt.Printf("ListBox.Click(*q)\n", l.Texts)
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

func main(){
	label := Label{Widget{10, 10}, "label"}
	button := Button{
		Label{Widget{10, 70}, "OK"},
	}
	listBox := ListBox{
		Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"},
		0,
	}

	for _, painter := range []Painter{label, listBox, button} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}
}