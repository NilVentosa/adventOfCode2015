package main

import (
	"fmt"

	aoc "github.com/nilventosa/adventOfCode2015"
)

func main() {
	input := aoc.GetInput(1, false)
	solve(input)
}

func solve(input []string) {
	ans := 0
	ans2 := 0

	for index, char := range input[0] {
		if string(char) == "(" {
			ans++
		} else if string(char) == ")" {
			ans--
		}
		if ans == -1 && ans2 == 0 {
			ans2 = index + 1
		}
	}

	fmt.Println(ans)
	fmt.Println(ans2)
}
