package main

import (
	"fmt"
	//lint:ignore ST1001 because this way is more readable
	. "github.com/justseyit/dijkstra/dijkstra"
)

func main() {

	//For 10 vertex, you must enter 11
	var g Graph = NewGraph(11)

	g.AddEdge(1, 3, 450)
	g.AddEdge(1, 4, 154)
	g.AddEdge(2, 3, 403)
	g.AddEdge(3, 4, 388)
	g.AddEdge(3, 6, 319)
	g.AddEdge(3, 7, 262)
	g.AddEdge(3, 10, 1001)
	g.AddEdge(4, 5, 345)
	g.AddEdge(4, 6, 183)
	g.AddEdge(5, 6, 335)
	g.AddEdge(5, 8, 460)
	g.AddEdge(7, 9, 345)
	g.AddEdge(8, 9, 609)
	g.AddEdge(9, 10, 523)

	var d *Dijkstra = &Dijkstra{}

	path, dist := d.ShortestPath(g, 1, 10)

	var s string

	for i := len(path) - 1; i >= 0; i-- {
		if i != 0 {
			s += fmt.Sprintf("%s -> ", path[i])
		} else {
			s += path[i].String()
		}
	}
	fmt.Println("Start: ", path[len(path)-1])
	fmt.Println("End: ", path[0])
	fmt.Println("Shortest path:", s)
	fmt.Printf("Minimum distance: %d meters\n", dist)
}
