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
	if o.head == nil || o.tail == nil {
		print("First Try \n")
		o.head = node
		o.tail = node
		return
	}
	tmp := o.tail
	tmp.next = node
	o.tail = tmp.next
}

func (o *List) Delete(item interface{}) {
	if o.head == nil {
		return
	}

	if o.head.item == item {
		o.head = o.head.next
		return
	}

	// previos := o.head

	// for previos.next.next != nil {
	// 	if previos.next.item != item {
	// 		previos = previos.next
	// 	} else {
	// 		previos.next = previos.next.next
	// 	}
	// }

	previos := o.head
	for previos.next != nil {
		o.tail = previos
		if previos.next.item != item {
			previos = previos.next
		} else {
			previos.next = previos.next.next
		}

	}

	// x := o.head
	// for x.next != nil {
	// 	x = x.next
	// }
	// o.tail = x
}

func (o *List) Update(old, new interface{}) {
	tmp := o.head

	for tmp != nil {
		if tmp.item == old {
			tmp.item = new
		}
		tmp = tmp.next
	}
}

func (o *List) Print() {
	t := o.head
	for t != nil {
		fmt.Printf("value : %v - %v\n", t.item, o.tail)
		t = t.next
	}
}
