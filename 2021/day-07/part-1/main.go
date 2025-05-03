package main

import (
	"path/filepath"
	"slices"
	"strings"

	"cdf144/aoc2021/utils"
)

func main() {
	input, printAnswer := utils.InitRaw(filepath.Join("..", "input.example.txt"))

	var crabs []int
	{
		inputSplit := strings.Split(strings.TrimRight(input, "\r\n"), ",")
		crabs = make([]int, len(inputSplit))
		for i, s := range inputSplit {
			crabs[i] = utils.Atoi(s)
		}
	}

	// The median minimizes the sum of absolute deviations from a set of points.
	slices.Sort(crabs)
	median := crabs[len(crabs)/2]

	minFuel := 0
	for _, crab := range crabs {
		minFuel += utils.Abs(crab - median)
	}

	printAnswer(minFuel)
}
