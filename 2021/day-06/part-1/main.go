package main

import (
	"path/filepath"
	"strings"

	"cdf144/aoc2021/utils"
)

func main() {
	input, printAnswer := utils.InitRaw(filepath.Join("..", "input.example.txt"))
	inputSplit := strings.Split(strings.TrimRight(input, "\r\n"), ",")

	lanternfishes := make([]int, len(inputSplit))
	for i, s := range inputSplit {
		lanternfishes[i] = utils.Atoi(s)
	}

	for range 80 {
		updateLanternfishes(&lanternfishes)
	}

	printAnswer(len(lanternfishes))
}

func updateLanternfishes(lanternfishes *[]int) {
	for i, fish := range *lanternfishes {
		if fish == 0 {
			(*lanternfishes)[i] = 6
			*lanternfishes = append(*lanternfishes, 8)
		} else {
			(*lanternfishes)[i]--
		}
	}
}
