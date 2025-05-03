package main

import (
	"path/filepath"
	"strconv"
	"strings"

	"cdf144/aoc2021/utils"
)

var displayToDigit = map[string]int{
	"1110111": 0,
	"0010010": 1,
	"1011101": 2,
	"1011011": 3,
	"0111010": 4,
	"1101011": 5,
	"1101111": 6,
	"1010010": 7,
	"1111111": 8,
	"1111011": 9,
}

func findOutputValue(line string) int {
	parts := strings.Split(line, " | ")
	signalPatterns := strings.Fields(parts[0])
	outputDigits := strings.Fields(parts[1])

	signalToSegment := make(map[string]int)

	// Signals which map to segment 3 (c) and 6 (f)
	for _, pattern := range signalPatterns {
		if len(pattern) == 2 {
			split := strings.Split(pattern, "")
			signal1, signal2 := split[0], split[1]
			signal1Freq := countFrequency(signal1, signalPatterns)
			signal2Freq := countFrequency(signal2, signalPatterns)

			if signal1Freq == 8 {
				signalToSegment[signal1] = 2
				signalToSegment[signal2] = 5
			} else if signal2Freq == 8 {
				signalToSegment[signal1] = 5
				signalToSegment[signal2] = 2
			}

			break
		}
	}

	// Signal which maps to segment 1 (a)
findSignalA:
	for _, pattern := range signalPatterns {
		if len(pattern) == 3 {
			for signal := range strings.SplitSeq(pattern, "") {
				_, ok := signalToSegment[signal]
				if !ok {
					signalToSegment[signal] = 0
					break findSignalA
				}
			}
		}
	}

	// Signals which map to segment 2 (b) and 4 (d)
	for _, pattern := range signalPatterns {
		if len(pattern) == 4 {
			split := strings.Split(pattern, "")
			unmapped1, unmapped2 := "", ""
			signalToSegment[unmapped1] = 5
			signalToSegment[unmapped2] = 7

			for _, signal := range split {
				_, ok := signalToSegment[signal]
				if !ok {
					if unmapped1 == "" {
						unmapped1 = signal
					} else {
						unmapped2 = signal
					}
				}
			}

			unmapped1Freq := countFrequency(unmapped1, signalPatterns)
			unmapped2Freq := countFrequency(unmapped2, signalPatterns)

			if unmapped1Freq == 6 {
				signalToSegment[unmapped1] = 1
				signalToSegment[unmapped2] = 3
			} else if unmapped2Freq == 6 {
				signalToSegment[unmapped1] = 3
				signalToSegment[unmapped2] = 1
			}

			break
		}
	}

	// Signals which map to segment 5 (e) and 7 (g)
	for _, pattern := range signalPatterns {
		if len(pattern) == 7 {
			split := strings.Split(pattern, "")
			unmapped1, unmapped2 := "", ""

			for _, signal := range split {
				_, ok := signalToSegment[signal]
				if !ok {
					if unmapped1 == "" {
						unmapped1 = signal
					} else {
						unmapped2 = signal
					}
				}
			}

			unmapped1Freq := countFrequency(unmapped1, signalPatterns)
			unmapped2Freq := countFrequency(unmapped2, signalPatterns)

			if unmapped1Freq == 4 {
				signalToSegment[unmapped1] = 4
				signalToSegment[unmapped2] = 6
			} else if unmapped2Freq == 4 {
				signalToSegment[unmapped1] = 6
				signalToSegment[unmapped2] = 4
			}

			break
		}
	}

	// Map done (finally)
	resultStr := ""

	for _, digit := range outputDigits {
		display := []string{"0", "0", "0", "0", "0", "0", "0"}

		for signal := range strings.SplitSeq(digit, "") {
			display[signalToSegment[signal]] = "1"
		}

		resultStr += strconv.Itoa(
			displayToDigit[strings.Join(display, "")],
		)
	}

	return utils.Atoi(resultStr)
}

func countFrequency(signal string, signalPatterns []string) int {
	var frequency int
	for _, pattern := range signalPatterns {
		if strings.Contains(pattern, signal) {
			frequency++
		}
	}
	return frequency
}

func main() {
	lines, printAnswer := utils.Init(filepath.Join("..", "input.txt"))

	var answer int

	for _, line := range lines {
		answer += findOutputValue(line)
	}

	printAnswer(answer)
}
