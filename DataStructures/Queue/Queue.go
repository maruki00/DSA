package queue

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}
func (o *Queue) InQueue(item interface{}) {
	o.items = append([]interface{}{item}, o.items...)
}

func (o *Queue) DeQueue() interface{} {
	lastIndex := len(o.items) - 1
	item := o.items[lastIndex]
	o.items = o.items[:lastIndex]
	return item
}
