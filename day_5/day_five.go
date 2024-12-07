package day5

import (
	"advent_of_code/util"
	"slices"
	"strconv"
	"strings"
)

func LoadPrintInstructions(path string) (rules, updates []string) {

	lines := util.ReadFileLines(path)

	var isUpdateSection bool
	for _, line := range lines {
		if line == "" {
			isUpdateSection = true
			continue
		}
		if isUpdateSection {
			updates = append(updates, line)
		} else {
			rules = append(rules, line)
		}
	}

	return rules, updates
}

func ParseRules(rules []string) map[int][]int {
	rulesMap := make(map[int][]int)
	splitFunc := func(r rune) bool {
		return r == '|'
	}
	for _, rule := range rules {
		ruleParts := strings.FieldsFunc(rule, splitFunc)
		key, _ := strconv.Atoi(ruleParts[0])
		value, _ := strconv.Atoi(ruleParts[1])
		rulesMap[key] = append(rulesMap[key], value)
	}

	return rulesMap
}

func ParseUpdates(updates []string) [][]int {
	updatesNumbers := [][]int{}
	splitFunc := func(r rune) bool {
		return r == ','
	}
	for _, update := range updates {
		var updateNumbers []int
		tokens := strings.FieldsFunc(update, splitFunc)
		for _, token := range tokens {
			value, _ := strconv.Atoi(token)
			updateNumbers = append(updateNumbers, value)
		}
		updatesNumbers = append(updatesNumbers, updateNumbers)
	}

	return updatesNumbers
}

func SumMiddlePages(rules map[int][]int, updates [][]int) (middlePageSum, badUpdatePageSum int) {
	for _, update := range updates {
		if ValidUpdate(update, rules) {
			middlePageSum += update[len(update)/2]
		} else {
			fixedUpdate := ReorderUpdate(update, rules)
			badUpdatePageSum += fixedUpdate[len(fixedUpdate)/2]
		}
	}
	return middlePageSum, badUpdatePageSum
}

func ValidUpdate(update []int, rules map[int][]int) bool {
	for index, page := range update {
		pageRules := rules[page]
		for _, rulePage := range pageRules {
			if slices.Contains(update[:index], rulePage) {
				return false
			}
		}
	}
	return true
}

func ReorderUpdate(update []int, rules map[int][]int) []int {
	fixedUpdate := make([]int, len(update))
	copy(fixedUpdate, update)

	for index, page := range update {
		pageRules := rules[page]
		for _, rulePage := range pageRules {
			checkedSection := fixedUpdate[:index]
			if !slices.Contains(checkedSection, rulePage) {
				continue
			} else {
				indexOfBadPage := slices.Index(checkedSection, rulePage)
				// Current page can move after each fix is applied
				indexOfCurrentPage := slices.Index(fixedUpdate, page)
				// Swap positions
				fixedUpdate[indexOfBadPage] = page
				fixedUpdate[indexOfCurrentPage] = rulePage
			}
		}
	}

	if ValidUpdate(fixedUpdate, rules) {
		return fixedUpdate
	} else {
		return ReorderUpdate(fixedUpdate, rules)
	}
}
