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

func solveOne() {
    distances := getAllDistances()
    fmt.Println(aoc.Min(distances))
}

func solveTwo() {
    distances := getAllDistances()
    fmt.Println(aoc.Max(distances))
}

func getAllDistances() []int {
	input := aoc.GetInput(9, test)

    var m map[string]int = make(map[string]int)

    for _, line := range input {
        splitted := strings.Split(line, " = ")
        distance, _ := strconv.Atoi(splitted[1])
        m[splitted[0]] = distance
    }

    options := aoc.Permutate(getAllCities(m))

    distances := []int{}

    for _, option := range options {
        dist := calculateDistance(option, m)
        if dist != 0 {
            distances = append(distances, dist)
        }
    }    

    return distances
}

func calculateDistance(input []string, m map[string]int) int {
    result := 0

    for i := 0; i < len(input)-1; i++ {
        key := input[i] + " to " + input[i+1]
        distance, ok :=  m[key]
        if !ok {
            key2 := input[i+1] + " to " + input[i]
            distance, ok2 :=  m[key2]
            if !ok2 {
                return 0
            } else {
                result += distance
            }
        } else {
            result += distance
        }
    }
    return result
}

func getAllCities(input map[string]int) []string {
    result := []string{}
    for item := range input {
        items := strings.Split(item, " to ")
        result = append(result, items[0])
        result = append(result, items[1])
    }
    return aoc.ExtractUniques(result)
}

