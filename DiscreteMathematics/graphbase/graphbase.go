package main

import "fmt"

func get_base(amount_cond_orgraph_vertexes int, list_cond_orgraph_adjacencies map[int][]int, list_bindings [][]int) []int {

	var moveable []int
	for i := 0; i < amount_cond_orgraph_vertexes; i++ {
		moveable = append(moveable, 0)
	}

	for i := 0; i < amount_cond_orgraph_vertexes; i++ {
		mini_list, status := list_cond_orgraph_adjacencies[i]

		if status {
			for _, v := range mini_list {
				moveable[v] = 1
			}
		}
	}

	var list_base_vertexes []int

	for i := 0; i < amount_cond_orgraph_vertexes; i++ {
		if moveable[i] == 0 {
			list_base_vertexes = append(list_base_vertexes, list_bindings[i][0])
		}
	}

	return list_base_vertexes
}

func make_cond_orgraph(amount_vertexes int, list_bindings [][]int, list_adjacencies map[int][]int) map[int][]int {
	list_cond_orgraph_adjacencies := make(map[int][]int, 20)

	for i, elem1 := range list_bindings {

		var bin_can_visit []int

		for i := 0; i < amount_vertexes; i++ {
			bin_can_visit = append(bin_can_visit, 0)
		}

		for _, v1 := range elem1 {
			for _, v2 := range list_adjacencies[v1] {
				bin_can_visit[v2] = 1
			}
		}

		for j, elem2 := range list_bindings {

			if i == j {
				continue
			}

			for _, v := range elem2 {
				if bin_can_visit[v] == 1 {
					list_cond_orgraph_adjacencies[i] = append(list_cond_orgraph_adjacencies[i], j)
					break
				}
			}
		}
	}

	return list_cond_orgraph_adjacencies
}

func find_all_strong_bindings(amount_vertexes int, list_adjacencies, inverse_list_adjacencies map[int][]int) [][]int {
	var all_visited []int

	for i := 0; i < amount_vertexes; i++ {
		all_visited = append(all_visited, 0)
	}

	var list_bindings [][]int
	have_not_visited := true
	count_bindings := 0

	for have_not_visited {

		have_not_visited = false

		for i, status := range all_visited {

			if status == 0 {
				list_bindings = append(list_bindings, find_one_strong_binding(i, amount_vertexes, list_adjacencies, inverse_list_adjacencies))

				for _, v := range list_bindings[count_bindings] {
					all_visited[v] = 1
				}

				have_not_visited = true
				count_bindings++
				break
			}
		}
	}

	return list_bindings
}

func find_one_strong_binding(v, amount_vertexes int, list_adjacencies, inverse_list_adjacencies map[int][]int) []int {
	var visited_A []int
	var visited_B []int

	for i := 0; i < amount_vertexes; i++ {
		visited_A = append(visited_A, 0)
		visited_B = append(visited_B, 0)
	}

	visited_A[v] = 1
	visited_B[v] = 1

	dfs(v, list_adjacencies, &visited_A)
	dfs(v, inverse_list_adjacencies, &visited_B)

	var returned_list []int
	for i := 0; i < amount_vertexes; i++ {
		if visited_A[i]*visited_B[i] != 0 {
			returned_list = append(returned_list, i)
		}
	}

	return returned_list
}

func dfs(v int, list_adjacencies map[int][]int, visited *[]int) {
	(*visited)[v] = 1

	for _, new_v := range list_adjacencies[v] {
		if (*visited)[new_v] == 0 {
			dfs(new_v, list_adjacencies, visited)
		}
	}
}

func main() {
	var amount_vertexes int
	var amount_edges int
	list_adjacencies := make(map[int][]int, 20)
	inverse_list_adjacencies := make(map[int][]int, 20)

	fmt.Scan(&amount_vertexes)
	fmt.Scan(&amount_edges)

	for i := 0; i < amount_edges; i++ {
		var v1, v2 int

		fmt.Scan(&v1)
		fmt.Scan(&v2)

		list_adjacencies[v1] = append(list_adjacencies[v1], v2)
		inverse_list_adjacencies[v2] = append(inverse_list_adjacencies[v2], v1)
	}

	list_bindings := find_all_strong_bindings(amount_vertexes, list_adjacencies, inverse_list_adjacencies)
	amount_cond_orgraph_vertexes := len(list_bindings)
	list_cond_orgraph_adjacencies := make_cond_orgraph(amount_vertexes, list_bindings, list_adjacencies)
	/*
		fmt.Println(amount_cond_orgraph_vertexes)
		fmt.Println(list_bindings)
		fmt.Println(list_cond_orgraph_adjacencies)
	*/
	for _, elem := range get_base(amount_cond_orgraph_vertexes, list_cond_orgraph_adjacencies, list_bindings) {
		fmt.Printf("%d ", elem)
	}
}