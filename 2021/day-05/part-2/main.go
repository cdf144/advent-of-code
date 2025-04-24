package main

import (
	"fmt"
	"path/filepath"

	"cdf144/aoc2021/utils"
)

type Vent struct {
	startX, startY int
	endX, endY     int
}

func (v Vent) isHorizontal() bool {
	return v.startY == v.endY
}

func (v Vent) isVertical() bool {
	return v.startX == v.endX
}

func findBorder(vents []Vent) (int, int) {
	maxX, maxY := 0, 0
	for _, vent := range vents {
		if vent.startX > maxX {
			maxX = vent.startX
		}
		if vent.endX > maxX {
			maxX = vent.endX
		}
		if vent.startY > maxY {
			maxY = vent.startY
		}
		if vent.endY > maxY {
			maxY = vent.endY
		}
	}
	return maxX, maxY
}

func drawVent(v Vent, diagram [][]int) {
	switch {
	case v.isHorizontal():
		for x := min(v.startX, v.endX); x <= max(v.startX, v.endX); x++ {
			diagram[v.startY][x]++
		}
	case v.isVertical():
		for y := min(v.startY, v.endY); y <= max(v.startY, v.endY); y++ {
			diagram[y][v.startX]++
		}
	default: // diagonal
		dx, dy := 1, 1
		if v.startX > v.endX {
			dx = -1
		}
		if v.startY > v.endY {
			dy = -1
		}

		x, y := v.startX, v.startY
		for {
			diagram[y][x]++
			if x == v.endX && y == v.endY {
				break
			}
			x += dx
			y += dy
		}
	}
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	vents := make([]Vent, len(lines))

	for i, line := range lines {
		var startX, startY, endX, endY int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &startX, &startY, &endX, &endY)
		vents[i] = Vent{startX, startY, endX, endY}
	}

	maxX, maxY := findBorder(vents)

	diagram := make([][]int, maxY+1)
	for i := range diagram {
		diagram[i] = make([]int, maxX+1)
	}

	for _, vent := range vents {
		drawVent(vent, diagram)
	}

	var overlaps int
	for _, row := range diagram {
		for _, cell := range row {
			if cell >= 2 {
				overlaps++
			}
		}
	}

	printAnswer(overlaps)
}
