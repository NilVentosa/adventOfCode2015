package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
)

func main() {
	dat := getInput(1, false)
	solve(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(day int, isTest bool) []byte {
	_, currentFileName, _, _ := runtime.Caller(0)
	filePath := path.Dir(currentFileName)
	var inputFileName string

	if isTest {
		inputFileName = strconv.Itoa(day) + "_test" + ".in"
	} else {
		inputFileName = strconv.Itoa(day) + ".in"
	}

	dat, err := ioutil.ReadFile(filePath + "/" + inputFileName)
	check(err)
	return dat
}

func solve(dat []byte) {
	ans := 0
	ans2 := 0

	for index, char := range string(dat) {
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
