package main

import (
	"cdf144/aoc2021/utils"
	"math"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	input, printAnswer := utils.InitRaw(filepath.Join("..", "input.txt"))

	var crabs []int
	{
		inputSplit := strings.Split(strings.TrimRight(input, "\r\n"), ",")
		crabs = make([]int, len(inputSplit))
		for i, s := range inputSplit {
			crabs[i] = utils.Atoi(s)
		}
	}

	minPos, maxPos := slices.Min(crabs), slices.Max(crabs)
	minFuel := math.MaxInt

	for i := minPos; i <= maxPos; i++ {
		fuel := 0
		for _, crab := range crabs {
			distance := utils.Abs(crab - i)
			fuel += (distance * (distance + 1)) / 2
		}
		minFuel = min(minFuel, fuel)
	}

	printAnswer(minFuel)
}
