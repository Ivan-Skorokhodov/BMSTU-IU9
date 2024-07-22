package main

import (
	"fmt"
	"strings"
)

type Rule struct {
  left  string
  right string
}

func main() {
  var N int
  var startSymbol string
  fmt.Scanf("%d %s", &N, &startSymbol)

  rules := make([]Rule, N)
  nonterminals := make(map[string]bool)

  for i := 0; i < N; i++ {
    var part string
    fmt.Scanf("%s", &part)
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
      return
    }
  }
  fmt.Println("yes")
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
          if strings.HasPrefix(right, nt) {
            isLeftRecursive = true
          }
          if strings.HasSuffix(right, nt) {
            isRightRecursive = true
          }
        }
      }
    }
  }

  return !(isLeftRecursive && isRightRecursive)
}