package main

import (
	"path/filepath"
	"strings"

	"cdf144/aoc2021/utils"
)

func updateLanternfishes(lanternfishes map[int]uint64) {
	produced := lanternfishes[0]

	for i := range 8 {
		lanternfishes[i] = lanternfishes[i+1]
	}

	lanternfishes[6] += produced
	lanternfishes[8] = produced
}

func main() {
	input, printAnswer := utils.InitRaw(filepath.Join("..", "input.example.txt"))
	inputSplit := strings.Split(strings.TrimRight(input, "\r\n"), ",")

	// Instead of straight up simulating the list of lanternfishes, we keep
	// track of how many lanternfishes are at each timer value.
	lanternfishes := make(map[int]uint64, 9)
	for _, s := range inputSplit {
		fish := utils.Atoi(s)
		lanternfishes[fish]++
	}

	for range 256 {
		updateLanternfishes(lanternfishes)
	}

	answer := uint64(0)
	for _, v := range lanternfishes {
		answer += v
	}

	printAnswer(answer)
}
