package main

import (
	"errors"
	"fmt"
)

var (
	InvalidIndexError = errors.New("Invalid Index")
)

var staticArray []int

func Remove(index int) error {
	if index < 0 || index >= len(staticArray) {
		return InvalidIndexError
	}
	staticArray = append(staticArray[:index], staticArray[index+1:]...)
	return nil
}

func Add(item int) {
	staticArray = append(staticArray, item)
}

func GetAt(index int) (int, error) {
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
	Add(1)
	Add(1)
	Add(1)
	Add(1)
	Add(3)
	err := Remove(4)
	if err != nil {
		panic(err.Error())
	}
	value, err := GetAt(3)
	if err != nil {
		panic(err.Error())
	}
	err = Update(3, 4)

	fmt.Println(staticArray, value)
}
