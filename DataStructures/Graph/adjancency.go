package graph

type AdjancencyGraph struct {
	nodes map[int][]int
}

func NewAdjancencyGraph() *AdjancencyGraph {
	return &AdjancencyGraph{
		nodes: make(map[int][]int, 0),
	}
}

func (this *AdjancencyGraph) Add(from, to int) {
	if _, ok := this.nodes[from]; ok == false {
		this.nodes[from] = make([]int, 0)
	}
	this.nodes[from] = append(this.nodes[from], to)
}

func (this *AdjancencyGraph) Delete(node int) {
	delete(this.nodes, node)

	for _, nodes := range this.nodes {
		for i, n := range nodes {
			if n == node {
				nodes = append(nodes[:i], nodes[i+1:]...)
			}
		}
	}

}
