package main

import (
	day19 "advent_of_code/day_19"
	"fmt"
)

func main() {
	towels, combinations := day19.LoadTowelData("./input.txt")
	possibleCombinations, totalPossibleCombinations := day19.CountPossibleCombinations(towels, combinations)
	fmt.Printf("Possible towel combinations %d\n", possibleCombinations)
	fmt.Printf("Total possible towel combinations %d\n", totalPossibleCombinations)
}
