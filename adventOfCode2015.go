package adventofcode

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

func Max(values []int) int {
	max := 0
	for i, e := range values {
		if i == 0 || e > max {
			max = e
		}
	}
	return max
}

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


func Permutate(input []string) [][]string {
	var permute func(int, []string)
	result := [][]string{}


	permute = func(k int, input []string) {
		if k == 1 {
			part := make([]string, len(input))
			copy(part, input)
			result = append(result, part)
		} else {
			permute(k-1, input)
			for i := 0; i < k-1; i++ {
				if k % 2 == 1 {
					input[0], input[k-1] = input[k-1], input[0]
				} else {
					input[i], input[k-1] = input[k-1], input[i]
				}
				permute(k-1, input)
			}
		}
	}
	permute(len(input), input)
	return result
}
