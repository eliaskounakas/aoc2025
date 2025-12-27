package solutions

import (
	util "aoc-2025/utils"
	"fmt"
	"strconv"
)

func Day01() {
	input := util.ReadInput("inputs/day01.txt")

	fmt.Println("Day 1 Part 1:", solvePart1(input))
	fmt.Println("Day 1 Part 2:", solvePart2(input))
}

func solvePart1(input []string) int {
	dial := 50
	zeros := 0

	for i := range input {
		curr := input[i]
		left := curr[0] == 'L'
		count, err := strconv.Atoi(curr[1:])

		if err != nil {
			continue
		}

		for range count {
			if left {
				dial--
			} else {
				dial++
			}

			if dial < 0 {
				dial = 99
			} else if dial > 99 {
				dial = 0
			}
		}

		if dial == 0 {
			zeros++
		}
	}

	return zeros
}

func solvePart2(input []string) int {
	dial := 50
	zeros := 0

	for i := range input {
		curr := input[i]
		left := curr[0] == 'L'
		count, err := strconv.Atoi(curr[1:])

		if err != nil {
			continue
		}

		for range count {
			if left {
				dial--
			} else {
				dial++
			}

			if dial < 0 {
				dial = 99
			} else if dial > 99 {
				dial = 0
			}

			// Difference is here
			if dial == 0 {
				zeros++
			}
		}
	}

	return zeros
}
