package main

import (
	"fmt"
	"path/filepath"

	"cdf144/aoc2021/utils"
)

type Point struct {
	X, Y int
}

type Fold struct {
	Axis byte
	Line int
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.txt"))

	points := make(map[Point]struct{})
	var folds []Fold
	for _, line := range lines {
		if line == "" {
			continue
		}
		if line[0] == 'f' {
			var foldAxis byte
			var foldLine int
			fmt.Sscanf(line, "fold along %c=%d", &foldAxis, &foldLine)
			folds = append(folds, Fold{Axis: foldAxis, Line: foldLine})
		} else {
			var x, y int
			fmt.Sscanf(line, "%d,%d", &x, &y)
			point := Point{X: x, Y: y}
			points[point] = struct{}{}
		}
	}

	for _, fold := range folds {
		newPoints := make(map[Point]struct{})
		for point := range points {
			if fold.Axis == 'x' && point.X > fold.Line {
				point.X = fold.Line - (point.X - fold.Line)
			} else if fold.Axis == 'y' && point.Y > fold.Line {
				point.Y = fold.Line - (point.Y - fold.Line)
			}
			newPoints[point] = struct{}{}
		}
		points = newPoints
	}

	printAnswer("")
	printPaper(points)
}

func printPaper(points map[Point]struct{}) {
	maxX, maxY := 0, 0
	for point := range points {
		maxX = max(maxX, point.X)
		maxY = max(maxY, point.Y)
	}

	for y := range maxY + 1 {
		for x := range maxX + 1 {
			if _, ok := points[Point{X: x, Y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
