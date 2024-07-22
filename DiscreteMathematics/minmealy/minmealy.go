package main

import "fmt"

type Mealy struct {
	count_states, alhabet_size, start_state int
	transition_matrix                       [][]int
	output_matrix                           [][]string
}

func Find(num int, table_parents *[]int) int {
	if (*table_parents)[num] == num {
		return num
	} else {
		return Find((*table_parents)[num], table_parents)
	}
}

func Union(num1, num2 int, table_parents, table_depth *[]int) {
	root_num1 := Find(num1, table_parents)
	root_num2 := Find(num2, table_parents)

	if (*table_depth)[root_num1] < (*table_depth)[root_num2] {
		(*table_parents)[root_num1] = root_num2
	} else {
		(*table_parents)[root_num2] = root_num1

		if (*table_depth)[root_num1] == (*table_depth)[root_num2] && root_num1 != root_num2 {
			(*table_depth)[root_num1]++
		}
	}
}

func Split1(mealy Mealy) (int, []int) {
	count_states := mealy.count_states
	var states_list, table_parents, table_depth []int

	for i := 0; i < mealy.count_states; i++ {
		table_parents = append(table_parents, i)
		table_depth = append(table_depth, 0)
	}

	for i := 0; i < mealy.count_states; i++ {
		for j := i + 1; j < mealy.count_states; j++ {

			var eq bool

			if Find(i, &table_parents) != Find(j, &table_parents) {
				eq = true
			}

			for x := 0; x < mealy.alhabet_size; x++ {
				if mealy.output_matrix[i][x] != mealy.output_matrix[j][x] {
					eq = false
					break
				}
			}

			if eq {
				Union(i, j, &table_parents, &table_depth)
				count_states--
			}
		}
	}

	for i := 0; i < mealy.count_states; i++ {
		states_list = append(states_list, Find(i, &table_parents))
	}

	return count_states, states_list
}

func Split(mealy Mealy, states_list []int) (int, []int) {
	var table_parents, table_depth []int
	transition_matrix := mealy.transition_matrix
	count_states := mealy.count_states

	for i := 0; i < mealy.count_states; i++ {
		table_parents = append(table_parents, i)
		table_depth = append(table_depth, 0)
	}

	for i := 0; i < mealy.count_states; i++ {
		for j := i + 1; j < mealy.count_states; j++ {

			if states_list[i] == states_list[j] && Find(i, &table_parents) != Find(j, &table_parents) {
				eq := true

				for x := 0; x < mealy.alhabet_size; x++ {
					w1 := transition_matrix[i][x]
					w2 := transition_matrix[j][x]
					if states_list[w1] != states_list[w2] {
						eq = false
						break
					}
				}

				if eq {
					Union(i, j, &table_parents, &table_depth)
					count_states--
				}
			}
		}
	}

	for i := 0; i < mealy.count_states; i++ {
		states_list[i] = Find(i, &table_parents)
	}

	return count_states, states_list
}

func AufenkampHohn(mealy *Mealy) *Mealy {
	m1, states_list := Split1(*mealy)

	var m2 int
	for {
		m2, states_list = Split(*mealy, states_list)
		if m1 == m2 {
			break
		}
		m1 = m2
	}

	helper1, helper2 := make([]int, mealy.count_states), make([]int, mealy.count_states)
	counter := 0
	for i := 0; i < mealy.count_states; i++ {
		if states_list[i] == i {
			helper1[counter] = i
			helper2[i] = counter
			counter++
		}
	}

	minimized := &Mealy{m1, mealy.alhabet_size, helper2[states_list[mealy.start_state]],
		make([][]int, m1), make([][]string, m1)}

	for i := 0; i < m1; i++ {
		minimized.transition_matrix[i] = make([]int, mealy.alhabet_size)
		minimized.output_matrix[i] = make([]string, mealy.alhabet_size)
	}

	for i := 0; i < minimized.count_states; i++ {
		for j := 0; j < mealy.alhabet_size; j++ {
			minimized.transition_matrix[i][j] = helper2[states_list[mealy.transition_matrix[helper1[i]][j]]]
			minimized.output_matrix[i][j] = mealy.output_matrix[helper1[i]][j]
		}
	}

	return minimized
}

func Canonic(mealy Mealy) *Mealy {
	var visited []int
	for i := 0; i < mealy.count_states; i++ {
		visited = append(visited, -1)
	}

	count := 0
	Dfs(mealy.start_state, &count, &visited, mealy.transition_matrix)

	res_transition_matrix := make([][]int, mealy.count_states)
	res_output_matrix := make([][]string, mealy.count_states)

	for old, new := range visited {
		res_transition_matrix[new] = mealy.transition_matrix[old]
		res_output_matrix[new] = mealy.output_matrix[old]
	}

	for _, array := range res_transition_matrix {
		for j := 0; j < mealy.alhabet_size; j++ {
			array[j] = visited[array[j]]
		}
	}

	return &Mealy{mealy.count_states, mealy.alhabet_size, 0, res_transition_matrix, res_output_matrix}
}

func Dfs(v int, count *int, visited *[]int, transition_matrix [][]int) {
	(*visited)[v] = *count // 0, 1, 2... not visited -1

	for _, new_v := range transition_matrix[v] {
		if (*visited)[new_v] == -1 {
			*count++
			Dfs(new_v, count, visited, transition_matrix)
		}
	}
}

func PrintGraph(mealy Mealy) {
	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")

	for i := 0; i < mealy.count_states; i++ {
		for j := 0; j < mealy.alhabet_size; j++ {
			signal := string(rune(97 + j))
			fmt.Printf("\t%d -> %d [label = \"%s(%s)\"]\n", i,
				mealy.transition_matrix[i][j], signal, mealy.output_matrix[i][j])
		}
	}

	fmt.Println("}")
}

func main() {
	var count_states, alhabet_size, start_state, num int
	var transition_matrix [][]int
	var output_matrix [][]string
	var s string

	fmt.Scan(&count_states)
	fmt.Scan(&alhabet_size)
	fmt.Scan(&start_state)

	for i := 0; i < count_states; i++ {
		var array []int
		for j := 0; j < alhabet_size; j++ {
			fmt.Scan(&num)
			array = append(array, num)
		}
		transition_matrix = append(transition_matrix, array)
	}

	for i := 0; i < count_states; i++ {
		var array []string
		for j := 0; j < alhabet_size; j++ {
			fmt.Scan(&s)
			array = append(array, s)
		}
		output_matrix = append(output_matrix, array)
	}

	mealy := Mealy{count_states, alhabet_size, start_state, transition_matrix, output_matrix}

	new := AufenkampHohn(&mealy)
	new_canonic := Canonic(*new)
	PrintGraph(*new_canonic)
}