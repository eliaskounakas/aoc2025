package solutions

import (
	util "aoc-2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day02() {
	input := util.ReadCommaSeperatedInput("inputs/day02.txt")

	fmt.Println("Day 2 Part 1:", solveDay2Part1(input))
	fmt.Println("Day 2 Part 2:", solveDay2Part2(input))
}

func solveDay2Part1(input []string) int {
	invalidIDsSum := 0

	for i := range input {
		ids := strings.Split(input[i], "-")
		firstID, err := strconv.Atoi(ids[0])

		if err != nil {
			return 0
		}

		secondID, err := strconv.Atoi(ids[1])

		if err != nil {
			return 0
		}

		for currID := firstID; currID <= secondID; currID++ {
			strID := strconv.Itoa(currID)

			if len(strID)%2 == 1 {
				continue
			}

			firstHalf := strID[:len(strID)/2]
			secondHalf := strID[len(strID)/2:]

			if firstHalf == secondHalf {
				invalidIDsSum += currID
			}
		}
	}

	return invalidIDsSum
}

func solveDay2Part2(input []string) int {

	return 0
}
