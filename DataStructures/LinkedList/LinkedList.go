package Linkedlist

import "fmt"

type node struct {
	item interface{}
	next *node
}

type List struct {
	head *node
	tail *node
}

func NewLinkedList() *List {
	return &List{nil, nil}
}

func (o *List) Add(item interface{}) {
	node := &node{item: item, next: nil}
	tmp := o.tail
	tmp.next = node
	o.tail = tmp.next
}

func (o *List) Delete(item interface{}) {
	tmp := o.head

	for tmp.next.next != nil {
		if tmp.next.item == item {
			tmp = tmp.next.next
		}
	}

	o.tail = tmp.next.next
}

func (o *List) Update(item interface{}) {
	tmp := o.head

	for tmp != nil {
		if tmp.item == item {
			tmp.item = item
		}
	}
}

func (o *List) Print() {
	t := o.head
	for t != nil {
		fmt.Printf("value : %v\n", t.item)
	}
}
