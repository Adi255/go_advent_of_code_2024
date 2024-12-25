package day20

import (
	"testing"
)

func TestDayTwenty(t *testing.T) {

	var testMaze = NavigateMaze("test_input.txt")

	t.Run("Can navigate through the maze", func(t *testing.T) {
		input := "test_input.txt"
		actualSolution := NavigateMaze(input)
		expected := 85
		if len(actualSolution) != expected {
			t.Errorf("got %d want %d", len(actualSolution), expected)
		}
	})

	t.Run("Can find 2ps saving 2 space cheats", func(t *testing.T) {
		thresholdPs := 2
		actualCheatCount := CountTimeSavingCheats(testMaze, thresholdPs, 2)
		expected := 44
		if actualCheatCount != expected {
			t.Errorf("got %d want %d", actualCheatCount, expected)
		}
	})

	t.Run("Can find 50ps saving 10 space cheats", func(t *testing.T) {
		thresholdPs := 50
		actualCheatCount := CountTimeSavingCheats(testMaze, thresholdPs, 20)
		expected := 44
		if actualCheatCount != expected {
			t.Errorf("got %d want %d", actualCheatCount, expected)
		}
	})

}
