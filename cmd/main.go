
package main

import (
    "fmt"
    "os"
    "github.com/eliaskounakas/aoc2025/days/day04"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: aoc2025 <day>")
        return
    }

    day := os.Args[1]

    switch day {
    case "4":
        day04.Solve()
    default:
        fmt.Println("Day not implemented yet")
    }
}
