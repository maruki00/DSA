package main

import (
	"errors"
	"fmt"
)

var (
	InvalidIndexError = errors.New("Invalid Index")
)

var dynamicArray []int

func Remove(index int) error {
	if index < 0 || index >= len(dynamicArray) {
		return InvalidIndexError
	}
	dynamicArray = append(dynamicArray[:index], dynamicArray[index+1:]...)
	return nil
}

func Add(item int) {
	dynamicArray = append(dynamicArray, item)
}

func GetAt(index int) (int, error) {
	if index < 0 || index >= len(dynamicArray) {
		return -1, InvalidIndexError
	}
	return dynamicArray[index], nil
}

func Update(index int, item int) error {
	if index < 0 || index >= len(dynamicArray) {
		return InvalidIndexError
	}
	dynamicArray[index] = item
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

	fmt.Println(dynamicArray, value)
}
