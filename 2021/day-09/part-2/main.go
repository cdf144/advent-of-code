package main

import (
	"path/filepath"
	"slices"

	"cdf144/aoc2021/utils"
)

var neighbors = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

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

	marked := make([][]bool, rows)
	for i := range rows {
		marked[i] = make([]bool, cols)
	}

	resetMarked := func() {
		for i := range rows {
			for j := range cols {
				marked[i][j] = false
			}
		}
	}

	var basins []int
	for i := range rows {
		for j := range cols {
			if isLowPoint(i, j) {
				resetMarked()
				basins = append(basins, getBasinSize(i, j, heightmap, marked))
			}
		}
	}

	slices.SortFunc(basins, func(a, b int) int {
		return b - a
	})
	printAnswer(basins[0] * basins[1] * basins[2])
}

func getBasinSize(i, j int, heightmap [][]int, marked [][]bool) int {
	if marked[i][j] {
		return 0
	}
	marked[i][j] = true

	result := 1
	rows := len(heightmap)
	cols := len(heightmap[0])
	currLocation := (heightmap)[i][j]

	for _, delta := range neighbors {
		x, y := i+delta[0], j+delta[1]
		if x >= 0 && x < rows && y >= 0 && y < cols {
			location := heightmap[x][y]
			if location != 9 && location >= currLocation {
				result += getBasinSize(x, y, heightmap, marked)
			}
		}
	}

	return result
}
