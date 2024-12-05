package main

import (
	day5 "advent_of_code/day_5"
	"fmt"
)

func main() {
	rules, rawUpdates := day5.LoadPrintInstructions("./../day5_input.txt")
	rulesMap := day5.ParseRules(rules)
	updates := day5.ParseUpdates(rawUpdates)
	goodUpdateSum, badUpdateSum := day5.SumMiddlePages(rulesMap, updates)
	fmt.Printf("Sum of middle pages of valid updates: %d\n", goodUpdateSum)
	fmt.Printf("Sum of middle pages of invalid updates: %d\n", badUpdateSum)
}
