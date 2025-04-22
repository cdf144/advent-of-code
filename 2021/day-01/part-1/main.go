package main

import (
	"cdf144/aoc2021/utils"
	"path/filepath"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	var increases int
	for i := 1; i < len(lines); i++ {
		if utils.Atoi(lines[i]) > utils.Atoi(lines[i-1]) {
			increases++
		}
	}

	printAnswer(increases)
}
