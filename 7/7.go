package main

import (
	"fmt"
	aoc "github.com/nilventosa/adventOfCode2015"
	"strconv"
	"strings"
)

const isTest bool = false

func main() {
	solveOne()
	solveTwo()
}

var result1 = make(map[string]int64)
var result2 = make(map[string]int64)

func solveTwo() {
	input2 := aoc.GetInput(7, isTest)

  for i := 0; i < len(input2); i++ {
    if input2[i] == "14146 -> b" {
      result2["b"] = result1["a"]
      input2 = append(input2[:i], input2[i+1:]...)
      break
    }
  }

	for len(input2) != 0 {
		for i := len(input2) - 1; i >= 0; i-- {
			key, value, ok := parseLinePart2(input2[i])
			if ok {
				result2[key] = value
				input2 = append(input2[:i], input2[i+1:]...)
			}
		}
	}
	fmt.Println(result2["a"])
}

func solveOne() {
	input := aoc.GetInput(7, isTest)

	for len(input) != 0 {
		for i := len(input) - 1; i >= 0; i-- {
			key, value, ok := parseLinePart1(input[i])
			if ok {
				result1[key] = value
				input = append(input[:i], input[i+1:]...)
			}
		}
	}
	fmt.Println(result1["a"])
}

func parseLinePart1(line string) (key string, value int64, ok bool) {
	leftright := strings.Split(line, " -> ")
	key = leftright[1]

	parts := strings.Split(leftright[0], " ")
	if len(parts) == 1 {
		value, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			value, ok := result1[parts[0]]
			if ok {
				return key, value, true
			}
			return "", 0, false
		}

		return key, value, true

	} else if len(parts) == 2 {
		value, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			value, ok := result1[parts[1]]
			if ok {
				return key, ^value, true
			}
			return "", 0, false
		}

		return key, ^value, true

	} else if len(parts) == 3 {
		valueL, errL := strconv.ParseInt(parts[0], 10, 64)
		valueR, errR := strconv.ParseInt(parts[2], 10, 64)

		if errL != nil {
			_, ok := result1[parts[0]]
			if !ok {
				return "", 0, false

			}
			valueL, _ = result1[parts[0]]
		}

		if errR != nil {
			_, ok := result1[parts[2]]
			if !ok {
				return "", 0, false

			}
			valueR, _ = result1[parts[2]]
		}

		switch {
		case parts[1] == "LSHIFT":
			return key, valueL << valueR, true

		case parts[1] == "RSHIFT":
			return key, valueL >> valueR, true

		case parts[1] == "AND":
			return key, valueL & valueR, true

		case parts[1] == "OR":
			return key, valueL | valueR, true

		}
	}
	return "", 0, false

}


func parseLinePart2(line string) (key string, value int64, ok bool) {
	leftright := strings.Split(line, " -> ")
	key = leftright[1]

	parts := strings.Split(leftright[0], " ")
	if len(parts) == 1 {
		value, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			value, ok := result2[parts[0]]
			if ok {
				return key, value, true
			}
			return "", 0, false
		}

		return key, value, true

	} else if len(parts) == 2 {
		value, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			value, ok := result2[parts[1]]
			if ok {
				return key, ^value, true
			}
			return "", 0, false
		}

		return key, ^value, true

	} else if len(parts) == 3 {
		valueL, errL := strconv.ParseInt(parts[0], 10, 64)
		valueR, errR := strconv.ParseInt(parts[2], 10, 64)

		if errL != nil {
			_, ok := result2[parts[0]]
			if !ok {
				return "", 0, false

			}
			valueL, _ = result2[parts[0]]
		}

		if errR != nil {
			_, ok := result2[parts[2]]
			if !ok {
				return "", 0, false

			}
			valueR, _ = result2[parts[2]]
		}

		switch {
		case parts[1] == "LSHIFT":
			return key, valueL << valueR, true

		case parts[1] == "RSHIFT":
			return key, valueL >> valueR, true

		case parts[1] == "AND":
			return key, valueL & valueR, true

		case parts[1] == "OR":
			return key, valueL | valueR, true

		}
	}
	return "", 0, false

}
