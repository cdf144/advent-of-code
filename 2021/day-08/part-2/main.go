package main

import (
	"path/filepath"
	"slices"
	"strings"

	"cdf144/aoc2021/utils"
)

var displayToDigit = map[string]rune{
	"abcefg":  '0',
	"cf":      '1',
	"acdeg":   '2',
	"acdfg":   '3',
	"bcdf":    '4',
	"abdfg":   '5',
	"abdefg":  '6',
	"acf":     '7',
	"abcdefg": '8',
	"abcdfg":  '9',
}

func findOutputValue(entry string) int {
	parts := strings.Split(entry, " | ")
	signalPatterns := strings.Fields(parts[0])
	outputDigits := strings.Fields(parts[1])

	signalToSegment := make(map[rune]rune, 7)

	signalFreq := make(map[rune]int, 7)
	for _, pattern := range signalPatterns {
		for _, signal := range pattern {
			signalFreq[signal]++
		}
	}

	// Segments b, e, f have unique frequencies (6, 4, 9 respectively)
	for signal, freq := range signalFreq {
		switch freq {
		case 6:
			signalToSegment[signal] = 'b'
		case 4:
			signalToSegment[signal] = 'e'
		case 9:
			signalToSegment[signal] = 'f'
		}
	}

	// Patterns for digits that are unique in number of segments
	pattern1 := findPatternByLength(signalPatterns, 2)
	pattern4 := findPatternByLength(signalPatterns, 4)
	pattern7 := findPatternByLength(signalPatterns, 3)

	// Pattern for digit 1 have two signals that map to 'c' and 'f'
	// Because one is already mapped to 'f', the other must be mapped to 'c'
	if _, mapped := signalToSegment[rune(pattern1[0])]; mapped {
		signalToSegment[rune(pattern1[1])] = 'c'
	} else {
		signalToSegment[rune(pattern1[0])] = 'c'
	}

	// Pattern for digit 7 have three signals that map to 'a', 'c' and 'f'
	// The only one that hasn't been mapped is the one that maps to 'a'
	for _, signal := range pattern7 {
		if _, mapped := signalToSegment[signal]; !mapped {
			signalToSegment[signal] = 'a'
			break
		}
	}

	// Pattern for digit 4 have four signals that map to 'b', 'c', 'd' and 'f'
	// Similarly we isolate the one that maps to 'd'
	for _, signal := range pattern4 {
		if _, mapped := signalToSegment[signal]; !mapped {
			signalToSegment[signal] = 'd'
			break
		}
	}

	// The only remaining unmapped signal is that which maps to 'g'
	for signal := 'a'; signal <= 'g'; signal++ {
		if _, mapped := signalToSegment[signal]; !mapped {
			signalToSegment[signal] = 'g'
			break
		}
	}

	// Hooray mapping done
	var resultStr strings.Builder
	for _, digit := range outputDigits {
		decodedDisplay := make([]rune, len(digit))
		for i, signal := range digit {
			decodedDisplay[i] = signalToSegment[signal]
		}
		slices.Sort(decodedDisplay)
		resultStr.WriteRune(displayToDigit[string(decodedDisplay)])
	}

	return utils.Atoi(resultStr.String())
}

func findPatternByLength(signalPatterns []string, length int) string {
	for _, pattern := range signalPatterns {
		if len(pattern) == length {
			return pattern
		}
	}
	return ""
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	var answer int
	for _, line := range lines {
		answer += findOutputValue(line)
	}

	printAnswer(answer)
}
