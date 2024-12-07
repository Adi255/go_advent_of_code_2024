package main

import (
	day6 "advent_of_code/day_6"
	"fmt"
	"reflect"
)

func main2() {
	guardMap := day6.LoadGuardMap("../day6_input.txt")
	updatedMap := day6.DeepCopyMap(guardMap)
	guardHeading := day6.CurrentGuardHeading(guardMap)
	var newHeading day6.Heading
	prevHeading := guardHeading
	for {
		updatedMap, newHeading = day6.GuardStep(updatedMap, prevHeading)
		if reflect.DeepEqual(newHeading, prevHeading) {
			break
		}
		prevHeading = newHeading
	}

	visitedPositions := day6.FindVisitedPositions(updatedMap)
	fmt.Printf("Visited positions: %d\n", len(visitedPositions))

	potentialObstacleCount := len(day6.PotentialObstaclePositions)

	// for i, pos := range visitedPositions {
	// 	if pos == guardStartingPosition {
	// 		visitedPositions = append(visitedPositions[:i], visitedPositions[i+1:]...)
	// 		break
	// 	}
	// }

	fmt.Printf("Good obstacle positions: %d", potentialObstacleCount)
}
