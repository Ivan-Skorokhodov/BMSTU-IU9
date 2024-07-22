package main

import (
	"fmt"
)

type Elem struct {
	v, color int
}

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

func bfs(distance *[]int, list_adjacencies map[int][]Elem, list_pred *map[int][]Elem) {
	(*distance)[1] = 0

	var queue []int
	enqueue(&queue, 1)

	for len(queue) > 0 {
		v := dequeue(&queue)

		for _, new_v := range list_adjacencies[v] {
			color_edge := new_v.color

			if (*distance)[new_v.v] == 1000000 {
				(*distance)[new_v.v] = (*distance)[v] + 1
				(*list_pred)[new_v.v] = append((*list_pred)[new_v.v], Elem{v, color_edge})
				enqueue(&queue, new_v.v)

			} else if (*distance)[new_v.v] == (*distance)[v]+1 {
				(*list_pred)[new_v.v] = append((*list_pred)[new_v.v], Elem{v, color_edge})
			}
		}
	}
}

func is_elem_in_list(elem Elem, array []Elem) bool {
	for _, array_elem := range array {
		if array_elem == elem {
			return true
		}
	}
	return false
}

func find_table_ways(old_v int, list_pred map[int][]Elem, table_ways_adjacencies *map[int][]Elem) {
	_, status := list_pred[old_v]

	if status {
		for _, pred := range list_pred[old_v] {
			new_elem := Elem{old_v, pred.color}

			if !is_elem_in_list(new_elem, (*table_ways_adjacencies)[pred.v]) {
				(*table_ways_adjacencies)[pred.v] = append((*table_ways_adjacencies)[pred.v], new_elem)
				find_table_ways(pred.v, list_pred, table_ways_adjacencies)
			}
		}
	}
}

func find_ways(old_v int, table_ways_adjacencies map[int][]Elem, way []int, ways *[][]int) {
	_, status := table_ways_adjacencies[old_v]

	if status {
		for _, new_elem := range table_ways_adjacencies[old_v] {
			var new_way []int

			new_way = append(new_way, way...)
			new_way = append(new_way, new_elem.color)

			find_ways(new_elem.v, table_ways_adjacencies, new_way, ways)
		}
	} else {
		*ways = append(*ways, way)
	}
}

func less(l1, l2 []int) bool {
	for i := 0; i < len(l1); i++ {
		if l1[i] > l2[i] {
			return false

		} else if l1[i] < l2[i] {
			return true
		}
	}
	return false
}

func main() {
	var amount_vertexes, amount_edges, v1, v2, color int
	list_adjacencies := make(map[int][]Elem, 20)
	list_pred := make(map[int][]Elem, 20)
	table_ways_adjacencies := make(map[int][]Elem, 20)
	var way []int
	var ways [][]int

	fmt.Scan(&amount_vertexes)
	fmt.Scan(&amount_edges)

	for i := 0; i < amount_edges; i++ {
		fmt.Scan(&v1)
		fmt.Scan(&v2)
		fmt.Scan(&color)

		list_adjacencies[v1] = append(list_adjacencies[v1], Elem{v2, color})
		list_adjacencies[v2] = append(list_adjacencies[v2], Elem{v1, color})

	}

	var distance []int
	for i := 0; i < amount_vertexes+1; i++ {
		distance = append(distance, 1000000)
	}

	bfs(&distance, list_adjacencies, &list_pred)
	find_table_ways(amount_vertexes, list_pred, &table_ways_adjacencies)

	fmt.Println(distance[amount_vertexes])

	find_ways(1, table_ways_adjacencies, way, &ways)

	min_n := ways[0]
	for i := 1; i < len(ways); i++ {
		if less(ways[i], min_n) {
			min_n = ways[i]
		}
	}

	for _, n := range min_n {
		fmt.Printf("%d ", n)
	}
}