package main

import (
	day16 "advent_of_code/day_16"
	"fmt"
)

func main() {
	reindeerMap := day16.LoadReindeerMap("input.txt")
	lowestScore, bestSeats := day16.LowestScoringPath(reindeerMap)
	fmt.Printf("Part 1 lowest score: %d\n", lowestScore)
	fmt.Printf("Part 2 best seats: %d\n", bestSeats)
}
