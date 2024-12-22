package main

import (
	day17 "advent_of_code/day_17"
	"fmt"
	"strconv"
	"strings"
)

const testProgram = "0,1,5,4,3,0"
const realProgram = "2,4,1,5,7,5,4,3,1,6,0,3,5,5,3,0"
const testA = 729
const realA = 61156655

func main() {
	program := realProgram
	programStringArray := strings.Split(program, ",")
	programIntArray := make([]int, len(programStringArray))
	for i, v := range programStringArray {
		programIntArray[i], _ = strconv.Atoi(v)
	}

	initialState := day17.State{
		A:       realA,
		Program: programIntArray,
	}

	output := day17.ExecuteInstructions(initialState)
	fmt.Printf("Part 1 Output: %s\n", output)

	targetRegister := day17.FindTargetRegister(program)
	fmt.Printf("Part 2 Target Register: %d\n", targetRegister)
}
