package backtracking

import (
	stack "DSA/DataStructures/Stack"
	trie "DSA/DataStructures/Trie"
)

func Backtracking(root *trie.Node, path stack.Stack) bool {
	if root == nil || root.GetValue() == 0 {
		return false
	}
	path.Push(root.GetValue())
	if root.Left == nil && root.Right == nil {
		return true
	}

	if Backtracking(root.Left, path) == true {
		return true
	}

	if Backtracking(root.Right, path) == true {
		return true
	}
	path.Pop()
	return false
}
