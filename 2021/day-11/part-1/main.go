package main

import (
	"cdf144/aoc2021/utils"
	"path/filepath"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	octopuses := make([][]int, len(lines))
	for i, line := range lines {
		octopusRow := make([]int, len(line))
		for j, char := range line {
			octopusRow[j] = int(char - '0')
		}
		octopuses[i] = octopusRow
	}

	var flashes int
	for range 100 {
		step(octopuses, &flashes)
	}

	printAnswer(flashes)
}

func step(octopuses [][]int, flashes *int) {
	for i := range octopuses {
		for j := range octopuses[i] {
			octopuses[i][j]++
		}
	}

	for i := range octopuses {
		for j := range octopuses[i] {
			if octopuses[i][j] > 9 {
				flash(octopuses, i, j, flashes)
			}
		}
	}
}

func flash(octopuses [][]int, i, j int, flashes *int) {
	(*flashes)++
	octopuses[i][j] = 0
	// PERF: Yes this is recalculated each recursive call but I ain't passing these as params
	m, n := len(octopuses), len(octopuses[0])

	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			// Surrounding octopus that is already flashed, or the current octopus itself
			if octopuses[x][y] == 0 {
				continue
			}
			octopuses[x][y]++
			if octopuses[x][y] > 9 {
				flash(octopuses, x, y, flashes)
			}
		}
	}
}
