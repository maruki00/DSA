package heap

type Heap struct {
	items []int
}

/*
parent Node = int(index / 2)
left  Node = (index*2)
right Node = (index*2) +1
*/
func NewHeap() *Heap {
	return &Heap{
		items: []int{0},
	}
}

func (this *Heap) Push(value int) {
	this.items = append(this.items, value)
	index := len(this.items) - 1
	for this.items[index] < this.items[int(index/2)] {
		this.items[index], this.items[int(index/2)] = this.items[int(index/2)], this.items[index]
		index = int(index / 2)
	}

}

func (this *Heap) Pop() int {

	return 0
}
