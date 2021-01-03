package main

import (
	"fmt"
	aoc "github.com/nilventosa/adventOfCode2015"
	"strconv"
	"strings"
)

const test bool = false

func main() {
	solveOne()
	solveTwo()
}

func solveTwo() {
	table := [1000][1000]int{}
	input := aoc.GetInput(6, test)
	for _, line := range input {
		if strings.Contains(line, "off") {
			parts := strings.Split(line, " ")
			x1, _ := strconv.Atoi(strings.Split(parts[2], ",")[0])
			y1, _ := strconv.Atoi(strings.Split(parts[2], ",")[1])
			x2, _ := strconv.Atoi(strings.Split(parts[4], ",")[0])
			y2, _ := strconv.Atoi(strings.Split(parts[4], ",")[1])
			for ; x1 <= x2; x1++ {
				y1, _ = strconv.Atoi(strings.Split(parts[2], ",")[1])
				for ; y1 <= y2; y1++ {
					if table[x1][y1] > 0 {
						table[x1][y1]--
					}
				}
			}
		} else if strings.Contains(line, "on") {
			parts := strings.Split(line, " ")
			x1, _ := strconv.Atoi(strings.Split(parts[2], ",")[0])
			y1, _ := strconv.Atoi(strings.Split(parts[2], ",")[1])
			x2, _ := strconv.Atoi(strings.Split(parts[4], ",")[0])
			y2, _ := strconv.Atoi(strings.Split(parts[4], ",")[1])
			for ; x1 <= x2; x1++ {
				y1, _ = strconv.Atoi(strings.Split(parts[2], ",")[1])
				for ; y1 <= y2; y1++ {
					table[x1][y1]++
				}
			}
		} else {
			parts := strings.Split(line, " ")
			x1, _ := strconv.Atoi(strings.Split(parts[1], ",")[0])
			y1, _ := strconv.Atoi(strings.Split(parts[1], ",")[1])
			x2, _ := strconv.Atoi(strings.Split(parts[3], ",")[0])
			y2, _ := strconv.Atoi(strings.Split(parts[3], ",")[1])
			for ; x1 <= x2; x1++ {
				y1, _ = strconv.Atoi(strings.Split(parts[1], ",")[1])
				for ; y1 <= y2; y1++ {
					table[x1][y1] += 2
				}
			}
		}
	}
	fmt.Println(countTwo(table))
}

func solveOne() {
	table := [1000][1000]bool{}
	input := aoc.GetInput(6, test)
	for _, line := range input {
		if strings.Contains(line, "off") {
			parts := strings.Split(line, " ")
			x1, _ := strconv.Atoi(strings.Split(parts[2], ",")[0])
			y1, _ := strconv.Atoi(strings.Split(parts[2], ",")[1])
			x2, _ := strconv.Atoi(strings.Split(parts[4], ",")[0])
			y2, _ := strconv.Atoi(strings.Split(parts[4], ",")[1])
			for ; x1 <= x2; x1++ {
				y1, _ = strconv.Atoi(strings.Split(parts[2], ",")[1])
				for ; y1 <= y2; y1++ {
					table[x1][y1] = false
				}
			}
		} else if strings.Contains(line, "on") {
			parts := strings.Split(line, " ")
			x1, _ := strconv.Atoi(strings.Split(parts[2], ",")[0])
			y1, _ := strconv.Atoi(strings.Split(parts[2], ",")[1])
			x2, _ := strconv.Atoi(strings.Split(parts[4], ",")[0])
			y2, _ := strconv.Atoi(strings.Split(parts[4], ",")[1])
			for ; x1 <= x2; x1++ {
				y1, _ = strconv.Atoi(strings.Split(parts[2], ",")[1])
				for ; y1 <= y2; y1++ {
					table[x1][y1] = true
				}
			}
		} else {
			parts := strings.Split(line, " ")
			x1, _ := strconv.Atoi(strings.Split(parts[1], ",")[0])
			y1, _ := strconv.Atoi(strings.Split(parts[1], ",")[1])
			x2, _ := strconv.Atoi(strings.Split(parts[3], ",")[0])
			y2, _ := strconv.Atoi(strings.Split(parts[3], ",")[1])
			for ; x1 <= x2; x1++ {
				y1, _ = strconv.Atoi(strings.Split(parts[1], ",")[1])
				for ; y1 <= y2; y1++ {
					table[x1][y1] = !table[x1][y1]
				}
			}
		}
	}
	fmt.Println(countOne(table))
}

func countOne(table [1000][1000]bool) int {
	counter := 0
	for _, j := range table {
		for _, i := range j {
			if i {
				counter++
			}
		}
	}
	return counter
}

func countTwo(table [1000][1000]int) int {
	counter := 0
	for _, j := range table {
		for _, i := range j {
			counter += i
		}
	}
	return counter
}