package main

import (
	"fmt"
	"strings"
)

type Rule struct {
	left  string
	right string
}

func dfs(node string, graph map[string][]string, visited map[string]bool, stack *[]string) {
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited, stack)
		}
	}
	*stack = append(*stack, node)
}

func getTranspose(graph map[string][]string) map[string][]string {
	transpose := make(map[string][]string)
	for node := range graph {
		for _, neighbor := range graph[node] {
			transpose[neighbor] = append(transpose[neighbor], node)
		}
	}
	return transpose
}

func stronglyConnectedComponents(graph map[string][]string) [][]string {
	stack := []string{}
	visited := make(map[string]bool)

	for node := range graph {
		if !visited[node] {
			dfs(node, graph, visited, &stack)
		}
	}

	transpose := getTranspose(graph)
	visited = make(map[string]bool)
	var components [][]string

	for i := len(stack) - 1; i >= 0; i-- {
		node := stack[i]
		if !visited[node] {
			componentStack := []string{}
			dfs(node, transpose, visited, &componentStack)
			components = append(components, componentStack)
		}
	}

	return components
}

func isComponentStronglyRegular(component []string, rules []Rule) bool {
	isLeftRecursive := false
	isRightRecursive := false

	for _, rule := range rules {
		for _, nonterminal := range component {

			if rule.left == nonterminal {
				right := rule.right
				for _, nt := range component {
					if len(right) != 1 {
						check1 := false
						check2 := false

						if strings.HasPrefix(right, nt) {
							check1 = true
							isLeftRecursive = true

						}
						if strings.HasSuffix(right, nt) {
							check2 = true
							isRightRecursive = true

						} 
						if check2 == false && check1 == false {
							for _, s := range right {
								if string(s) == nt {
									return false
								}
							}
						}
					}
				}
			}

		}	
	}

	return !(isLeftRecursive && isRightRecursive)
}

func main() {
	var N int
	var startSymbol string
	fmt.Scan(&N)
	fmt.Scan(&startSymbol)

	rules := make([]Rule, N)
	nonterminals := make(map[string]bool)

	for i := 0; i < N; i++ {
		var part string
		fmt.Scan(&part)
		parts := strings.Split(part, "->")
		
		left := parts[0]
		right := parts[1]
		rules[i] = Rule{left, right}
		nonterminals[left] = true
	}

	graph := make(map[string][]string)

	for _, rule := range rules {
		for _, char := range rule.right {
			if char >= 'A' && char <= 'Z' {
				graph[rule.left] = append(graph[rule.left], string(char))
			}
		}
	}

	components := stronglyConnectedComponents(graph)

	for _, component := range components {
		if !isComponentStronglyRegular(component, rules) {
			fmt.Println("no")
			/*
			fmt.Println(rules)
			fmt.Println(graph)
			fmt.Println(components)
			*/
			return
		}
	}

	fmt.Println("yes")
	/*
	fmt.Println(rules)
	fmt.Println(graph)
	fmt.Println(components)
	*/
}

/*
func main() {
	var n, a int

	fmt.Scan(&n)
	fmt.Scan(&a)

	list_adjacencies := make(map[string][]string, n)

	for i := 0; i < n + 1; i++ {
		var str string
		fmt.Scan(&str)

		var list_symbols []string
		for _, s := range str[3:] {
			fmt.Println(string(s))
			if s <= 'Z' && s >= 'A' {
				list_symbols = append(list_symbols, string(s))
			}
		}

		list_adjacencies[string(str[0])] = append(list_adjacencies[string(str[0])], list_symbols...)
	}

	fmt.Println(list_adjacencies)
}
*/
