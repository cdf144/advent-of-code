package main

import (
	"cdf144/aoc2021/utils"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	// Doesn't have to be a map since we know the input is binary, I'm just exploring Go
	freqs := make([]map[string]int, utf8.RuneCountInString(lines[0]))

	for i := range freqs {
		freqs[i] = make(map[string]int)
	}

	for _, line := range lines {
		for i, char := range line {
			freqs[i][string(char)]++
		}
	}

	gammaBin := strings.Builder{}
	epsilonBin := strings.Builder{}

	for _, freq := range freqs {
		var maxChar, minChar string
		if freq["0"] > freq["1"] {
			maxChar = "0"
			minChar = "1"
		} else {
			maxChar = "1"
			minChar = "0"
		}

		gammaBin.WriteString(maxChar)
		epsilonBin.WriteString(minChar)
	}

	gamma, err := strconv.ParseInt(gammaBin.String(), 2, 64)
	if err != nil {
		log.Fatalf("Error converting gamma to decimal: %v", err)
	}

	epsilon, err := strconv.ParseInt(epsilonBin.String(), 2, 64)
	if err != nil {
		log.Fatalf("Error converting epsilon to decimal: %v", err)
	}

	fmt.Printf("Gamma: %d\nEpsilon: %d\n", gamma, epsilon)
	printAnswer(gamma * epsilon)
}
