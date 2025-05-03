package main

import (
	"path/filepath"
	"strings"

	"cdf144/aoc2021/utils"
)

func calculateFuel(crabs []int, position int) int {
	fuel := 0
	for _, crab := range crabs {
		distance := utils.Abs(crab - position)
		fuel += (distance * (distance + 1)) / 2
	}
	return fuel
}

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

	// The mean minimizes the sum of squared deviations from a set of points.
	// (n * (n + 1)) / 2 = (n^2 + n) / 2
	// To minimize sum of fuel, we either minimize n^2 (mean) or n (median).
	// Since n^2 has a larger impact on the sum, we use the mean.
	sum := 0
	for _, crab := range crabs {
		sum += crab
	}
	mean := float32(sum) / float32(len(crabs))

	printAnswer(
		min(
			calculateFuel(crabs, int(mean)),
			calculateFuel(crabs, int(mean+1)),
		),
	)
}
