package main

import (
	"fmt"

	aoc "github.com/nilventosa/adventOfCode2015"
)

const test bool = false

func main() {
	solveOne()
	solveTwo()
}

func solveTwo() {
	santaX := 0
	santaY := 0

	robotX := 0
	robotY := 0

	visited := make(map[string]struct{})
	visited = addIfFirstVisit(santaX, santaY, visited)

	input := aoc.GetInput(3, test)
	line := input[0]

	for index, b := range line {
		if index%2 == 0 {
			x, y := move(b)
			santaX += x
			santaY += y
			visited = addIfFirstVisit(santaX, santaY, visited)
		} else {
			x, y := move(b)
			robotX += x
			robotY += y
			visited = addIfFirstVisit(robotX, robotY, visited)
		}
	}

	fmt.Println(len(visited))
}

func solveOne() {
	currentX := 0
	currentY := 0

	visited := make(map[string]struct{})
	visited = addIfFirstVisit(currentX, currentY, visited)

	input := aoc.GetInput(3, test)
	line := input[0]

	for _, b := range line {
		x, y := move(b)
		currentX += x
		currentY += y
		visited = addIfFirstVisit(currentX, currentY, visited)
	}
	fmt.Println(len(visited))
}

func addIfFirstVisit(x int, y int, visited map[string]struct{}) map[string]struct{} {
	currentString := fmt.Sprint(x) + "," + fmt.Sprint(y)
	_, ok := visited[currentString]

	if !ok {
		visited[currentString] = struct{}{}
	}

	return visited
}

func move(b int32) (int, int) {
	arrow := string(b)
	x := 0
	y := 0
	if arrow == ">" {
		x++
	} else if arrow == "<" {
		x--
	} else if arrow == "^" {
		y++
	} else if arrow == "v" {
		y--
	}
	return x, y
}
