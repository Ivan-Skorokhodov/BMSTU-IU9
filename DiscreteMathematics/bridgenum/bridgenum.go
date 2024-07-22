package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(v, prev_v, time int, list_adjacencies map[int][]int, visited *[]bool, up, in *[]int, count *int) {
	(*visited)[v] = true
	time++
	(*in)[v] = time
	(*up)[v] = time

	//fmt.Printf("%d %d\n", v, (*up)[v])

	for _, new_v := range list_adjacencies[v] {
		if new_v == prev_v {
			continue
		}

		if !(*visited)[new_v] {
			dfs(new_v, v, time, list_adjacencies, visited, up, in, count)
			(*up)[v] = min((*up)[v], (*up)[new_v])
		} else {
			(*up)[v] = min((*up)[v], (*in)[new_v])
		}
	}

	if (*up)[v] >= (*in)[v] {
		*count++
	}

}

func main() {
	var amount_vertexes, amount_edges int
	var visited []bool
	var up, in []int
	list_adjacencies := make(map[int][]int, 20)

	fmt.Scan(&amount_vertexes)
	fmt.Scan(&amount_edges)

	for i := 0; i < amount_vertexes; i++ {
		visited = append(visited, false)
		up = append(up, 0)
		in = append(in, 0)
	}

	for i := 0; i < amount_edges; i++ {
		var v1, v2 int

		fmt.Scan(&v1)
		fmt.Scan(&v2)

		list_adjacencies[v1] = append(list_adjacencies[v1], v2)
		list_adjacencies[v2] = append(list_adjacencies[v2], v1)
	}

	count := 0
	have_false := true

	for have_false {

		have_false = false

		for v, status := range visited {
			if !status {
				new_count := -1
				dfs(v, -1, 0, list_adjacencies, &visited, &up, &in, &new_count)

				count += new_count
				have_false = true
			}
		}
	}
	/*
		fmt.Println(list_adjacencies)
		fmt.Println(in)
		fmt.Println(up)
	*/
	fmt.Println(count)
}