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
			if isInvalidID(currID) {
				invalidIDsSum += currID
			}
		}
	}

	return invalidIDsSum
}

func isInvalidID(id int) bool {
	strID := strconv.Itoa(id)

	for subStringSize := 1; subStringSize <= len(strID)/2; subStringSize++ {
		if subStringSize == 0 || (len(strID)%subStringSize) != 0 {
			continue
		}

		subStrings := chunkString(strID, subStringSize)

		invalid := true
		for _, s := range subStrings {
			if s != subStrings[0] {
				invalid = false
				break
			}
		}

		if invalid {
			fmt.Println(strID)
			return true
		}
	}

	return false
}

// Splits a string into substrings of size argument.
// for example: s is a string of length 8, if size is 2,
// it will get split into 4 strings of length 2
func chunkString(s string, size int) []string {
	var chunks []string
	for i := 0; i < len(s); i += size {
		end := min(i+size, len(s))
		chunks = append(chunks, s[i:end])
	}
	return chunks
}
