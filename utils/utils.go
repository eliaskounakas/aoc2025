package util

import (
	"bufio"
	"os"
	"strings"
)

func ReadInput(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadCommaSeperatedInput(path string) []string {
	bytes, _ := os.ReadFile(path)

	return strings.Split(string(bytes), ",")
}
