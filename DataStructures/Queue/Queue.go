package queue

import (
	"fmt"
	"sync"
)

type Queue struct {
	items []interface{}
	sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		items: []interface{}{},
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

	fmt.Println(o.items)
	item := o.items[lastIndex]
	o.items = o.items[:lastIndex]
	o.Unlock()
	return item
}
