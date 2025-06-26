package main

import (
	"fmt"
	"math"
	"path/filepath"

	"cdf144/aoc2021/utils"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	template := lines[0]
	pairCounts := make(map[string]int)
	for i := range len(template) - 1 {
		pair := template[i : i+2]
		pairCounts[pair]++
	}

	rules := make(map[string]string)
	for _, line := range lines[2:] {
		var pair, insert string
		if _, err := fmt.Sscanf(line, "%s -> %s", &pair, &insert); err != nil {
			panic(err)
		}
		rules[pair] = insert
	}

	for range 40 {
		newPairCounts := make(map[string]int)

		for pair, count := range pairCounts {
			if insert, ok := rules[pair]; ok {
				newPairCounts[string(pair[0])+insert] += count
				newPairCounts[insert+string(pair[1])] += count
			} else {
				newPairCounts[pair] += count
			}
		}

		pairCounts = newPairCounts
	}

	counts := make(map[byte]int)
	for pair, count := range pairCounts {
		counts[pair[0]] += count
		counts[pair[1]] += count
	}
	// Adjust for the first and last characters which don't appear in 2 pairs
	counts[template[0]]++
	counts[template[len(template)-1]]++

	min, max := math.MaxInt, 0
	for _, count := range counts {
		count /= 2 // Each character is counted twice (appears in two pairs)
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	printAnswer(max - min)
}
