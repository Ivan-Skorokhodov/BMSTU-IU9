package main

import (
	"fmt"
	"os"
)

func enqueue(queue *[]int, element int) {
	*queue = append(*queue, element)
}

func dequeue(queue *[]int) int {
	element := (*queue)[0]

	if len(*queue) == 1 {
		*queue = []int{}
		return element

	}

	*queue = (*queue)[1:]
	return element
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bfs(start int, distance *[]int, list_adjacencies map[int][]int) {
	(*distance)[start] = 0

	var queue []int
	enqueue(&queue, start)

	for len(queue) > 0 {
		v := dequeue(&queue)

		for _, new_v := range list_adjacencies[v] {
			if (*distance)[new_v] == 1000000 {
				(*distance)[new_v] = (*distance)[v] + 1
				enqueue(&queue, new_v)
			}
		}
	}

}

func main() {
	var amount_vertexes, amount_edges, amount_main_vertexes int
	var list_main_vertexes, list_ans []int
	var list_distances [][]int
	list_adjacencies := make(map[int][]int, 20)

	fmt.Scan(&amount_vertexes)
	fmt.Scan(&amount_edges)

	for i := 0; i < amount_edges; i++ {
		var v1, v2 int

		fmt.Scan(&v1)
		fmt.Scan(&v2)

		list_adjacencies[v1] = append(list_adjacencies[v1], v2)
		list_adjacencies[v2] = append(list_adjacencies[v2], v1)
	}

	fmt.Scan(&amount_main_vertexes)
	for i := 0; i < amount_main_vertexes; i++ {
		var v int
		fmt.Scan(&v)
		list_main_vertexes = append(list_main_vertexes, v)
	}

	for _, start := range list_main_vertexes {
		var distance []int
		for i := 0; i < amount_vertexes; i++ {
			distance = append(distance, 1000000)
		}

		bfs(start, &distance, list_adjacencies)
		list_distances = append(list_distances, distance)
	}

	for i := 0; i < amount_vertexes; i++ {
		dist := list_distances[0][i]

		if dist == 1000000 {
			continue
		}

		all_the_same := true

		for j := 0; j < len(list_distances); j++ {
			if list_distances[j][i] != dist {
				all_the_same = false
			}
		}

		if all_the_same {
			list_ans = append(list_ans, i)
		}
	}

	if len(list_ans) == 0 {
		fmt.Println("-")
		os.Exit(0)
	}

	for _, v := range list_ans {
		fmt.Printf("%d ", v)
	}
}