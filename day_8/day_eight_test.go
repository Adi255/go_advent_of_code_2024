package day8

import (
	"reflect"
	"slices"
	"testing"
)

func TestDayEight(t *testing.T) {

	var testMap = [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '0', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '0', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '0', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '0', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', 'A', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'A', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	var testAntennaGroup = AntennaGroup{
		'0': []Position{{1, 8}, {2, 5}, {3, 7}, {4, 4}},
		'A': []Position{{5, 6}, {8, 8}, {9, 9}}}

	t.Run("Load antenna map from file", func(t *testing.T) {
		got := LoadAntennaMap("test_input.txt")

		expected := testMap

		if !reflect.DeepEqual(got, expected) {
			println("Expected:")
			PrintAntennaMap(expected)
			println("Got:")
			PrintAntennaMap(got)
			t.Error("Map not loaded correctly")
		}
	})

	t.Run("Can locate antennas", func(t *testing.T) {
		input := testMap

		got := LocateAntennas(input)

		expected := testAntennaGroup

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v but expected %v", got, expected)
		}
	})

	t.Run("Can calculate antinode position for antenna group", func(t *testing.T) {
		input := AntennaGroup{
			'0': []Position{{1, 8}, {2, 5}}}

		got := LocateAntinodes(input, testMap)

		expected := []Position{{0, 11}, {3, 2}}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v but expected %v", got, expected)
		}

	})

	t.Run("Can count unique antinode positions", func(t *testing.T) {
		got := CountUniqueAntinodes(testMap)

		expected := 14

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})

	t.Run("Can calculate unique antinode position for antenna group, with harmonics", func(t *testing.T) {
		input := AntennaGroup{'T': []Position{{0, 0}, {1, 3}, {2, 1}}}

		got := LocateHarmonicAntinodes(input, testMap)

		expected := []Position{
			{0, 0}, {1, 3}, {2, 6}, {3, 9}, {0, 5}, {6, 3}, {8, 4}, {10, 5}, {2, 1}, {4, 2}}

		if len(got) != len(expected) {
			t.Errorf("got length of %d but expected %d", len(got), len(expected))
		}

		for _, position := range got {
			if !slices.Contains(expected, position) {
				t.Errorf("missing position %v", position)
			}
		}

	})
}

func PrintAntennaMap(antennaMap [][]byte) {
	for _, row := range antennaMap {
		for _, cell := range row {
			print(string(cell))
		}
		println()
	}
}
