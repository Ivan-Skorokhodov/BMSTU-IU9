package main

import "fmt"

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

	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			signal := string(rune(97 + j))
			fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n", i, transition_matrix[i][j], signal, output_matrix[i][j])
		}
	}

	fmt.Println("}")
}