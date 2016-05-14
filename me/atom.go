package main

import (
	"fmt"
)

func initClumpTree(a *Atom) *ClumpTree {
	return &ClumpTree{
		item:     a,
		children: make([]*ClumpTree, 0),
	}
}

type ClumpTree struct {
	item     *Atom
	children []*ClumpTree
}

func (c *ClumpTree) print() {
	arr := make([]*Atom,0)
	c.printPath(arr)
}

func (c *ClumpTree) printPath(arr []*Atom) {
	arr = append(arr,c.item)
	if len(c.children)==0 {
		printArray(arr)
		return
	}
	for i:= 0;i<len(c.children);i++ {
		c.children[i].printPath(arr)
	}
}

func printArray(atoms []*Atom) {
	for i:=0;i<len(atoms);i++ {
		fmt.Print(*atoms[i]," ")
	}
	fmt.Println()
}

func (c *ClumpTree) iterator() {
	fmt.Print(c.item.arr[0:])

	for i := 0; i < len(c.children); i++ {
		c.children[i].iterator()
	}

}

func (c *ClumpTree) initChild(a *Atom) {
	child := &ClumpTree{
		item: a,
	}
	c.children = append(c.children, child)
}

func (c *ClumpTree) generateTree() {
	for i := 0; i < len(c.children); i++ {
		c.generateByPos(i)
	}
}

func (c *ClumpTree) generateByPos(pos int) {
	currClump := c.children[pos]
	for i := pos + 1; i < len(c.children); i++ {
		if !currClump.item.isExist(c.children[i].item) {
			child := initClumpTree(c.children[i].item)
			currClump.children = append(currClump.children, child)
		}
	}
	for i := 0; i < len(currClump.children); i++ {
		currClump.generateByPos(i)
	}
}

func initAtom(a []int) *Atom {
	return &Atom{
		arr: a,
	}
}

type Atom struct {
	arr []int
}

func (a *Atom) isExist(atom *Atom) bool {
	for i := 0; i < len(a.arr); i++ {
		for j := 0; j < len(atom.arr); j++ {
			if a.arr[i] == atom.arr[j] {
				return true
			}
		}
	}
	return false
}
