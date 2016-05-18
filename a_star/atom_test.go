package main

import (
	"log"
	"testing"
)

func TestAtom(t *testing.T) {
	node := initNode()
	//	log.Println(node.next[0].item)
	atom := &Atom{
		item:   node,
		weight: 0,
	}
	atom1 := &Atom{
		item:   node.next[0],
		weight: 10,
	}
	atom2 := &Atom{
		item:   node.next[1],
		weight: 5,
	}
	atom3 := &Atom{
		item:   node.next[2],
		weight: 7,
	}
	atom4 := &Atom{
		item:   node.next[0].next[0],
		weight: 1,
	}
	atom = atom.sortInsert(atom1)
	atom = atom.sortInsert((atom2))
	atom = atom.sortInsert((atom3))
	atom = atom.sortInsert((atom4))

	log.Println("atom0", atom.item, atom.weight)
	log.Println("atom1", atom.next.item, atom.next.weight)
	log.Println("atom2", atom.next.next.item, atom.next.next.weight)
	log.Println("atom3", atom.next.next.next.item, atom.next.next.next.weight)
	log.Println("atom4", atom.next.next.next.next.item, atom.next.next.next.next.weight)

	log.Println("root", atom.weight, atom.next)
	a := atom.pop()
	log.Println("pop", a.weight, "root", atom.weight, "next", atom.next)
}
