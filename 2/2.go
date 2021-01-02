package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/nilventosa/adventOfCode2015"
)

type present struct {
	L, W, H int
}

func main() {
	input := aoc.GetInput(2, false)
	presents := extractPresents(input)
	solve(presents)
}

func solve(presents []present) {
	resultOne := 0
	resultTwo := 0

	for _, present := range presents {
		resultOne += getTotalArea(present)
	}

	for _, present := range presents {
		resultTwo += getRibbonLength(present)
	}

	fmt.Println(resultOne)
	fmt.Println(resultTwo)
}

func getRibbonLength(present present) int {
	per1 := present.L + present.L + present.H + present.H
	per2 := present.W + present.W + present.H + present.H
	per3 := present.L + present.L + present.W + present.W

	return aoc.Min([]int{per1, per2, per3}) + (present.H * present.L * present.W)
}

func getTotalArea(present present) int {
	side1 := 2 * present.L * present.W
	side2 := 2 * present.W * present.H
	side3 := 2 * present.H * present.L
	smallSide := aoc.Min([]int{side1, side2, side3}) / 2

	return side1 + side2 + side3 + smallSide
}

func extractPresents(input []string) []present {
	result := []present{}

	for _, line := range input {
		s := strings.Split(line, "x")

		l, _ := strconv.Atoi(s[0])
		w, _ := strconv.Atoi(s[1])
		h, _ := strconv.Atoi(s[2])
		p := present{l, w, h}

		result = append(result, p)
	}
	return result
}
