package main

import (
	"fmt"
	"math"
	"sort"
)

type Vertex struct {
	init_number int
	x           float64
	y           float64
}

type Edge struct {
	dist float64
	v1   Vertex
	v2   Vertex
}

type ByDist []Edge

func (a ByDist) Len() int           { return len(a) }
func (a ByDist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDist) Less(i, j int) bool { return a[i].dist < a[j].dist }

func is_equal_sets(list_a, list_b []Vertex) bool {
	list_a_table := make(map[Vertex]int, 20)

	for _, v1 := range list_a {
		list_a_table[v1] = 1
	}

	for _, v2 := range list_b {
		_, status := list_a_table[v2]

		if !status {
			return false
		} else {
			list_a_table[v2] = 2
		}
	}

	for _, v1 := range list_a {
		value, _ := list_a_table[v1]

		if value == 1 {
			return false
		}
	}

	return true
}

func Kruskal(list_vertexes []Vertex, list_edges []Edge, amount_vertexes int) float64 {
	table_vertexes := make(map[Vertex][]Vertex, 20)
	var list_min_dist_edges []Edge

	for _, vertex := range list_vertexes {
		table_vertexes[vertex] = append(table_vertexes[vertex], vertex)
	}

	count_edges := 0

	for _, edge := range list_edges {
		v1_set, _ := table_vertexes[edge.v1]
		v2_set, _ := table_vertexes[edge.v2]

		if !is_equal_sets(v1_set, v2_set) {
			list_min_dist_edges = append(list_min_dist_edges, edge)

			for _, vertex_2 := range v2_set {
				for _, vertex_1 := range v1_set {
					table_vertexes[vertex_2] = append(table_vertexes[vertex_2], vertex_1)
					table_vertexes[vertex_1] = append(table_vertexes[vertex_1], vertex_2)
				}
			}

			count_edges++
		}

		if count_edges == amount_vertexes-1 {
			break
		}
	}

	var ans float64

	for _, edge := range list_min_dist_edges {
		ans += edge.dist
	}
	return ans
}

func main() {
	var amount_vertexes int
	var list_vertexes []Vertex
	var list_edges []Edge

	fmt.Scan(&amount_vertexes)

	for i := 0; i < amount_vertexes; i++ {
		var x, y float64

		fmt.Scan(&x)
		fmt.Scan(&y)

		list_vertexes = append(list_vertexes, Vertex{i, x, y})
	}

	for i := 0; i < amount_vertexes; i++ {
		for j := i + 1; j < amount_vertexes; j++ {
			x := list_vertexes[i].x - list_vertexes[j].x
			y := list_vertexes[i].y - list_vertexes[j].y
			dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

			list_edges = append(list_edges, Edge{dist, list_vertexes[i], list_vertexes[j]})
		}
	}

	sort.Sort(ByDist(list_edges))

	fmt.Printf("%.2f\n", Kruskal(list_vertexes, list_edges, amount_vertexes))
}