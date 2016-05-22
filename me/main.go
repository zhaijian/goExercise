package main

func main() {

	c := [][]int{[]int{1, 2, 3}, []int{3, 8}, []int{4, 6}, []int{5, 7}, []int{4, 7}, []int{1, 2, 8}, []int{4, 8, 9}, []int{12, 8, 9, 11}}
//	c := [][]int{[]int{1, 2, 3}}
	for i := 0; i < len(c); i++ {

		a := initAtom(c[i])
		//每一个集合依次和他们以后的节点做互斥，并创建一棵集合树
		clumpTree := initClumpTree(a)

		//首先构造树的第一级子树
		for j := i + 1; j < len(c); j++ {
			at := initAtom(c[j])

			//如果不互斥，则不加入这颗树
			if a.isExist(at) {
				continue
			}
			//创建这颗树
			clumpTree.initChild(at)
		}
		//生成每棵树的2-N级子树
		clumpTree.generateTree()
		clumpTree.print()
	}

}
