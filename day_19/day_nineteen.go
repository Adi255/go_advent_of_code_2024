package day19

import (
	"advent_of_code/util"
	"strings"
)

func splitOnComma(r rune) bool {
	return r == ','
}

var possibleCombinations = map[string]int{}

func LoadTowelData(path string) (towels []string, combinations []string) {
	lines := util.ReadFileLines(path)
	towels = strings.FieldsFunc(lines[0], splitOnComma)
	towels = trimEachTowel(towels)
	combinations = lines[2:]
	return
}

func trimEachTowel(towels []string) []string {
	for i, towel := range towels {
		towels[i] = strings.TrimSpace(towel)
	}
	return towels
}

func CountPossibleCombinations(towels []string, combinations []string) (int, int) {
	var possibleCombinations int
	totalPossibleCombinations := 0
	for _, combination := range combinations {
		if isPossibleCombination(towels, combination) {
			possibleCombinations++
		}
		totalPossibleCombinations += countPossibleCombination(towels, combination)
	}
	return possibleCombinations, totalPossibleCombinations
}

func isPossibleCombination(towels []string, combination string) bool {
	if combination == "" {
		return true
	}
	for _, towel := range towels {
		towelLength := len(towel)
		if len(combination) >= towelLength && combination[:towelLength] == towel {
			if isPossibleCombination(towels, combination[len(towel):]) {
				return true
			}
		}
	}
	return false
}

func countPossibleCombination(towels []string, combination string) int {
	if possible, cached := possibleCombinations[combination]; cached {
		return possible
	}
	if combination == "" {
		return 1
	}
	totalCombinations := 0
	for _, towel := range towels {
		towelLength := len(towel)
		if len(combination) >= towelLength && combination[:towelLength] == towel {
			totalCombinations += countPossibleCombination(towels, combination[towelLength:])
			possibleCombinations[combination] = totalCombinations
		}
	}
	return totalCombinations
}
