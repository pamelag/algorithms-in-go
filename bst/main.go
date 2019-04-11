package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}
type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) BFS() []int {
	node := bst.root
	data := make([]int, 0)
	queue := make([]*Node, 0)

	queue = append(queue, node)
	for len(queue) > 0 {
		node = queue[0]
		data = append(data, node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return data
}

func (bst *BinarySearchTree) DFS() []int {
	data := []int{}
	bst.traverse(bst.root, data)
	return data
}

func (bst *BinarySearchTree) traverse(node *Node, data []int) {
	if node.left != nil {
		bst.traverse(node.left, data)
	}
	data = append(data, node.value)
	if node.right != nil {
		bst.traverse(node.right, data)
	}
}

func createBST(val int) *BinarySearchTree {
	root := &Node{
		value: val,
		left:  nil,
		right: nil,
	}
	bst := &BinarySearchTree{
		root: root,
	}
	return bst
}

func main() {
	fmt.Println("BFS")

	bst := createBST(10)
	bst.root.left.value = 7
	bst.root.right.value = 15
	bst.root.left.right.value = 9

}
