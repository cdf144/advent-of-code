package main

import (
	"cdf144/aoc2021/utils"
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	actions := make([]string, len(lines))
	units := make([]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		actions[i] = parts[0]
		units[i] = utils.Atoi(parts[1])
	}

	var x, y, aim int

	for i, action := range actions {
		switch action {
		case "forward":
			x += units[i]
			y += aim * units[i]
		case "down":
			aim += units[i]
		case "up":
			aim -= units[i]
		}
	}

	fmt.Printf("Final position: x=%d, y=%d\n", x, y)
	printAnswer(x * y)
}
