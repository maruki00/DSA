package main

import (
	Linkedlist "DSA/DataStructures/LinkedList"
	queue "DSA/DataStructures/Queue"
	stack "DSA/DataStructures/Stack"
	"fmt"
)

func StackTest() {

	s := stack.NewStack()
	s.Push(1)
	s.Push(6)
	s.Push(4)
	s.Push(3)
	fmt.Println(s, s.Pop(), s)
}

func QueueTest() {

	s := queue.NewQueue()
	s.InQueue(1)
	s.InQueue(6)
	s.InQueue(4)
	s.InQueue(3)
	fmt.Println(s)
	fmt.Println(s.DeQueue())
	fmt.Println(s.DeQueue())
	fmt.Println(s)
}

func LinkedLIstTest() {
	l := Linkedlist.NewLinkedList()
	l.Add(4)
	l.Add(40)
	l.Add(54)
	l.Print()
}

func main() {
	// StackTest()
	// QueueTest()
	LinkedLIstTest()
}
