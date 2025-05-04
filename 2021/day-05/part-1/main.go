package main

import (
	"fmt"
	"path/filepath"

	"cdf144/aoc2021/utils"
)

type vent struct {
	startX, startY int
	endX, endY     int
}

func (v vent) isHorizontal() bool {
	return v.startY == v.endY
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))
	// Skip diagonal lines
	lines = utils.Filter(lines, func(line string) bool {
		var startX, startY, endX, endY int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &startX, &startY, &endX, &endY)
		return startX == endX || startY == endY
	})

	vents := make([]vent, len(lines))

	for i, line := range lines {
		var startX, startY, endX, endY int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &startX, &startY, &endX, &endY)
		vents[i] = vent{startX, startY, endX, endY}
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

func findBorder(vents []vent) (int, int) {
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

func drawVent(v vent, diagram [][]int) {
	if v.isHorizontal() {
		for x := min(v.startX, v.endX); x <= max(v.startX, v.endX); x++ {
			diagram[v.startY][x]++
		}
	} else {
		for y := min(v.startY, v.endY); y <= max(v.startY, v.endY); y++ {
			diagram[y][v.startX]++
		}
	}
}
