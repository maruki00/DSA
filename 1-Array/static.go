package main

import (
	"errors"
	"fmt"
)

var (
	InvalidIndexError = errors.New("Invalid Index")
)

var staticArray []int

func remove(index int) error {
	if index < 0 || index >= len(staticArray) {
		return InvalidIndexError
	}
	staticArray = append(staticArray[:index], staticArray[index+1:]...)
	return nil
}

func add(item int) {
	staticArray = append(staticArray, item)
}

func getAt(index int) (int, error) {
	if index < 0 || index >= len(staticArray) {
		return -1, InvalidIndexError
	}
	return staticArray[index], nil
}

func Update(index int, item int) error {
	if index < 0 || index >= len(staticArray) {
		return InvalidIndexError
	}
	staticArray[index] = item
	return nil
}

func main() {
	add(1)
	add(1)
	add(1)
	add(1)
	add(3)
	fmt.Println(staticArray)
}
