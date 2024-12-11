package day11

import (
	"strconv"
)

const defaultFactor = 2024

var stoneCountCache = map[[2]int]int{}

func CountAfterBlinks(stones []int, blinks int) int {
	total := 0
	for _, stone := range stones {
		total += blinkStoneNTimes(stone, blinks)
	}
	return total
}

func blinkStoneNTimes(stone int, times int) int {
	if times == 0 {
		return 1
	}
	cached, found := stoneCountCache[[2]int{stone, times}]
	if found {
		return cached
	} else {
		stones := engrave(stone)
		result := 0
		for _, stone := range stones {
			stoneCount := blinkStoneNTimes(stone, times-1)
			result += stoneCount
		}
		stoneCountCache[[2]int{stone, times}] = result
		return result
	}
}

func engrave(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	stoneString := strconv.Itoa(stone)
	stoneStringLength := len(stoneString)
	if stoneStringLength%2 == 0 {
		return splitStoneString(stoneString, stoneStringLength)
	} else {
		return []int{stone * defaultFactor}
	}
}

func splitStoneString(stoneString string, length int) []int {
	halfIndex := length / 2
	firstNumber, _ := strconv.Atoi(stoneString[:halfIndex])
	secondNumber, _ := strconv.Atoi(stoneString[halfIndex:])
	return []int{firstNumber, secondNumber}
}
