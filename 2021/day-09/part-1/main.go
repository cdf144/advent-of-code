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

	var totalRiskLevel int
	for i, row := range heightmap {
		for j := range row {
			if isLowPoint(i, j, heightmap) {
				totalRiskLevel += 1 + heightmap[i][j]
			}
		}
	}

	printAnswer(totalRiskLevel)
}

func isLowPoint(i, j int, heightmap [][]int) bool {
	rows := len(heightmap)
	cols := len(heightmap[0])
	pointHeight := heightmap[i][j]

	neighbors := [][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for _, delta := range neighbors {
		x, y := i+delta[0], j+delta[1]
		if x >= 0 && x < rows && y >= 0 && y < cols {
			if heightmap[x][y] <= pointHeight {
				return false
			}
		}
	}

	return true
}
