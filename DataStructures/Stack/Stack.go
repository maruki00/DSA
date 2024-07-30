package stack

import "sync"

type Stack struct {
	items []interface{}
	sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}
func (o *Stack) Push(item interface{}) {
	o.Lock()
	o.items = append(o.items, item)
	o.Unlock()
}

func (o *Stack) Pop() interface{} {
	o.Lock()
	lastIndex := len(o.items) - 1
	item := o.items[lastIndex]
	o.items = o.items[:lastIndex]
	o.Unlock()
	return item
}
