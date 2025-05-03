package main

import (
	"path/filepath"
	"strings"

	"cdf144/aoc2021/utils"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	var answer int

	for _, line := range lines {
		outputDigits := strings.Fields(
			strings.Split(line, " | ")[1],
		)
		for _, digit := range outputDigits {
			switch len(digit) {
			case 2, 3, 4, 7:
				answer++
			}
		}
	}

	printAnswer(answer)
}
