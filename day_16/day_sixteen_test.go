package day16

import (
	"reflect"
	"testing"
)

func TestDaySixteen(t *testing.T) {
	t.Run("Can load map", func(t *testing.T) {
		got := LoadReindeerMap("./test_input.txt")

		expected := [][]byte{
			{'#', '#', '#', '#', '#', '#'},
			{'#', '.', '.', '.', 'E', '#'},
			{'#', '.', '#', '.', '#', '#'},
			{'#', '.', '#', '.', '.', '#'},
			{'#', 'S', '#', '.', '.', '#'},
			{'#', '#', '#', '#', '#', '#'}}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("Finds lowest scoring path", func(t *testing.T) {

		reindeerMap := LoadReindeerMap("./larger_test_input.txt")

		score, seats := LowestScoringPath(reindeerMap)

		expectedScore := 7036
		expectedSeats := 45

		if score != expectedScore {
			t.Errorf("Expected score of %d, got %d", expectedScore, score)
		}

		if seats != expectedSeats {
			t.Errorf("Expected seats of %d, got %d", expectedSeats, seats)
		}

	})

	t.Run("Finds lowest scoring path in larger example", func(t *testing.T) {

		reindeerMap := LoadReindeerMap("./even_larger_example.txt")

		score, seats := LowestScoringPath(reindeerMap)

		expectedScore := 11048
		expectedSeats := 64

		if score != expectedScore {
			t.Errorf("Expected %d, got %d", expectedScore, score)
		}

		if seats != expectedSeats {
			t.Errorf("Expected seats of %d, got %d", expectedSeats, seats)
		}
	})

}
