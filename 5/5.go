package main

import (
	"fmt"
	aoc "github.com/nilventosa/adventOfCode2015"
	"strings"
)

const test bool = false

var vowels = "aeiou"

func main() {
	solveOne()
	solveTwo()
}

func solveOne() {
	input := aoc.GetInput(5, test)
	nice := 0
	for _, line := range input {
		if isLineNicePart1(line) {
			nice++
		}
	}
	fmt.Println(nice)
}

func solveTwo() {
	input := aoc.GetInput(5, test)
	nice := 0
	for _, line := range input {
		if isLineNicePart2(line) {
			nice++
		}
	}
	fmt.Println(nice)
}

func isLineNicePart2(line string) bool {
	hasPair := false
	for i := 2; i <= len(line); i++ {
		if strings.Count(line, line[i-2:i]) > 1 {
			hasPair = true
			break
		}
	}
	if !hasPair {
		return hasPair
	}

	hasBetweener := false
	for i := 3; i <= len(line); i++ {
		if line[i-3:i-2] == line[i-1:i] {
			hasBetweener = true
			break
		}
	}
	return hasBetweener
}

func isLineNicePart1(line string) bool {
	vowelCount := 0
	prevLetter := ""
	twoInARowResult := false
	for _, letter := range line {
		if strings.Contains(vowels, string(letter)) {
			vowelCount++
		}
		if string(letter) == prevLetter {
			twoInARowResult = true
		}
		prevLetter = string(letter)
	}

	if vowelCount < 3 {
		return false
	}

	if !twoInARowResult {
		return false
	}

	if strings.Contains(line, "ab") {
		return false
	}
	if strings.Contains(line, "cd") {
		return false
	}
	if strings.Contains(line, "pq") {
		return false
	}
	if strings.Contains(line, "xy") {
		return false
	}
	return true
}
