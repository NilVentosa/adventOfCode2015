package main

import (
	"fmt"
	aoc "github.com/nilventosa/adventOfCode2015"
	"strconv"
)

var isTest bool = false

func main() {
	solveOne()
	solveTwo()
}

func solveOne() {
	input := aoc.GetInput(8, isTest)

	lit := 0
	mem := 0

	for _, line := range input {
		lit += len(line)
		m, _ := strconv.Unquote(line)
		mem += len(m)

	}

	fmt.Println(lit - mem)
}

func solveTwo() {
	input := aoc.GetInput(8, isTest)
	norm := 0
	quoted := 0

	for _, line := range input {
		norm += len(line)
		quoted += len(strconv.Quote(line))
	}

	fmt.Println(quoted - norm)
}
