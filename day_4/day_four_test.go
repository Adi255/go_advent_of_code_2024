package main

import (
	"reflect"
	"testing"
)

func TestDayTwo(t *testing.T) {

	testInput := [][]rune{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}

	t.Run("Can load data", func(t *testing.T) {
		got := LoadGrid("./test_input.txt")

		expected := testInput

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can determine search directions in top left corner", func(t *testing.T) {
		got := DetermineSearchDirections(0, 0, testInput)

		expected := []Direction{South, East, SouthEast}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can determine search directions in top right corner", func(t *testing.T) {
		got := DetermineSearchDirections(0, 9, testInput)

		expected := []Direction{West, South, SouthWest}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can determine search directions in bottom left corner", func(t *testing.T) {
		got := DetermineSearchDirections(9, 0, testInput)

		expected := []Direction{North, East, NorthEast}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can determine search directions in bottom right corner", func(t *testing.T) {
		got := DetermineSearchDirections(9, 9, testInput)

		expected := []Direction{North, West, NorthWest}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can determine search directions in centre", func(t *testing.T) {
		got := DetermineSearchDirections(4, 4, testInput)

		expected := []Direction{North, West, South, East, NorthWest, NorthEast, SouthWest, SouthEast}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can count number of XMAS occurrences", func(t *testing.T) {
		got := CountXmasOccurrences(testInput)

		expected := 18

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

	t.Run("Can count cross-mas occurrences", func(t *testing.T) {
		got := CountCrossMasOccurrences(testInput)

		expected := 9

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

}
