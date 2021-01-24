package main

import (
	"fmt"
	"strings"
)

var alphabet string = "abcdefghijklmnopqrstuvwxyz"

func main() {
	// input := aoc.GetInput(11, false)[0]
	input := "aabcc"

	fmt.Println(isPasswordValid(input))
}

func solveOne(input string) string {
	return input
}

func isPasswordValid(input string) bool {
	return passFirstRule(input) && passSecondRule(input) && passThirdRule(input)
}

func passFirstRule(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if strings.Contains(alphabet, input[i:i+3]) {
			return true
		}
	}
	return false
}

func passSecondRule(input string) bool {
	if strings.Contains(input, "i") || strings.Contains(input, "l") || strings.Contains(input, "o") {
		return false
	}
	return true
}

func passThirdRule(input string) bool {
	return groupsOfTwo(input) > 1
}

func groupsOfTwo(input string) int {
	result := 0
	current := ""
	currentAmount := 0

	for _, c := range input {
		if current == string(c) {
			currentAmount++
			if currentAmount > 0 {
				result++
				currentAmount = 0
			}
		} else {
			currentAmount = 0
		}
		current = string(c)
	}

	return result
}
