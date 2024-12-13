package day13

import (
	"advent_of_code/util"
	"regexp"
	"strconv"
)

type ClawMachine struct {
	A, B, Target [2]int
}

var costCoefficients = [2]int{3, 1}

func LoadClawMachineData(path string) []ClawMachine {
	lines := util.ReadFileLines(path)
	var clawMachines []ClawMachine
	for i := 0; i+3 <= len(lines); i += 4 {
		clawMachines = append(clawMachines, parseClawMachineData(lines[i:i+3]))
	}
	return clawMachines
}

func parseClawMachineData(lines []string) ClawMachine {
	return ClawMachine{
		extractXYMoves(lines[0]),
		extractXYMoves(lines[1]),
		extractTarget(lines[2])}
}

func extractXYMoves(line string) [2]int {
	movePattern := regexp.MustCompile(`: X\+(\d+), Y\+(\d+)`)
	match := movePattern.FindStringSubmatch(line)
	xMove, _ := strconv.Atoi(match[1])
	yMove, _ := strconv.Atoi(match[2])
	return [2]int{xMove, yMove}
}

func extractTarget(line string) [2]int {
	targetPattern := regexp.MustCompile(`X\=(\d+), Y\=(\d+)`)
	match := targetPattern.FindStringSubmatch(line)
	xMove, _ := strconv.Atoi(match[1])
	yMove, _ := strconv.Atoi(match[2])
	return [2]int{xMove, yMove}
}

func SumAllCosts(machines []ClawMachine, offset int) int {
	var totalCost int
	for _, machine := range machines {
		minCost, found := FindMinCostAfterOffset(machine, offset)
		if found {
			totalCost += minCost
		}
	}
	return totalCost
}

func FindMinCostAfterOffset(machine ClawMachine, offset int) (int, bool) {
	targetX := float64(machine.Target[0] + offset)
	targetY := float64(machine.Target[1] + offset)

	a1, a2, b1, b2 := float64(machine.A[0]), float64(machine.A[1]), float64(machine.B[0]), float64(machine.B[1])

	aPushesFloat := (b2*targetX - targetY*b1) / (b2*a1 - a2*b1)
	bPushes := int((targetY - aPushesFloat*a2) / b2)
	aPushes := int(aPushesFloat)

	if aPushes < 0 || bPushes < 0 {
		return 0, false
	}
	if aPushes*int(a1)+bPushes*int(b1) != int(targetX) || aPushes*int(a2)+bPushes*int(b2) != int(targetY) {
		return 0, false
	}

	return calculateCost([2]int{aPushes, bPushes}), true
}

func calculateCost(pushes [2]int) int {
	return pushes[0]*costCoefficients[0] + pushes[1]*costCoefficients[1]
}
