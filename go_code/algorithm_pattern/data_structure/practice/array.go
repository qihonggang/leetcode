package practice

import (
	"errors"
	"fmt"
)

type Array struct {
	data []int
	length uint
}

func NewArray(cpacity uint) *Array {
	if cpacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, cpacity, cpacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

func (this *Array) Insert(index uint, v int) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

func (this *Array) insertToTail(v int) error {
	return this.Insert(this.Len(), v)
}

func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index ragne")
	}
	v := this.data[index]
	for i := index; i < this.Len(); i-- {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Len(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}