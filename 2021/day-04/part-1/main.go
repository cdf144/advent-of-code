package main

import (
	"cdf144/aoc2021/utils"
	"fmt"
	"path/filepath"
	"strings"
)

type bingoField struct {
	num    int
	marked bool
}

func checkBoard(board [][]*bingoField, draws []int, minDraws, winScore *int) {
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
			if i+1 < *minDraws {
				*minDraws = i + 1
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
	minDraws := len(draws)
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

		checkBoard(board, draws, &minDraws, &winScore)
	}

	fmt.Printf("Minimum draws: %d\n", minDraws)
	fmt.Printf("Last draw: %d\n", draws[minDraws-1])
	printAnswer(winScore)
}
