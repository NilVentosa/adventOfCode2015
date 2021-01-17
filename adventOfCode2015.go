package adventofcode

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strconv"
)

// Given a error panics if the error is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Given the number of day and bool if it or not a test returns a slice of the input lines
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

// Given a slice of ints returns the min value
func Min(values []int) int {
	min := 0
	for i, e := range values {
		if i == 0 || e < min {
			min = e
		}
	}
	return min
}

// Given a slice of ints returns the max value
func Max(values []int) int {
	max := 0
	for i, e := range values {
		if i == 0 || e > max {
			max = e
		}
	}
	return max
}

// Given a slice of strings removes all duplicates
func ExtractUniques(input []string) []string {
	seen := make(map[string]struct{}, len(input))
	j := 0
	for _, v := range input {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		input[j] = v
		j++
	}
	return input[:j]
}

// Given a slice of strings returns a slice of all possible permutations
func Permutate(input []string) [][]string {
	result := [][]string{}
	var helper func(int, []string)

	helper = func(k int, input []string) {
		if k == 1 {
			part := make([]string, len(input))
			copy(part, input)
			result = append(result, part)
		} else {
			helper(k-1, input)
			for i := 0; i < k-1; i++ {
				if k % 2 == 1 {
					input[0], input[k-1] = input[k-1], input[0]
				} else {
					input[i], input[k-1] = input[k-1], input[i]
				}
				helper(k-1, input)
			}
		}
	}
	helper(len(input), input)
	return result
}
