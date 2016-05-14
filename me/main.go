package main

func main() {

//	c := [][]int{[]int{1, 2, 3}, []int{3, 8}, []int{4, 6}, []int{5, 7}, []int{4, 7}, []int{1, 2, 8}, []int{4, 8, 9}, []int{12, 8, 9, 11}}
	c := [][]int{[]int{1, 2, 3}}
	for i := 0; i < len(c); i++ {

		a := initAtom(c[i])
		clumpTree := initClumpTree(a)

		for j := i + 1; j < len(c); j++ {
			at := initAtom(c[j])

			if a.isExist(at) {
				continue
			}
			clumpTree.initChild(at)
		}
		clumpTree.generateTree()
		clumpTree.print()
	}

}
