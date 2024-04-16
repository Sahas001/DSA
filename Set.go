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

func (set *Set) IntersectionSet(secondSet *Set) *Set {
	intersectionSet := &Set{}
	intersectionSet.New()
	for value := range set.integerMap {
		if secondSet.ContainsElement(value) {
			intersectionSet.AddElement(value)
		}
	}
	return intersectionSet
}

func (set *Set) UnionSet(secondSet *Set) *Set {
	unionSet := &Set{}
	unionSet.New()
	var value int
	for value = range set.integerMap {
		unionSet.AddElement(value)
	}
	for value = range secondSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}

func main() {
	set := &Set{}
	var anotherSet *Set
	set.New()
	set.AddElement(2)
	set.AddElement(1)
	set.AddElement(5)
	set.AddElement(5)
	set.AddElement(7)
	set.DeleteElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(7))
	anotherSet = &Set{}
	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)
	fmt.Println(set.IntersectionSet(anotherSet))
	fmt.Println(set.UnionSet(anotherSet))
}
