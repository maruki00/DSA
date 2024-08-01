package trees

import (
	queue "DSA/DataStructures/Queue"
	trie "DSA/DataStructures/Trie"
	"fmt"
)

var (
	q queue.Queue = *queue.NewQueue()
)

func BFS(root *trie.Node) {
	q.InQueue(root)

	for !q.IsEmpty() {
		tmp := q.DeQueue()
		fmt.Printf("value : %v\n", tmp.GetValue())
		if tmp.Left != nil {
			q.InQueue(tmp.Left)
		}
		if tmp.Right != nil {
			q.InQueue(tmp.Right)
		}
	}
}
