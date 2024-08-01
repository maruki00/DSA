package trees

import (
	trie "DSA/DataStructures/Trie"
	"fmt"
)

func Preorder(root *trie.Node) {
	if root == nil {
		return
	}

	fmt.Printf("value : %v\n", root.GetValue())
	Preorder(root.Left)
	Preorder(root.Right)
}
