package main

import "fmt"

func get_min_edge(list_added_vertexes, bin_list_added_vertexes []int, table_vertexes map[int][][]int) []int {
	min_edge := []int{-1, 1000000}

	for _, start_vertex := range list_added_vertexes {
		for _, elem := range table_vertexes[start_vertex] {

			if bin_list_added_vertexes[elem[0]] != 1 && elem[1] < min_edge[1] {
				copy(min_edge, elem)
			}

		}

	}

	return min_edge
}

func Prim(amount_vertexes int, table_vertexes map[int][][]int) int {
	var bin_list_added_vertexes []int
	for i := 0; i < amount_vertexes; i++ {
		bin_list_added_vertexes = append(bin_list_added_vertexes, 0)
	}
	bin_list_added_vertexes[0] = 1

	var list_added_vertexes []int
	list_added_vertexes = append(list_added_vertexes, 0)

	ans := 0
	count := 1

	for count < amount_vertexes {
		edge := get_min_edge(list_added_vertexes, bin_list_added_vertexes, table_vertexes)

		bin_list_added_vertexes[edge[0]] = 1
		list_added_vertexes = append(list_added_vertexes, edge[0])
		ans += edge[1]
		count++
	}

	return ans
}

func main() {
	var amount_vertexes int
	var amount_edges int

	table_vertexes := make(map[int][][]int, 20)

	fmt.Scan(&amount_vertexes)
	fmt.Scan(&amount_edges)

	for i := 0; i < amount_edges; i++ {
		var v1, v2, dist int

		fmt.Scan(&v1)
		fmt.Scan(&v2)
		fmt.Scan(&dist)

		table_vertexes[v1] = append(table_vertexes[v1], []int{v2, dist})
		table_vertexes[v2] = append(table_vertexes[v2], []int{v1, dist})
	}

	fmt.Println(Prim(amount_vertexes, table_vertexes))

}