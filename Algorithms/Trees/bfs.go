package trees

import (
	trie "DSA/DataStructures/Trie"
	"fmt"
)

type Queue struct {
	items []*trie.Node
}

func (o *Queue) InQueue(node *trie.Node) {
	o.items = append([]*trie.Node{node}, o.items...)
}

func (o *Queue) Items() []*trie.Node {
	return o.items
}

func (o *Queue) IsEmpty() bool {
	return len(o.items) == 0
}

func (o *Queue) DeQueue() *trie.Node {
	l := len(o.items) - 1
	item := o.items[l]
	o.items = o.items[:l]
	return item

}

func BFS(root *trie.Node) {
	level := 1
	q := &Queue{items: make([]*trie.Node, 0)}
	if root != nil {
		q.InQueue(root)
	}
	for !q.IsEmpty() {
		fmt.Println("---------------level : ", level, "---------------------")

		for range len(q.Items()) {

			tmp := q.DeQueue()
			fmt.Printf("value : %v\n", tmp.GetValue())

			if tmp.Left != nil {
				q.InQueue(tmp.Left)
			}
			if tmp.Right != nil {
				q.InQueue(tmp.Right)
			}
		}
		level++
	}
}
