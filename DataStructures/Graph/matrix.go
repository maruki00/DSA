package graph

type Graph struct {
	nodes [][]int
}

func NewGraph(c, r int) *Graph {
	nodes := make([][]int, r, 0)
	for i := range r {
		nodes[i] = make([]int, c, 0)
	}
	return &Graph{
		nodes: nodes,
	}

}

func (this *Graph) Insert(from, to int) {
	this.nodes[from][to] = 1
}

func (this *Graph) Delete(from, to int) {
	this.nodes[from][to] = 0
}
