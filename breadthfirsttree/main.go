package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySerachTree struct {
	Root *Node
}

func (bst *BinarySerachTree) insert(val int) *BinarySerachTree {
	newNode := &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
	if bst.Root == nil {
		bst.Root = newNode
		return bst
	}
	current := bst.Root
	for {
		if val == current.Value {
			return nil
		}
		if val < current.Value {
			if current.Left == nil {
				current.Left = newNode
				return bst
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return bst
			}
			current = current.Right
		}
		return bst
	}
}

func (bst *BinarySerachTree) find(val int) *Node {
	if bst.Root == nil {
		return nil
	}
	current := bst.Root
	found := false
	for current != nil && !found {
		if val < current.Value {
			current = current.Left
		} else if val > current.Value {
			current = current.Right
		} else {
			found = true
		}
	}
	if !found {
		return nil
	}
	return current
}
func (bst *BinarySerachTree) contains(val int) bool {
	if bst.Root == nil {
		return false
	}
	var current = bst.Root
	found := false
	for current != nil && !found {
		if val < current.Value {
			current = current.Left
		} else if val > current.Value {
			current = current.Right
		} else {
			return true
		}
	}
	return false
}

func (bst *BinarySerachTree) bfs() []int {
	data := make([]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, bst.Root)

	for i := 0; i < len(queue); i++ {
		node, arr := queue[0], queue[1:]
		queue = arr
		data = append(data, node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return data
}

func (bst *BinarySerachTree) Print() {
	var current = bst.Root
	queue := make([]*Node, 0)
	queue = append(queue, bst.Root)

	for i := 0; i < len(queue); i++ {
		node, arr := queue[0], queue[1:]
		fmt.Println("node", node.Value)
		fmt.Println("arr", len(arr))
		queue = arr
		if current.Left != nil {
			queue = append(queue, node.Left)
		}
		if current.Right != nil {
			queue = append(queue, node.Right)
		}
		fmt.Print(current.Value)
	}
}

func (bst *BinarySerachTree) dfsPreOrder() []int {
	data := make([]int, 0)
	traverse(bst.Root, data)
	return data
}

func traverse(node *Node, data []int) {
	data = append(data, node.Value)
	if node.Left != nil {
		traverse(node.Left, data)
	}
	if node.Right != nil {
		traverse(node.Right, data)
	}
	fmt.Println(node.Value)
}

func main() {
	node := &Node{Value: 10, Left: nil, Right: nil}
	tree := &BinarySerachTree{
		Root: node,
	}
	tree.insert(10)
	tree.insert(6)
	tree.insert(15)
	tree.insert(3)
	tree.insert(8)
	tree.insert(20)

	data := tree.dfsPreOrder()
	fmt.Println(data)

	//data := tree.bfs()
	//fmt.Println(data)
}
