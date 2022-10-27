package dijkstra

import "math"

type Dijkstra struct{}

func (d *Dijkstra) ShortestPath(g *Graph, from, to int) ([]Vertex, int) {
	distance := make([]int, g.GetNumOfVertices())
	visited := make([]bool, g.GetNumOfVertices())
	previous := make([]int, g.GetNumOfVertices())
	for i := 0; i < g.GetNumOfVertices(); i++ {
		distance[i] = math.MaxInt32
		visited[i] = false
		previous[i] = -1
	}
	distance[from] = 0
	for i := 0; i < g.GetNumOfVertices(); i++ {
		u := d.minDistance(distance, visited)
		visited[u] = true
		for v := 0; v < g.GetNumOfVertices(); v++ {
			if !visited[v] && g.GetWeight(u, v) != 0 && distance[u] != math.MaxInt32 && distance[u]+g.GetWeight(u, v) < distance[v] {
				distance[v] = distance[u] + g.GetWeight(u, v)
				previous[v] = u
			}
		}
	}

	return d.getPath(previous, to), distance[to]
}

func (d *Dijkstra) minDistance(distance []int, visited []bool) int {
	min := math.MaxInt32
	minIndex := 0
	for v := 0; v < len(distance); v++ {
		if !visited[v] && distance[v] <= min {
			min = distance[v]
			minIndex = v
		}
	}
	return minIndex
}

func (d *Dijkstra) getPath(previous []int, to int) []Vertex {
	var path []Vertex
	for v := to; v != -1; v = previous[v] {
		path = append(path, Vertex(v))
	}
	return path
}
