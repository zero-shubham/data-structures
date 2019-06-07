package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func initNode(d int) *node {
	tmpNode := node{data: d, left: nil, right: nil}
	return &tmpNode
}

func insertToNode(root *node, d int) *node {

	if root == nil {
		root = initNode(d)
		return root
	}
	tmp := &root

	for true {
		if (*tmp).data > d {
			if (*tmp).left != nil {
				(tmp) = &((*tmp).left)
			} else {
				(*tmp).left = initNode(d)
				break
			}
		} else {
			if (*tmp).right != nil {
				(tmp) = &((*tmp).right)
			} else {
				(*tmp).right = initNode(d)
				break
			}
		}
	}

	fmt.Println(tmp, root)
	return root
}

func printTreeInorder(ptr *node) {
	if ptr == nil {
		return
	}
	printTreeInorder(ptr.left)
	fmt.Println(ptr.data)
	printTreeInorder(ptr.right)
}

func main() {
	var inp int
	var root *node
	for i := 1; i <= 5; i++ {
		fmt.Println("\nEnter element:")
		fmt.Scanf("%d", &inp)
		if i == 1 {
			root = insertToNode(root, inp)
		} else {
			root = insertToNode(root, inp)
		}
	}

	printTreeInorder(root)
}
