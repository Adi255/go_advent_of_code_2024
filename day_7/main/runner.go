package main

import (
	day7 "advent_of_code/day_7"
	"fmt"
)

func main() {
	calibrations := day7.LoadCalibrations("../day7_input.txt")
	fmt.Printf("Valid Calibration sum: %d\n", day7.SumValidCalibrations(calibrations, []byte{day7.Plus, day7.Multiply}))
	fmt.Printf("Valid Calibration sum (with concat): %d\n", day7.SumValidCalibrations(calibrations, []byte{day7.Plus, day7.Multiply, day7.Concatenate}))
}
