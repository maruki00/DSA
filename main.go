package main

import (
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
func main() {
	StackTest()
}
