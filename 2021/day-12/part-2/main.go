package main

import (
	"path/filepath"
	"strings"

	"cdf144/aoc2021/utils"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	graph := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		u, v := parts[0], parts[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var paths int
	backtrack("start", false, graph, make(map[string]int), &paths)

	printAnswer(paths)
}

func backtrack(
	node string,
	hasVisitedSmallCaveTwice bool,
	graph map[string][]string,
	visitedCount map[string]int,
	paths *int,
) {
	if node == "end" {
		*paths++
		return
	}

	// node cannot be "end" per the above check
	// If node is "start", it will only be visited once thanks to the check in the below loop
	if !isBigCave(node) {
		visitedCount[node]++
	}

	for _, neighbor := range graph[node] {
		if neighbor == "start" {
			continue
		}

		neighborVisitCount := visitedCount[neighbor]

		if isBigCave(neighbor) || neighborVisitCount == 0 {
			backtrack(neighbor, hasVisitedSmallCaveTwice, graph, visitedCount, paths)
		} else if neighborVisitCount == 1 && !hasVisitedSmallCaveTwice {
			backtrack(neighbor, true, graph, visitedCount, paths)
		}
	}

	if !isBigCave(node) {
		visitedCount[node]--
	}
}

func isBigCave(name string) bool {
	return strings.ToUpper(name) == name
}
