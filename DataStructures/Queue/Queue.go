package queue

import (
	"sync"
)

type Queue struct {
	items []interface{}
	sync.Mutex
}

func (o *Queue) IsEmpty() bool {
	return len(o.items) == 0
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}
func (o *Queue) InQueue(item interface{}) {
	o.Lock()
	o.items = append([]interface{}{item}, o.items...)
	o.Unlock()
}

func (o *Queue) DeQueue() interface{} {
	o.Lock()
	lastIndex := len(o.items) - 1
	item := o.items[lastIndex]
	o.items = o.items[:lastIndex]
	o.Unlock()
	return item
}
