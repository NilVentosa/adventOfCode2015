package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type present struct {
	L, W, H int
}

func main() {
	input := getInput(2, false)
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
		resultTwo += getRibbon(present)
	}

	fmt.Println(resultOne)
	fmt.Println(resultTwo)
}

func getRibbon(present present) int {
	//smallest perimeter plus volume
	per1 := present.L + present.L + present.H + present.H
	per2 := present.W + present.W + present.H + present.H
	per3 := present.L + present.L + present.W + present.W

	return min([]int{per1, per2, per3}) + (present.H * present.L * present.W)
}

func getTotalArea(present present) int {
	side1 := 2 * present.L * present.W
	side2 := 2 * present.W * present.H
	side3 := 2 * present.H * present.L
	smallSide := min([]int{side1, side2, side3}) / 2

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(day int, isTest bool) []string {
	_, currentFileName, _, _ := runtime.Caller(0)
	filePath := path.Dir(currentFileName)
	var inputFileName string

	if isTest {
		inputFileName = strconv.Itoa(day) + "_test" + ".in"
	} else {
		inputFileName = strconv.Itoa(day) + ".in"
	}

	file, err := os.Open(filePath + "/" + inputFileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func min(values []int) int {
	min := 0
	for i, e := range values {
		if i == 0 || e < min {
			min = e
		}
	}
	return min
}
