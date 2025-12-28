package solutions

import (
	util "aoc-2025/utils"
	"fmt"
)

func Day03() {
	input := util.ReadInput("inputs/day03.txt")

	fmt.Println("Day 3 Part 1:", solveDay3Part1(input))
	fmt.Println("Day 3 Part 2:", solveDay3Part2(input))
}

func solveDay3Part1(input []string) int {
	totalJoltage := 0

	for _, s := range input {
		max := int(s[0] - '0')
		maxIndex := 0

		for i := range s {
			if i >= len(s)-1 {
				break
			}

			curr := int(s[i] - '0')
			if curr > max {
				max = curr
				maxIndex = i
			}
		}

		firstDigit := max

		max = int(s[maxIndex+1] - '0')
		for i := maxIndex + 2; i < len(s); i++ {
			curr := int(s[i] - '0')
			if curr > max {
				max = curr
			}
		}

		secondDigit := max

		fmt.Println(s)
		fmt.Println(firstDigit, secondDigit)

		totalJoltage += firstDigit*10 + secondDigit
	}

	return totalJoltage
}

func solveDay3Part2(input []string) int {
	return 0
}
