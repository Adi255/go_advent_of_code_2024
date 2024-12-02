package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadReports(path string) [][]int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reportsData [][]int

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		var reportValues []int
		for _, token := range tokens {
			reportValue, _ := strconv.Atoi(token)
			reportValues = append(reportValues, reportValue)
		}
		reportsData = append(reportsData, reportValues)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return reportsData
}

func CountSafeReports(input [][]int) int {
	safeReportCount := 0

	for _, report := range input {
		isSafeReport, _ := isSafe(report)
		if isSafeReport {
			safeReportCount += 1
		}
	}

	return safeReportCount
}

func isSafe(report []int) (bool, [2]int) {
	ascending, firstValue := inspectReport(report)

	previous := firstValue
	for idx, value := range report[1:] {
		if ascending {
			if value-previous > 3 || value-previous < 1 {
				return false, [2]int{idx, idx + 1}
			}
		} else {
			if previous-value > 3 || previous-value < 1 {
				return false, [2]int{idx, idx + 1}
			}
		}
		previous = value
	}

	return true, [2]int{-1, -1}
}

func CountSafeDampenedReports(input [][]int) int {
	safeReportCount := 0

	for _, report := range input {
		isSafeReport, badIndices := isSafe(report)
		if isSafeReport {
			safeReportCount += 1
		} else {
			permutationOne := removeElement(report, badIndices[0])
			permutationTwo := removeElement(report, badIndices[1])
			permutationOneSafe, _ := isSafe(permutationOne)
			permutationTwoSafe, _ := isSafe(permutationTwo)
			if permutationOneSafe || permutationTwoSafe {
				safeReportCount += 1
			}
		}
	}

	return safeReportCount
}

func removeElement(report []int, idx int) []int {
	elementRemoved := append([]int{}, report[:idx]...)
	elementRemoved = append(elementRemoved, report[idx+1:]...)
	return elementRemoved
}

func inspectReport(report []int) (bool, int) {
	var ascending bool
	firstValue := report[0]
	lastValue := report[len(report)-1]
	if lastValue >= firstValue {
		ascending = true
	} else {
		ascending = false
	}
	return ascending, firstValue
}

func main() {
	data := LoadReports("./day2_input.txt")
	fmt.Printf("Count of safe reports: %d\n", CountSafeReports(data))
	fmt.Printf("Count of safe reports after dampening: %d\n", CountSafeDampenedReports(data))
}
