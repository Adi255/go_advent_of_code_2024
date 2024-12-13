package main

import (
	day13 "advent_of_code/day_13"
	"fmt"
)

func main() {
	machines := day13.LoadClawMachineData("../day13_input.txt")

	minCost := day13.SumAllCosts(machines, 0)

	fmt.Printf("Minimum cost: %v\n", minCost)

	offset := 10000000000000
	fmt.Printf("Minimum cost with offset: %v\n", day13.SumAllCosts(machines, offset))
}
