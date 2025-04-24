package main

import (
	"path/filepath"

	"cdf144/aoc2021/utils"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	var increases int
	var window int
	prevWindow := -1

	for i, line := range lines {
		if i >= 3 {
			prevWindow = window
			window -= utils.Atoi(lines[i-3])
		}

		window += utils.Atoi(line)

		if prevWindow != -1 && window > prevWindow {
			increases++
		}

	}

	printAnswer(increases)
}
