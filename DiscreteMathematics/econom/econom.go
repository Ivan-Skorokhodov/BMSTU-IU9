package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsOperator(symbol string) bool {
	switch symbol {
	case "#":
		return true
	case "$":
		return true
	case "@":
		return true
	}
	return false
}

func IsElem(symbol string) bool {
	switch symbol {
	case "#":
		return false
	case "$":
		return false
	case "@":
		return false
	case "(":
		return false
	case ")":
		return false
	}
	return true
}

func Convert(expression_list []string, table map[string]int, count int) int {
	var new_expression_list []string
	i := 0

	for i < len(expression_list) {

		if i+4 < len(expression_list) &&
			expression_list[i] == "(" &&
			IsOperator(expression_list[i+1]) &&
			IsElem(expression_list[i+2]) &&
			IsElem(expression_list[i+3]) &&
			expression_list[i+4] == ")" {

			new_elem := expression_list[i] +
				expression_list[i+1] +
				expression_list[i+2] +
				expression_list[i+3] +
				expression_list[i+4]

			_, status := table[new_elem]

			if status == false {
				table[new_elem] = 1
				count++
			}

			new_expression_list = append(new_expression_list, new_elem)
			i += 5

		} else {
			new_expression_list = append(new_expression_list, expression_list[i])
			i++
		}
	}

	if len(new_expression_list) != 1 {
		return Convert(new_expression_list, table, count)
	} else {
		return count
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression_line := scanner.Text()
	var expression_list []string

	for _, i := range expression_line {
		if i != 32 {
			expression_list = append(expression_list, string(i))
		}
	}

	table := make(map[string]int, 20)

	res := Convert(expression_list, table, 0)
	fmt.Println(res)
}