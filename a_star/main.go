package main

func main() {
	node := initNode()

	atom := &Atom{
		item:   node,
		weight: 0,
	}

	var genAtom *Atom
	var rootGenAtom *Atom
	for currAtom := atom; currAtom != nil; currAtom = currAtom.iterator() {

		if genAtom == nil {
			genAtom = currAtom
			rootGenAtom = genAtom
			continue
		}

		genAtom.next = currAtom
		genAtom = genAtom.next

		if currAtom.item.item == "n5" {
			currAtom.next = nil
			break
		}
	}
	rootGenAtom.print()
}

func initNode() *Node {
	n5 := &Node{
		item: "n5",
	}
	n3 := &Node{
		item:   "n3",
		next:   []*Node{n5},
		weight: []int{10},
	}
	n4 := &Node{
		item:   "n4",
		next:   []*Node{n3, n5},
		weight: []int{20, 60},
	}
	n2 := &Node{
		item:   "n2",
		next:   []*Node{n3},
		weight: []int{50},
	}
	n0 := &Node{
		item:   "n0",
		next:   []*Node{n2, n4, n5},
		weight: []int{10, 30, 100},
	}
	return n0
}

type Node struct {
	item   string
	next   []*Node
	weight []int
}

func (n *Node) print(atom *Atom) {


}
