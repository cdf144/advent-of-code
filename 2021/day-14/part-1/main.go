package main

import (
	"fmt"
	"math"
	"path/filepath"
	"strings"

	"cdf144/aoc2021/utils"
)

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	template := lines[0]
	rules := make(map[string]string)
	for _, line := range lines[2:] {
		var pair, insert string
		if _, err := fmt.Sscanf(line, "%s -> %s", &pair, &insert); err != nil {
			panic(err)
		}
		rules[pair] = insert
	}

	for range 10 {
		var newTemplate strings.Builder
		newTemplate.WriteByte(template[0])

		for i := range len(template) - 1 {
			pair := template[i : i+2]
			if insert, ok := rules[pair]; ok {
				newTemplate.WriteString(insert)
			}
			newTemplate.WriteByte(template[i+1])
		}

		template = newTemplate.String()
	}

	counts := make(map[rune]int)
	for _, c := range template {
		counts[c]++
	}

	min, max := math.MaxInt, 0
	for _, count := range counts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	printAnswer(max - min)
}
