package day17

import (
	"math"
	"strconv"
)

type State struct {
	A, B, C int64
	Program []int
	Pointer int
	output  []string
}

func combo(operand int, state State) int64 {
	if 0 <= operand && operand <= 3 {
		return int64(operand)
	}
	switch operand {
	case 4:
		return state.A
	case 5:
		return state.B
	case 6:
		return state.C
	default:
		panic("Invalid operand")
	}
}

var InstructionsMap = map[int]func(int, State) State{
	0: adv,
	1: bxl,
	2: bst,
	3: jnz,
	4: bxc,
	5: out,
	6: bdv,
	7: cdv,
}

func performDivision(operand int, state State) int64 {
	numerator := state.A
	comboOperand := combo(operand, state)
	denominator := math.Pow(2.0, float64(comboOperand))
	return numerator / int64(denominator)
}

var adv = func(operand int, state State) State {
	state.A = performDivision(operand, state)
	state.Pointer += 2
	return state
}

var bxl = func(operand int, state State) State {
	state.B = state.B ^ int64(operand)
	state.Pointer += 2
	return state
}

var bst = func(operand int, state State) State {
	comboOperand := combo(operand, state)
	state.B = comboOperand % 8
	state.Pointer += 2
	return state
}

var jnz = func(operand int, state State) State {
	if state.A != 0 {
		state.Pointer = int(operand)
	} else {
		state.Pointer += 2
	}
	return state
}

var bxc = func(_ int, state State) State {
	state.B = state.B ^ state.C
	state.Pointer += 2
	return state
}

var out = func(operand int, state State) State {
	comboOperand := combo(operand, state)
	value := comboOperand % 8
	state.output = append(state.output, strconv.Itoa(int(value)))
	state.Pointer += 2
	return state
}

var bdv = func(operand int, state State) State {
	state.B = performDivision(operand, state)
	state.Pointer += 2
	return state
}

var cdv = func(operand int, state State) State {
	state.C = performDivision(operand, state)
	state.Pointer += 2
	return state
}
