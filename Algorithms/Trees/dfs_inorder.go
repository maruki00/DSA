package trees

import (
	trie "DSA/DataStructures/Trie"
	"fmt"
)

func Ireorder(root *trie.Node) {
	if root == nil {
		return
	}

	Preorder(root.Left)
	fmt.Printf("value : %v\n", root.GetValue())
	Preorder(root.Right)
}
