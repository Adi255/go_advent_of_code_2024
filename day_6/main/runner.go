package main

import (
	day6 "advent_of_code/day_6"
	"fmt"
	"slices"
)

func main() {
	guardMap := day6.LoadGuardMap("../day6_input.txt")
	encounteredHeadings, _ := day6.TraverseMap(guardMap)
	visitedPositions := [][2]int{}

	for _, heading := range encounteredHeadings {
		if !slices.Contains(visitedPositions, heading.Position) {
			visitedPositions = append(visitedPositions, heading.Position)
		}
	}

	fmt.Println(len(visitedPositions))

	// Part 2
	countOfGoodObstacles := 0

	_, startPosition := day6.FindGuard(guardMap)

	idxToRemove := slices.Index(visitedPositions, startPosition)
	visitedPositions = append(visitedPositions[:idxToRemove], visitedPositions[idxToRemove+1:]...)

	for _, position := range visitedPositions {
		guardMapVariation := day6.AddObstacleToMap(guardMap, position)
		_, couldExit := day6.TraverseMap(guardMapVariation)
		if !couldExit {
			countOfGoodObstacles++
		}
	}

	fmt.Println(countOfGoodObstacles)

}
