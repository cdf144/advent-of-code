package main

import (
	"cdf144/aoc2021/utils"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
)

func calculateFrequency(slice []string, index int) map[byte]int {
	freq := make(map[byte]int)
	freq['0'] = 0
	freq['1'] = 0
	for _, s := range slice {
		freq[s[index]]++
	}
	return freq
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	oxygenFilter := make([]string, len(lines))
	copy(oxygenFilter, lines)

	co2Filter := make([]string, len(lines))
	copy(co2Filter, lines)

	for i := range len(lines[0]) {
		if len(oxygenFilter) == 1 {
			break
		}

		freq := calculateFrequency(oxygenFilter, i)

		mostCommon := byte('1')
		if freq['0'] > freq['1'] {
			mostCommon = byte('0')
		}

		oxygenFilter = utils.Filter(oxygenFilter, func(s string) bool {
			return s[i] == mostCommon
		})
	}

	for i := range len(lines[0]) {
		if len(co2Filter) == 1 {
			break
		}

		freq := calculateFrequency(co2Filter, i)

		leastCommon := byte('0')
		if freq['1'] < freq['0'] {
			leastCommon = byte('1')
		}

		co2Filter = utils.Filter(co2Filter, func(s string) bool {
			return s[i] == leastCommon
		})
	}

	oxygen, err := strconv.ParseInt(oxygenFilter[0], 2, 64)
	if err != nil {
		log.Fatalf("Error parsing oxygen value: %v", err)
	}

	co2, err := strconv.ParseInt(co2Filter[0], 2, 64)
	if err != nil {
		log.Fatalf("Error parsing CO2 value: %v", err)
	}

	fmt.Printf("Oxygen generator rating: %d\nCO2 scrubber rating: %d\n", oxygen, co2)
	printAnswer(oxygen * co2)
}
