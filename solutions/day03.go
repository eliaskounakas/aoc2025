package solutions

import (
	util "aoc-2025/utils"
	"fmt"
	"math"
)

func Day03() {
	input := util.ReadInput("inputs/day03.txt")

	fmt.Println("Day 3 Part 1:", solveDay3Part1(input))
	fmt.Println("Day 3 Part 2:", solveDay3Part2(input))
}

func solveDay3Part1(input []string) int {
	totalJoltage := 0

	for _, s := range input {
		totalJoltage += findJoltage(s, 1, 0)
	}

	return totalJoltage
}

func solveDay3Part2(input []string) int {
	totalJoltage := 0

	for _, s := range input {
		totalJoltage += findJoltage(s, 11, 0)
	}

	return totalJoltage
}

func findJoltage(s string, level int, index int) int {
	max := int(s[index] - '0')
	maxIndex := index

	for ; index < len(s)-level; index++ {
		curr := int(s[index] - '0')
		if curr > max {
			max = curr
			maxIndex = index
		}
	}

	if level == 0 {
		return max
	}

	return max*int(math.Pow(10, float64(level))) + findJoltage(s, level-1, maxIndex+1)
}
