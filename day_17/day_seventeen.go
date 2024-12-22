package day17

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func ExecuteInstructions(initialState State) string {
	state := initialState
	endIndex := len(state.Program) - 1
	for state.Pointer <= endIndex {
		instruction := state.Program[state.Pointer]
		operand := state.Program[state.Pointer+1]
		state = InstructionsMap[instruction](operand, state)
	}

	output := state.output
	return strings.Join(output, ",")
}

func FindTargetRegister(program string) int64 {
	programStringArray := strings.Split(program, ",")
	fmt.Printf("Program: %v\n", programStringArray)
	programIntArray := make([]int, len(programStringArray))
	endIndex := len(programStringArray) - 1
	for i, v := range programStringArray {
		programIntArray[i], _ = strconv.Atoi(v)
	}

	trialValue := int64(math.Pow(8, float64(len(programStringArray))))
	trialValue = 844424930131968
	for {
		initialState := State{
			A:       trialValue,
			Program: programIntArray,
		}

		state := initialState
		for state.Pointer <= endIndex {
			instruction := state.Program[state.Pointer]
			operand := state.Program[state.Pointer+1]
			state = InstructionsMap[instruction](operand, state)
		}

		fmt.Printf("Trial value: %d, Output: %v\n", trialValue, state.output)
		if slices.Equal(programStringArray, state.output) {
			return trialValue
		} else {
			trialValue += 1000000
		}
	}

}
