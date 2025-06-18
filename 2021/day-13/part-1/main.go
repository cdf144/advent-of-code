package main

import (
	"path/filepath"
	"strings"

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
			fold := strings.Split(line, " ")[2]
			foldParts := strings.Split(fold, "=")
			line := utils.Atoi(foldParts[1])
			folds = append(folds, Fold{Axis: fold[0], Line: line})
		} else {
			parts := strings.Split(line, ",")
			point := Point{
				X: utils.Atoi(parts[0]),
				Y: utils.Atoi(parts[1]),
			}
			points[point] = struct{}{}
		}
	}

	fold := folds[0]
	newPoints := make(map[Point]struct{})
	for point := range points {
		if fold.Axis == 'x' && point.X > fold.Line {
			point.X = fold.Line - (point.X - fold.Line)
		} else if fold.Axis == 'y' && point.Y > fold.Line {
			point.Y = fold.Line - (point.Y - fold.Line)
		}
		newPoints[point] = struct{}{}
	}

	printAnswer(len(newPoints))
}
