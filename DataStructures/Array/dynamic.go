package array

import (
	"errors"
	"sync"
)

var (
	ErrInvalidIndex = errors.New("invalid index")
)

type Dynamic struct {
	items []interface{}
	sync.Mutex
}

func (o *Dynamic) Remove(index int) error {
	if index < 0 || index >= len(o.items) {
		return ErrInvalidIndex
	}
	o.items = append(o.items[:index], o.items[index+1:]...)
	return nil
}

func (o *Dynamic) Add(item int) {
	o.items = append(o.items, item)
}

func (o *Dynamic) GetAt(index int) (interface{}, error) {
	if index < 0 || index >= len(o.items) {
		return -1, ErrInvalidIndex
	}
	return o.items[index], nil
}

func (o *Dynamic) Update(index int, item int) error {
	if index < 0 || index >= len(o.items) {
		return ErrInvalidIndex
	}
	o.items[index] = item
	return nil
}
