package set

import "sync"

type Set struct {
	items map[interface{}]bool
	sync.Mutex
}

func NewSet() *Set {
	return &Set{items: make(map[interface{}]bool)}
}

func (o *Set) Add(item interface{}) {
	o.Lock()
	defer o.Unlock()

	o.items[item] = true
}

func (o *Set) Delete(item interface{}) {
	delete(o.items, item)
}

func (o *Set) Update(old, item interface{}) {
	o.Lock()
	defer o.Unlock()
	o.items[item] = true
	delete(o.items, old)
}
