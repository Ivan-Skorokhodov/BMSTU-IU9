package main

import "fmt"

type Component struct {
	amount_vertexes int
	list_vertexes   []int
	amount_edges    int
	minimal_vertex  int
}

func dfs(v, color int, list_adjacencies map[int][]int, visited *[]int, my_component *Component) {
	(*visited)[v] = color
	(*my_component).amount_vertexes++
	(*my_component).list_vertexes = append(my_component.list_vertexes, v)

	for _, new_v := range list_adjacencies[v] {
		if (*visited)[new_v] == 0 {
			dfs(new_v, color, list_adjacencies, visited, my_component)
		}
	}
}

func find_min_v(my_component Component) int {
	min_v := my_component.list_vertexes[0]
	for _, v := range my_component.list_vertexes {
		if v < min_v {
			min_v = v
		}
	}
	return min_v
}

func find_amount_edges(my_component Component, list_adjacencies map[int][]int) int {
	count := 0
	for _, v := range my_component.list_vertexes {
		count += len(list_adjacencies[v])
	}

	return count / 2
}

func find_max_component(list_components []Component) []int {
	max_amount_vertexes := list_components[0].amount_vertexes
	max_amount_edges := list_components[0].amount_edges
	index := 0

	for i, component := range list_components {

		if component.amount_vertexes > max_amount_vertexes {
			max_amount_vertexes = component.amount_vertexes
			max_amount_edges = component.amount_edges
			index = i

		} else if component.amount_vertexes == max_amount_vertexes {

			if component.amount_edges > max_amount_edges {
				max_amount_edges = component.amount_edges
				index = i
			}
		}
	}

	return list_components[index].list_vertexes
}

func main() {
	var amount_vertexes, amount_edges int
	var visited []int
	var list_components []Component
	var list_edges [][]int
	list_adjacencies := make(map[int][]int, 20)

	fmt.Scan(&amount_vertexes)
	fmt.Scan(&amount_edges)

	for i := 0; i < amount_vertexes; i++ {
		visited = append(visited, 0)
	}

	for i := 0; i < amount_edges; i++ {
		var v1, v2 int
		var mini_list []int

		fmt.Scan(&v1)
		fmt.Scan(&v2)

		mini_list = append(mini_list, v1)
		mini_list = append(mini_list, v2)
		list_edges = append(list_edges, mini_list)

		list_adjacencies[v1] = append(list_adjacencies[v1], v2)
		list_adjacencies[v2] = append(list_adjacencies[v2], v1)
	}

	color := 1
	have_zero := true
	for have_zero {
		start_v := -1

		for index, v := range visited {
			if v == 0 {
				start_v = index
				break
			}
		}

		if start_v == -1 {
			have_zero = false
		} else {
			var new_list []int
			my_component := Component{0, new_list, 0, 0}

			dfs(start_v, color, list_adjacencies, &visited, &my_component)
			list_components = append(list_components, my_component)
			color++
		}
	}

	for i := 0; i < len(list_components); i++ {
		list_components[i].minimal_vertex = find_min_v(list_components[i])
	}

	for i := 0; i < len(list_components); i++ {
		list_components[i].amount_edges = find_amount_edges(list_components[i], list_adjacencies)
	}

	table_vertexes := make(map[int]int, 30)
	for _, v := range find_max_component(list_components) {
		table_vertexes[v] = 1
	}

	fmt.Printf("graph {\n")
	for v, _ := range visited {
		_, status := table_vertexes[v]
		if status {
			fmt.Printf("\t%d [color = red]\n", v)
		} else {
			fmt.Printf("\t%d\n", v)
		}
	}

	for i := 0; i < len(list_edges); i++ {
		mini_list := list_edges[i]
		v1 := mini_list[0]
		v2 := mini_list[1]

		_, status := table_vertexes[v1]
		if status {
			fmt.Printf("\t%d -- %d [color = red]\n", v1, v2)
		} else {
			fmt.Printf("\t%d -- %d\n", v1, v2)
		}
	}

	fmt.Println("}")

	//fmt.Println(list_adjacencies)
	//fmt.Println(visited)
	//fmt.Println(list_components)
	//fmt.Println(table_vertexes)
}