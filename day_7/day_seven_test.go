package day7

import (
	"reflect"
	"testing"
)

func TestDaySeven(t *testing.T) {

	var testCalibrations = []Calibration{
		{190, []int{10, 19}},
		{3267, []int{81, 40, 27}},
		{83, []int{17, 5}},
		{156, []int{15, 6}},
		{7290, []int{6, 8, 6, 15}},
		{161011, []int{16, 10, 13}},
		{192, []int{17, 8, 14}},
		{21037, []int{9, 7, 18, 13}},
		{292, []int{11, 6, 16, 20}}}

	var plusAndMultiply = []byte{Plus, Multiply}
	var allOperators = []byte{Plus, Multiply, Concatenate}

	t.Run("Can load calibrations", func(t *testing.T) {

		got := LoadCalibrations("test_input.txt")

		expected := testCalibrations

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	})

	t.Run("Can come up with all possible operator combinations", func(t *testing.T) {
		got := OperatorCombinations(3, []byte{Plus, Multiply})

		expected := [][]byte{
			{'+', '+', '+'},
			{'*', '+', '+'},
			{'+', '*', '+'},
			{'*', '*', '+'},
			{'+', '+', '*'},
			{'*', '+', '*'},
			{'+', '*', '*'},
			{'*', '*', '*'},
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v but got %v", expected, got)
		}

	})

	t.Run("Can come up with all possible operator combinations for 4 operations", func(t *testing.T) {
		gotLength := len(OperatorCombinations(4, []byte{Plus, Multiply}))

		expectedLength := 16

		if gotLength != expectedLength {
			t.Errorf("Expected length of %d but got %d", expectedLength, gotLength)
		}
	})

	t.Run("Can detect valid calibration", func(t *testing.T) {
		calibration := Calibration{3267, []int{81, 40, 27}}
		got := calibration.IsValid(plusAndMultiply)

		expected := true

		if got != expected {
			t.Errorf("Expected %t but got %t", expected, got)
		}
	})

	t.Run("Can detect invalid calibration", func(t *testing.T) {
		calibration := Calibration{3267, []int{81, 40, 28}}
		got := calibration.IsValid(plusAndMultiply)

		expected := false

		if got != expected {
			t.Errorf("Expected %t but got %t", expected, got)
		}
	})

	t.Run("Can sum number of valid calibrations", func(t *testing.T) {
		got := SumValidCalibrations(testCalibrations, plusAndMultiply)

		expected := int64(3749)

		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})

	t.Run("Can sum number of valid calibrations with concat", func(t *testing.T) {
		got := SumValidCalibrations(testCalibrations, allOperators)

		expected := int64(11387)

		if got != expected {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})

	t.Run("Can come up with operator combinations for 3 operations, including |", func(t *testing.T) {
		gotLength := len(OperatorCombinations(3, []byte{'+', '*', '|'}))

		expectedLength := 27

		if gotLength != expectedLength {
			t.Errorf("Expected length of %d but got %d", expectedLength, gotLength)
		}
	})

	t.Run("Can detect valid concatenation calibration", func(t *testing.T) {
		calibration := Calibration{1010, []int{10, 10}}
		got := calibration.IsValid(allOperators)

		expected := true

		if got != expected {
			t.Errorf("Expected %t but got %t", expected, got)
		}
	})

}
