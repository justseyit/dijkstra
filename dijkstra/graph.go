package dijkstra

import "fmt"

type Graph struct {
	adjacencyMatrix [][]int
	numOfVertices   int
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.adjacencyMatrix[from][to] = weight
}

func NewGraph(numOfVertices int) *Graph {
	g := &Graph{
		numOfVertices: numOfVertices,
	}
	g.adjacencyMatrix = make([][]int, numOfVertices)
	for i := range g.adjacencyMatrix {
		g.adjacencyMatrix[i] = make([]int, numOfVertices)
	}
	return g
}

func (g *Graph) GetAdjacencyMatrix() [][]int {
	return g.adjacencyMatrix
}

func (g *Graph) GetNumOfVertices() int {
	return g.numOfVertices
}

func (g *Graph) GetWeight(from, to int) int {
	return g.adjacencyMatrix[from][to]
}

func (g *Graph) String() string {
	var s string
	for i := 0; i < g.numOfVertices; i++ {
		for j := 0; j < g.numOfVertices; j++ {
			s += fmt.Sprintf("%d ", g.adjacencyMatrix[i][j])
		}
		s += " "
	}
	return s
}
