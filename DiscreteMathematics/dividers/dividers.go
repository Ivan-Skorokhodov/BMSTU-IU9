package main

import (
	"fmt"
	"os"
)

func find_del(num int) []int {
	var res []int
	i := 2
	for i*i <= num {
		if num%i == 0 {

			if i*i != num {
				res = append(res, i)
				res = append(res, num/i)
			} else {
				res = append(res, i)
			}
		}
		i++
	}
	return res
}

func find_all_table_del_elem(table_del_elem map[int][]int, list_keys *[]int) map[int][]int {
	len1 := len(*list_keys)

	for _, key := range *list_keys {
		for _, del := range table_del_elem[key] {

			_, status := table_del_elem[del]

			if status == false {
				*list_keys = append(*list_keys, del)
				table_del_elem[del] = find_del(del)
			}
		}
	}

	if len1 == len(*list_keys) {
		return table_del_elem
	}
	return find_all_table_del_elem(table_del_elem, list_keys)
}

func concat_stop_lists(stop_list *[]int, num int, stop_table, table_del_elem map[int][]int) {
	for _, j := range stop_table[num] {
		*stop_list = append(*stop_list, j)
	}

	for _, j := range table_del_elem[num] {
		*stop_list = append(*stop_list, j)
	}
}

func main() {
	var num int
	var list_keys []int

	table_del_elem := make(map[int][]int, 50)
	stop_table := make(map[int][]int, 50)

	fmt.Scan(&num)

	table_del_elem[num] = find_del(num)

	if num == 1 {
		fmt.Printf("graph {\n")
		fmt.Printf("\t%d\n", num)
		fmt.Println("}")
		os.Exit(0)
	}

	if len(table_del_elem[num]) == 0 {
		fmt.Printf("graph {\n")
		fmt.Printf("\t%d\n", num)
		fmt.Printf("\t%d\n", 1)
		fmt.Printf("\t%d--1\n", num)
		fmt.Println("}")
		os.Exit(0)
	}

	list_keys = append(list_keys, num)

	table_del_elem = find_all_table_del_elem(table_del_elem, &list_keys)

	fmt.Printf("graph {\n")
	for _, key := range list_keys {
		fmt.Printf("\t%d\n", key)
	}

	for i := 0; i <= len(table_del_elem[num]); i++ {
		for _, key := range list_keys {

			if len(table_del_elem[key]) == i {

				if i == 0 {

					stop_table[key] = append(stop_table[key], 1)
					fmt.Printf("\t%d--%d\n", key, 1)

				} else {

					var stop_list []int
					for _, del := range table_del_elem[key] {
						concat_stop_lists(&stop_list, del, stop_table, table_del_elem)
					}

					litle_stop_table := make(map[int]int, 10)
					for _, elem := range stop_list {
						litle_stop_table[elem] = 1
					}

					for _, elem := range table_del_elem[key] {
						_, status := litle_stop_table[elem]
						if status == false {
							fmt.Printf("\t%d--%d\n", key, elem)
						}
					}
				}
			}
		}
	}
	fmt.Println("}")
}