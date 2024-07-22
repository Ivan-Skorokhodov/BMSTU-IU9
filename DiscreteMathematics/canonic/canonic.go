package main

import "fmt"

func dfs(v int, count *int, visited *[]int, transition_matrix [][]int) {
	(*visited)[v] = *count // 0, 1, 2... not visited -1

	for _, new_v := range transition_matrix[v] {
		if (*visited)[new_v] == -1 {
			*count++
			dfs(new_v, count, visited, transition_matrix)
		}
	}
}

func main() {
	var s string
	var n, m, q0, num int
	var transition_matrix [][]int
	var output_matrix [][]string

	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q0)

	for i := 0; i < n; i++ {
		var array []int
		for j := 0; j < m; j++ {
			fmt.Scan(&num)
			array = append(array, num)
		}
		transition_matrix = append(transition_matrix, array)
	}

	for i := 0; i < n; i++ {
		var array []string
		for j := 0; j < m; j++ {
			fmt.Scan(&s)
			array = append(array, s)
		}
		output_matrix = append(output_matrix, array)
	}

	var visited []int
	for i := 0; i < n; i++ {
		visited = append(visited, -1)
	}

	count := 0
	dfs(q0, &count, &visited, transition_matrix)

	res_transition_matrix := make([][]int, n)
	res_output_matrix := make([][]string, n)

	for old, new := range visited {
		res_transition_matrix[new] = transition_matrix[old]
		res_output_matrix[new] = output_matrix[old]
	}

	for _, array := range res_transition_matrix {
		for j := 0; j < m; j++ {
			array[j] = visited[array[j]]
		}
	}

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println("0")

	for _, array := range res_transition_matrix {
		for _, i := range array {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}

	for _, array := range res_output_matrix {
		for _, s := range array {
			fmt.Printf("%s ", s)
		}
		fmt.Println()
	}
}