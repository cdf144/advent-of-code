package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"cdf144/aoc2021/utils"
)

type bingoField struct {
	num    int
	marked bool
}

func checkBoard(board [][]*bingoField, draws []int, maxDraws, winScore *int) {
	for i, draw := range draws {
	markLoop:
		for _, row := range board {
			for _, field := range row {
				if field.num == draw {
					field.marked = true
					break markLoop
				}
			}
		}

		if isBingoBoard(board) {
			if i+1 > *maxDraws {
				*maxDraws = i + 1
				*winScore = calculateScore(board, draw)
			}
			break
		}
	}
}

func isBingoBoard(board [][]*bingoField) bool {
	for _, row := range board {
		isBingo := true
		for _, field := range row {
			if !field.marked {
				isBingo = false
				break
			}
		}
		if isBingo {
			return true
		}
	}

	for j := range board[0] {
		isBingo := true
		for i := range board {
			if !board[i][j].marked {
				isBingo = false
				break
			}
		}
		if isBingo {
			return true
		}
	}

	return false
}

func calculateScore(board [][]*bingoField, draw int) int {
	sum := 0
	for _, row := range board {
		for _, field := range row {
			if !field.marked {
				sum += field.num
			}
		}
	}
	return sum * draw
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	drawsRaw := strings.Split(lines[0], ",")
	draws := make([]int, len(drawsRaw))
	for i, draw := range drawsRaw {
		draws[i] = utils.Atoi(draw)
	}

	boardSize := len(strings.Fields(lines[2]))
	maxDraws := 0
	winScore := 0

	for i := 2; i < len(lines); i += boardSize + 1 {
		board := make([][]*bingoField, boardSize)

		for j := range board {
			row := strings.Fields(lines[i+j])
			board[j] = make([]*bingoField, boardSize)
			for k, num := range row {
				board[j][k] = &bingoField{num: utils.Atoi(num), marked: false}
			}
		}

		checkBoard(board, draws, &maxDraws, &winScore)
	}

	fmt.Printf("Maximum draws: %d\n", maxDraws)
	fmt.Printf("Last draw: %d\n", draws[maxDraws-1])
	printAnswer(winScore)
}
