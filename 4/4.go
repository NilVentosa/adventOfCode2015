package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/nilventosa/adventOfCode2015"
)

const test bool = false

func main() {
	solveOne()
	solveTwo()
}

func firstWithPrefix(thing string, prefix string) int {
	for i := 0; true; i++ {
		h := md5.Sum([]byte(thing + strconv.Itoa(i)))
		if strings.HasPrefix(fmt.Sprintf("%x", h), prefix) {
			return i
		}
	}
	return 0
}

func solveOne() {
	input := aoc.GetInput(4, test)
	fmt.Println(firstWithPrefix(input[0], "00000"))
}

func solveTwo() {
	input := aoc.GetInput(4, test)
	fmt.Println(firstWithPrefix(input[0], "000000"))
}
