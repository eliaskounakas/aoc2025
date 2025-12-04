package day04

import (
    "fmt"
    "io/ioutil"
)

func Solve() {
    data, _ := ioutil.ReadFile("days/day04/input.txt")
    input := string(data)
    fmt.Println("Day 4 input length:", len(input))
}
