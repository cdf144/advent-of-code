package main

import (
	"path/filepath"
	"slices"

	"cdf144/aoc2021/utils"
)

var closeToScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var chunkOpenToClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	var scores []int

lineLoop:
	for _, line := range lines {
		var stack []rune

		for _, char := range line {
			if _, isOpen := chunkOpenToClose[char]; isOpen {
				stack = append(stack, char)
			} else {
				stackLastIdx := len(stack) - 1
				if stackLastIdx < 0 || char != chunkOpenToClose[stack[stackLastIdx]] {
					continue lineLoop
				}
				stack = stack[:stackLastIdx]
			}
		}

		if len(stack) == 0 {
			continue
		}

		var score int
		for len(stack) > 0 {
			stackLastIdx := len(stack) - 1
			char := stack[stackLastIdx]
			stack = stack[:stackLastIdx]
			score = score*5 + closeToScore[chunkOpenToClose[char]]
		}

		scores = append(scores, score)
	}

	slices.Sort(scores)
	printAnswer(scores[len(scores)/2])
}
