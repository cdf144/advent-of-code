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
	visited := make(map[string]struct{})
	visited["start"] = struct{}{}
	backtrack("start", graph, visited, &paths)

	printAnswer(paths)
}

func backtrack(node string, graph map[string][]string, visited map[string]struct{}, paths *int) {
	if node == "end" {
		*paths++
		return
	}
	for _, neighbor := range graph[node] {
		if _, ok := visited[neighbor]; !ok || strings.ToUpper(neighbor) == neighbor {
			visited[neighbor] = struct{}{}
			backtrack(neighbor, graph, visited, paths)
			delete(visited, neighbor)
		}
	}
}
