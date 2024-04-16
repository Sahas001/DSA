package main

import (
	"fmt"
)

type Set struct {
	integerMap map[int]bool
}

func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

func (set *Set) AddElement(m int) {
	if !set.ContainsElement(m) {
		set.integerMap[m] = true
	}
}

func (set *Set) DeleteElement(m int) {
	delete(set.integerMap, m)
}

func (set *Set) ContainsElement(m int) bool {
	var exists bool
	_, exists = set.integerMap[m]
	return exists
}

func main() {
	set := &Set{}
	set.New()
	set.AddElement(2)
	set.AddElement(1)
	set.AddElement(5)
	set.AddElement(5)
	set.AddElement(7)
	set.DeleteElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(7))
}
