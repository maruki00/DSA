package datastructures

var stack []interface{}

func Push(item interface{}) {
	stack = append(stack, item)
}

func Pop() interface{} {
	lastIndex := len(stack) - 1
	item := stack[lastIndex]
	stack = stack[:lastIndex-2]
	return item
}
