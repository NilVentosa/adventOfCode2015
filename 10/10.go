package main

import (
    "strings"
    "strconv"
    "fmt"
	aoc "github.com/nilventosa/adventOfCode2015"
)

var isTest bool = false

func main() {
    solve()
}

func solve() {
    input := aoc.GetInput(10, isTest)
    current := input[0]

    for i := 0; i < 50; i++ {
        if i == 40 {
            fmt.Println(len(current))
        }
        current = lookAndSay(current)
    }

    fmt.Println(len(current))
}

func lookAndSay(num string) string {
    result := ""
    splitted := strings.Split(num, "")
    current := ""
    currentAmount := 1

    for _, item := range splitted {
        if item == current {
            currentAmount++
        } else {
            if currentAmount > 0 && current != ""{
                result += strconv.Itoa(currentAmount)
                result += current
                currentAmount = 1
            }
            current = item
        }
    }
    result += strconv.Itoa(currentAmount)
    result += current

    return result
}
