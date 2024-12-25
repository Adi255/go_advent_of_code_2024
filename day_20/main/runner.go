package main

import (
	day20 "advent_of_code/day_20"
	"fmt"
)

func main() {
	raceTrack := day20.NavigateMaze("./input.txt")
	cheatCount := day20.CountTimeSavingCheats(raceTrack, 100, 2)
	fmt.Printf("Part 1: time-saving cheats - %d\n", cheatCount)
	partTwoCheatCount := day20.CountTimeSavingCheats(raceTrack, 100, 20)
	fmt.Printf("Part 2: time-saving cheats - %d\n", partTwoCheatCount)
}
