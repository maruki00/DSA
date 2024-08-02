package graph

type AdjancencyGraph struct {
	nodes map[int][]int
}

func NewAdjancencyGraph() *AdjancencyGraph {
	return &AdjancencyGraph{
		nodes: make(map[int][]int, 0),
	}
}
