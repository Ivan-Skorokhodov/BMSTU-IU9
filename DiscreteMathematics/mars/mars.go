package main

import "fmt"

var list_of_combs [][]int

func get_stop_list(res []int, stop_list_candidats map[int][]int) map[int]bool {
	stop_list := make(map[int]bool, 5)

	for _, candidat := range res {
		for _, stop_candidat := range stop_list_candidats[candidat] {
			stop_list[stop_candidat] = true
		}
	}

	return stop_list
}

func get_table_candidats(list_candidats []int) map[int]bool {
	table := make(map[int]bool, 5)

	for _, candidat := range list_candidats {
		table[candidat] = true
	}

	return table
}

func get_another_group(n int, group_table map[int]bool) []int {
	var another_group []int

	for i := 0; i < n; i++ {
		if !group_table[i] {
			another_group = append(another_group, i)
		}
	}

	return another_group
}

func comb(res, input []int, n, global_n int, stop_list_candidats map[int][]int) {
	if n != 0 && len(res) <= global_n/2 {
		stop_list := get_stop_list(res, stop_list_candidats)

		for i, _ := range input {

			if !stop_list[input[i]] {
				var new_res []int
				new_res = append(new_res, res...)
				new_res = append(new_res, input[i])

				first_gropup_table := get_table_candidats(new_res)
				second_group := get_another_group(global_n, first_gropup_table)
				second_group_stop_list := get_stop_list(second_group, stop_list_candidats)

				check := true
				for _, candidat := range second_group {
					if second_group_stop_list[candidat] {
						check = false
						break
					}
				}

				if check {
					var array []int
					for _, i := range new_res {
						array = append(array, i+1)
					}
					list_of_combs = append(list_of_combs, array)
				}

				comb(new_res, input[i+1:], n-1, global_n, stop_list_candidats)
			}
		}
	}
}

func main() {
	var n int
	var sign string
	var input, res []int
	list_of_combs = make([][]int, 0)

	stop_list_candidats := make(map[int][]int, 20)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&sign)
			if sign == "+" {
				stop_list_candidats[i] = append(stop_list_candidats[i], j)
			}
		}
	}

	for i := 0; i < n; i++ {
		input = append(input, i)
	}

	comb(res, input, n, n, stop_list_candidats)

	length := n / 2
	find := false
	for length > 0 && !find {
		for _, array := range list_of_combs {
			if len(array) == length {

				for _, candidat := range array {
					fmt.Printf("%d ", candidat)
				}

				find = true
				break
			}
		}
		length--
	}

	if !find {
		fmt.Println("No solution")
	}
	/*
		for _, array := range list_of_combs {
			fmt.Println(array)
		}
	*/
}