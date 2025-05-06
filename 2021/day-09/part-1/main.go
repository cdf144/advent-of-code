package main

import (
	"path/filepath"

	"cdf144/aoc2021/utils"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	heightmap := make([][]int, len(lines))
	for i, row := range lines {
		heightmap[i] = make([]int, len(row))
		for j, point := range row {
			heightmap[i][j] = int(point - '0')
		}
	}

	rows := len(heightmap)
	cols := len(heightmap[0])
	neighbors := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	isLowPoint := func(i, j int) bool {
		currLocation := heightmap[i][j]

		for _, delta := range neighbors {
			x, y := i+delta[0], j+delta[1]
			if x >= 0 && x < rows && y >= 0 && y < cols {
				if heightmap[x][y] <= currLocation {
					return false
				}
			}
		}

		return true
	}

	var totalRiskLevel int
	for i := range rows {
		for j := range cols {
			if isLowPoint(i, j) {
				totalRiskLevel += 1 + heightmap[i][j]
			}
		}
	}

	printAnswer(totalRiskLevel)
}
