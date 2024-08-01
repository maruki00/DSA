package trees

import (
	trie "DSA/DataStructures/Trie"
	"fmt"
)

func Postorder(root *trie.Node) {
	if root == nil {
		return
	}

	Preorder(root.Left)
	Preorder(root.Right)
	fmt.Printf("value : %v\n", root.GetValue())

}
