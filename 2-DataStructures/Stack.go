package datastructures

type Stack struct {
	items []interface{}
}

func New() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}
func (o *Stack) Push(item interface{}) {
	o.items = append(o.items, item)
}

func (o *Stack) Pop() interface{} {
	lastIndex := len(o.items) - 1
	item := o.items[lastIndex]
	o.items = o.items[:lastIndex-2]
	return item
}
