package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("1.in")
	check(err)
	solveOne(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func solveOne(dat []byte) {
	ans := 0
	ans2 := 0
	for i, c := range string(dat) {
		if string(c) == "(" {
			ans++
		} else if string(c) == ")" {
			ans--
		}
		if ans == -1 && ans2 == 0 {
			ans2 = i + 1
		}
	}

	fmt.Println(ans)
	fmt.Println(ans2)
}
