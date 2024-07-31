package trie

type Node struct {
	value interface{}
	Left  *Node
	Right *Node
}
