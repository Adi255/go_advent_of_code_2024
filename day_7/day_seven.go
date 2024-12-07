package day7

import (
	"advent_of_code/util"
	"strconv"
	"strings"
)

const Plus = '+'
const Multiply = '*'
const Concatenate = '|'

var operatorSequences = map[int][][]byte{}

func splitOnColon(r rune) bool {
	return r == ':'
}

func convertOperandsToInts(operandStrings []string) []int {
	operands := make([]int, len(operandStrings))
	for idx, operandString := range operandStrings {
		operand, _ := strconv.Atoi(operandString)
		operands[idx] = operand
	}
	return operands
}

func LoadCalibrations(path string) []Calibration {

	lines := util.ReadFileLines(path)
	calibrations := make([]Calibration, len(lines))

	for idx, line := range lines {
		tokens := strings.FieldsFunc(line, splitOnColon)
		target, _ := strconv.Atoi(tokens[0])
		operandStrings := strings.Fields(tokens[1])
		calibrations[idx] = Calibration{target, convertOperandsToInts(operandStrings)}
	}

	return calibrations
}

func OperatorCombinations(count int, operators []byte) [][]byte {

	operatorSequence, ok := operatorSequences[count]
	if ok {
		return operatorSequence
	} else {

		operatorCombos := [][]byte{}
		for _, operator := range operators {
			operatorCombos = append(operatorCombos, []byte{operator})
		}

		for i := 1; i < count; i++ {
			var newSlices = [][]byte{}
			for _, operator := range operators {
				newSlice := addOperator(operatorCombos, operator)
				newSlices = append(newSlices, newSlice...)
			}
			operatorCombos = newSlices
		}

		operatorSequences[count] = operatorCombos
		return operatorCombos
	}
}

func addOperator(operators [][]byte, operator byte) [][]byte {
	newOperators := make([][]byte, len(operators))
	for idx, op := range operators {
		opCopy := copySlice(op)
		newOp := append(opCopy, operator)
		newOperators[idx] = newOp
	}
	return newOperators
}

func copySlice(slice []byte) []byte {
	newSlice := make([]byte, len(slice))
	for idx, value := range slice {
		newSlice[idx] = value
	}
	return newSlice
}

func (calibration *Calibration) IsValid(operators []byte) bool {
	operatorCombinations := OperatorCombinations(len(calibration.operands)-1, operators)

	for _, operators := range operatorCombinations {
		total := calibration.operands[0]
		for opIdx, operator := range operators {
			switch operator {
			case Plus:
				total += calibration.operands[opIdx+1]
			case Multiply:
				total *= calibration.operands[opIdx+1]
			case Concatenate:
				total = concatenateNumbers(total, calibration.operands[opIdx+1])
			}
			if total > calibration.target {
				break
			}
		}
		if total == calibration.target {
			return true
		}
	}
	return false
}

func concatenateNumbers(a int, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	concatenated, _ := strconv.Atoi(aStr + bStr)
	return concatenated
}

func SumValidCalibrations(calibrations []Calibration, operators []byte) int64 {
	var validSum int64 = 0
	operatorSequences = map[int][][]byte{}
	for _, calibration := range calibrations {
		if calibration.IsValid(operators) {
			validSum += int64(calibration.target)
		}
	}
	return validSum
}
