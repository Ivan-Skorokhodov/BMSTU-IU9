package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IsOperator(symbol string) bool {
	switch symbol {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	}
	return false
}

func IsDigit(symbol string) bool {
	switch symbol {
	case "+":
		return false
	case "-":
		return false
	case "*":
		return false
	case "(":
		return false
	case ")":
		return false
	}
	return true
}

func Convert(expression_list []string) []string {
	var new_expression_list []string
	i := 0

	for i < len(expression_list) {

		if i+4 < len(expression_list) &&
			expression_list[i] == "(" &&
			IsOperator(expression_list[i+1]) &&
			IsDigit(expression_list[i+2]) &&
			IsDigit(expression_list[i+3]) &&
			expression_list[i+4] == ")" {

			new_expression_list = append(new_expression_list, strconv.Itoa(Solver_elem(expression_list, i+1)))
			i += 5

		} else {
			new_expression_list = append(new_expression_list, expression_list[i])
			i++
		}
	}

	if len(new_expression_list) != 1 {
		return Convert(new_expression_list)
	} else {
		return new_expression_list
	}
}

func Solver_elem(expression_list []string, index int) int {
	symbol := expression_list[index]

	if IsDigit(symbol) {
		i, _ := strconv.Atoi(symbol)
		return i
	}

	switch symbol {
	case "+":
		return Solver_elem(expression_list, index+1) + Solver_elem(expression_list, index+2)
	case "-":
		return Solver_elem(expression_list, index+1) - Solver_elem(expression_list, index+2)
	case "*":
		return Solver_elem(expression_list, index+1) * Solver_elem(expression_list, index+2)
	}

	return 0
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

	res := Convert(expression_list)
	fmt.Println(res[0])
}