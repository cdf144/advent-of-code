package main

import (
	"path/filepath"

	"cdf144/aoc2021/utils"
)

var errorToScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var chunkOpenToClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	var totalScore int

	for _, line := range lines {
		var stack []rune
		for _, char := range line {
			if _, isOpen := chunkOpenToClose[char]; isOpen {
				stack = append(stack, char)
			} else {
				stackLastIdx := len(stack) - 1
				if stackLastIdx < 0 || char != chunkOpenToClose[stack[stackLastIdx]] {
					totalScore += errorToScore[char]
					break
				}
				stack = stack[:stackLastIdx]
			}
		}
	}

	printAnswer(totalScore)
}
