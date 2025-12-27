package main

import (
	"aoc-2025/solutions" // Update with your module name
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]

	switch day {
	case "1":
		solutions.Day01()
	case "2":
		solutions.Day02()
	default:
		fmt.Printf("Day %s not implemented yet\n", day)
	}
}
