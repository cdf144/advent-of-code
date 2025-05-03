package main

import (
	"path/filepath"
	"slices"
	"strings"

	"cdf144/aoc2021/utils"
)

var displayToDigit = map[string]string{
	"abcefg":  "0",
	"cf":      "1",
	"acdeg":   "2",
	"acdfg":   "3",
	"bcdf":    "4",
	"abdfg":   "5",
	"abdefg":  "6",
	"acf":     "7",
	"abcdefg": "8",
	"abcdfg":  "9",
}

func findOutputValue(line string) int {
	parts := strings.Split(line, " | ")
	signalPatterns := strings.Fields(parts[0])
	outputDigits := strings.Fields(parts[1])

	signalToSegment := make(map[string]string)

	// Signals which map to segment c and f
	// Analyze pattern for digit 1 and do frequency analysis
	for _, pattern := range signalPatterns {
		if len(pattern) == 2 {
			split := strings.Split(pattern, "")
			signal1, signal2 := split[0], split[1]
			signal1Freq := countFrequency(signal1, signalPatterns)
			signal2Freq := countFrequency(signal2, signalPatterns)

			if signal1Freq == 8 {
				signalToSegment[signal1] = "c"
				signalToSegment[signal2] = "f"
			} else if signal2Freq == 8 {
				signalToSegment[signal1] = "f"
				signalToSegment[signal2] = "c"
			}

			break
		}
	}

	// Signal which maps to segment a
	// Digit 7 has two overlapping segments with digit 1 so we can isolate the third segment
findSignalA:
	for _, pattern := range signalPatterns {
		if len(pattern) == 3 {
			for signal := range strings.SplitSeq(pattern, "") {
				_, ok := signalToSegment[signal]
				if !ok {
					signalToSegment[signal] = "a"
					break findSignalA
				}
			}
		}
	}

	// Signals which map to segment b and d
	// Digit 4 has two overlapping segments with digit 1 so we isolate the other two
	// segments and do frequency analysis
	for _, pattern := range signalPatterns {
		if len(pattern) == 4 {
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

			if unmapped1Freq == 6 {
				signalToSegment[unmapped1] = "b"
				signalToSegment[unmapped2] = "d"
			} else if unmapped2Freq == 6 {
				signalToSegment[unmapped1] = "d"
				signalToSegment[unmapped2] = "b"
			}

			break
		}
	}

	// Signals which map to segment e and g
	// All segment mappings are done except for e and g. We can isolate the two
	// remaining segments and do frequency analysis.
	for _, pattern := range signalPatterns {
		// Digit 8 is the last with unique number of segments so I just use it
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
				signalToSegment[unmapped1] = "e"
				signalToSegment[unmapped2] = "g"
			} else if unmapped2Freq == 4 {
				signalToSegment[unmapped1] = "g"
				signalToSegment[unmapped2] = "e"
			}

			break
		}
	}

	// Map done (finally)
	resultStr := strings.Builder{}

	for _, digit := range outputDigits {
		display := make([]string, len(digit))
		for i, signal := range strings.Split(digit, "") {
			display[i] = signalToSegment[signal]
		}

		slices.Sort(display)
		resultStr.WriteString(displayToDigit[strings.Join(display, "")])
	}

	return utils.Atoi(resultStr.String())
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
	lines, printAnswer := utils.Init(filepath.Join("..", "input.example.txt"))

	var answer int

	for _, line := range lines {
		answer += findOutputValue(line)
	}

	printAnswer(answer)
}
