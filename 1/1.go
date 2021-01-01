package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
)

func main() {
	input := getInput(1, false)
	solve(input)
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
