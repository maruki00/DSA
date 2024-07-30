package main

import "fmt"

type Static struct {
	lenght      int
	staticArray []interface{}
}

func NewStaticArray(lenght int) (*Static, error) {
	if lenght <= 0 {
		return nil, fmt.Errorf("lenght should be great then 0")
	}
	return &Static{
		lenght:      lenght,
		staticArray: make([]interface{}, lenght),
	}, nil
}

func (o *Static) SetOrUpdate(index int, item interface{}) error {
	if index < 0 || index >= o.lenght {
		return fmt.Errorf("invalid index %d", index)
	}
	o.staticArray[index] = item
	return nil
}

func (o Static) Get(index int) (interface{}, error) {
	if index < 0 || index >= o.lenght {
		return nil, fmt.Errorf("invalid index %d", index)
	}
	return o.staticArray[index], nil
}
