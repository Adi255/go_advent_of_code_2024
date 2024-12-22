package day19

import (
	"reflect"
	"testing"
)

func TestDayNineteen(t *testing.T) {
	var testTowels = []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	var testCombinations = []string{
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}
	t.Run("Can load towel data", func(t *testing.T) {
		towels, combinations := LoadTowelData("test_input.txt")
		expectedTowels := testTowels
		expectedCombinations := testCombinations
		assertSlicesEqual(t, towels, expectedTowels)
		assertSlicesEqual(t, combinations, expectedCombinations)
	})

	t.Run("Can count possible combinations", func(t *testing.T) {
		possible, totalPossible := CountPossibleCombinations(testTowels, testCombinations)
		expectedPossible, expectedTotalPossible := 6, 16

		if possible != expectedPossible {
			t.Errorf("got %v want %v", possible, expectedPossible)
		}

		if totalPossible != expectedTotalPossible {
			t.Errorf("got %v want %v", totalPossible, expectedTotalPossible)
		}
	})

}

func assertSlicesEqual(t *testing.T, towels []string, expected_towels []string) {
	t.Helper()
	if !reflect.DeepEqual(towels, expected_towels) {
		t.Errorf("got %v want %v", towels, expected_towels)
	}
}
