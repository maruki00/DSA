package heap

type Heap struct {
	items []int
}

func (this *Heap) PushItem(item int) {
	this.items = append(this.items, item)
}
func (this *Heap) PopItem() int {
	l := len(this.items) - 1
	item := this.items[l]
	this.items = this.items[:l]
	return item
}

/*
|---> parent Node = int(index / 2)
|---> left  Node = (index*2)
|---> right Node = (index*2) +1
*/
func NewHeap() *Heap {
	return &Heap{
		items: []int{0},
	}
}

func (this *Heap) Count() int {
	return len(this.items)
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
	if this.Count() == 1 {
		return -1
	}
	if this.Count() == 2 {
		return this.PopItem()
	}

	res := this.items[1]
	this.items[1] = this.PopItem()

	i := 1
	for 2*i < this.Count() {
		if 2*i+1 < this.Count() &&
			this.items[2*i+1] < this.items[2*i] &&
			this.items[i] > this.items[2*i+1] {
			this.items[i], this.items[2*i+1] = this.items[2*i+1], this.items[i]
			i = 2*i + 1
		} else if this.items[i] > this.items[2*i] {
			this.items[i], this.items[2*i] = this.items[2*i], this.items[i]
			i = 2 * i
		} else {
			break
		}
	}
	return res
}
