package main

import (
	day11 "advent_of_code/day_11"
	"fmt"
)

func main() {
	puzzleInput := []int{2, 72, 8949, 0, 981038, 86311, 246, 7636740}

	fmt.Printf("Part one count: %v\n", day11.CountAfterBlinks(puzzleInput, 25))
	fmt.Printf("Part two count: %v\n", day11.CountAfterBlinks(puzzleInput, 75))
}
