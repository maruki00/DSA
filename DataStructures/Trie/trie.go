package trie

import "fmt"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

func (o *Node) GetValue() int {
	return o.value
}

func NewTrie(value int) *Node {
	return &Node{value: value, Left: nil, Right: nil}
}

func printTrie(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("Val : %v\n", node.value)
	printTrie(node.Left)
	printTrie(node.Right)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value, Left: nil, Right: nil}
	}

	if node.value > value {
		node.Left = insert(node.Left, value)

	} else {
		node.Right = insert(node.Right, value)
	}
	return node

}

func minValueNode(root *Node) *Node {
	curr := root

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func remove(root *Node, value int) *Node {

	if root == nil {
		return nil
	}

	if value > root.value {
		root.Right = remove(root.Right, value)
	} else if value < root.value {
		root.Left = remove(root.Left, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := minValueNode(root.Right)
			root.value = minNode.value
			root.Right = remove(root.Right, minNode.value)
		}
	}
	return root
}

func (o *Node) Insert(value int) {

	insert(o, value)
}

func (o *Node) Remove(value int) {
	remove(o, value)
}

func (o *Node) Print() {
	printTrie(o)
}
