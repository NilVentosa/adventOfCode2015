package adventOfCode2015

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInput(day int, isTest bool) []string {
	_, currentFileName, _, _ := runtime.Caller(0)
	filePath := path.Dir(currentFileName)
	var inputFileName string

	if isTest {
		inputFileName = strconv.Itoa(day) + "/" + strconv.Itoa(day) + "_test" + ".in"
	} else {
		inputFileName = strconv.Itoa(day) + "/" + strconv.Itoa(day) + ".in"
	}

	file, err := os.Open(filePath + "/" + inputFileName)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func Min(values []int) int {
	min := 0
	for i, e := range values {
		if i == 0 || e < min {
			min = e
		}
	}
	return min
}
