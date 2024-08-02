package graph

type MatrixGraph struct {
	nodes [][]int
}

func NewGraph(c, r int) *MatrixGraph {
	nodes := make([][]int, r, 0)
	for i := range r {
		nodes[i] = make([]int, c, 0)
	}
	return &MatrixGraph{
		nodes: nodes,
	}
}

func (this *MatrixGraph) Insert(from, to int) {
	this.nodes[from][to] = 1
}

func (this *MatrixGraph) Delete(from, to int) {
	this.nodes[from][to] = 0
}
