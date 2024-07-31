package trie

import "fmt"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

func NewTrie(value int) *Node {
	return &Node{value: value, Left: nil, Right: nil}
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value, Left: nil, Right: nil}
	}

	if node.value < value {
		node.Left = insert(node.Left, value)

	} else {
		node.Right = insert(node.Right, value)
	}
	return node

}
func (o *Node) Insert(root *Node, value int) {

	insert(root, value)
}

func printTrie(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("Val : %v\n", node.value)
	print(node.Left)
	print(node.Right)
}
func (o *Node) Print() {
	printTrie(o)
}
