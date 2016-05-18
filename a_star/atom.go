package main

import (
	"fmt"
)

type Atom struct {
	item   *Node
	next   *Atom
	weight int
}

func (a *Atom) iterator() *Atom {
	nextAtom := a.next
	children := a.item.next
	for i := 0; i < len(children); i++ {
		child := children[i]
		weight := a.item.weight[i]
		atom := &Atom{
			item:   child,
			weight: a.weight + weight,
		}
		if nextAtom == nil {
			nextAtom = atom
			continue
		}
		nextAtom = nextAtom.sortInsert(atom)
	}
	return nextAtom
}

func (a *Atom) sortInsert(newAtom *Atom) *Atom {

	if a.weight >= newAtom.weight {
		newAtom.next = a
		return newAtom
	}

	if a.next == nil {
		a.next = newAtom
		return a
	}
	for curr := a; curr != nil; curr = curr.next {
		if curr.weight >= newAtom.weight {
			continue
		}

		currNext := curr.next
		if currNext == nil {
			curr.next = newAtom
			return a
		}

		if newAtom.weight <= currNext.weight {
			curr.next = newAtom
			newAtom.next = currNext
			return a
		}
	}
	return a
}

func (a *Atom) print() {
	if a == nil {
		return
	}
	for curr := a; curr != nil; curr = curr.next {
		fmt.Print("(item ", curr.item.item, " weight ", curr.weight, ") ")
	}
	fmt.Println()
}

func (a *Atom) pop() *Atom {
	curr := a
	a = a.next
	curr.next = nil
	return curr
}
