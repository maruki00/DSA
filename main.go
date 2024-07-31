package main

import (
	searching "DSA/Algorithms/Searching"
	sorting "DSA/Algorithms/Sorting"
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
	l.Add(54)
	l.Add(54)
	l.Delete(54)
	l.Delete(54)
	l.Delete(54)
	l.Delete(40)
	l.Add(54)
	l.Add(54)
	l.Add(54)
	l.Add(54)
	l.Update(54, 99)
	l.Add(54)

	l.Add(54)

	l.Add(54)

	l.Print()

}

func InsertionSortingTest() {

	items := []int{1, 6, 3, 78, 32, 5, 7, 98, 1, 8}
	items = sorting.InsertionSort(items)
	fmt.Println(items)
}

func MergeSortingTest() {

	items1 := []int{1, 3, 6, 32, 78}
	items2 := []int{1, 6, 8, 9, 13, 34, 56, 76, 78}
	items := sorting.MergeSort(items1, items2)
	fmt.Println("Merge Sort : ", items)
}

func BucketSortingTest() {

	items1 := []int{0, 4, 5, 7, 4, 6, 1, 6, 9, 5, 4, 365, 8, 4, 3, 5, 7, 8}
	items := sorting.BucketSort(items1)
	fmt.Println(items)
}

func CustomSort() {

	items1 := []int{0, 4, 5, 7, 4, 6, 1, 6, 9, 5, 4, 365, 8, 4, 3, 5, 7, 8}
	items := sorting.CustomSort(items1)
	fmt.Println(items)
}

func BinarySearchTest() {
	items := []int{1, 4, 5, 7, 8, 9, 10, 12, 23, 34}
	res := searching.BinarySearch(items, 19)
	fmt.Println("Binary Search : ", res)
}

func main() {
	// StackTest()
	// QueueTest()
	//LinkedLIstTest()

	//InsertionSortingTest()
	//MergeSortingTest()
	// BucketSortingTest()
	// CustomSort()

	BinarySearchTest()
}
