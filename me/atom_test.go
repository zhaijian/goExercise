package main

import "testing"

func TestAtom(t *testing.T) {
	oldAtom := &Atom{[]int{1, 2, 5, 8}}
	newAtom := &Atom{[]int{7, 2}}
	println(oldAtom.isExist(newAtom))
}

func TestInit(t *testing.T) {
	str := []int{1, 2, 4, 9}
	a := initAtom(str)
	println(a.arr[3])
	tree := initClumpTree(a)
	println(tree.item.arr[3])
	println(len(tree.children))
}
