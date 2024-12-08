package main

import (
	day8 "advent_of_code/day_8"
	"fmt"
)

func main() {
	antennaMap := day8.LoadAntennaMap("../day8_input.txt")
	uniqueAntinodePositions := day8.CountUniqueAntinodes(antennaMap)

	fmt.Printf("Unique antinode positions: %d\n", uniqueAntinodePositions)

	uniqueHarmonicAntinodePositions := day8.CountUniqueHarmonicAntinodes(antennaMap)
	fmt.Printf("Unique harmonic antinode positions: %d\n", uniqueHarmonicAntinodePositions)
}
