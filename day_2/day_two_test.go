package main

import (
	"reflect"
	"testing"
)

func TestDayTwo(t *testing.T) {

	t.Run("Can load data", func(t *testing.T) {
		got := LoadReports("./test_data.txt")
		expected := [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9}}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can calculate number of safe reports", func(t *testing.T) {
		input := [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9}}

		got := CountSafeReports(input)
		expected := 2

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

	t.Run("Can calculate number of safe reports with dampener", func(t *testing.T) {
		input := [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9}}

		got := CountSafeDampenedReports(input)
		expected := 4

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

	t.Run("Can create permutations from slice with missing elements", func(t *testing.T) {
		input := []int{61, 64, 67, 64, 67, 68, 72}

		expected := [][]int{{61, 64, 67, 64, 67, 72}, {61, 64, 67, 64, 67, 68}}

		idx1, idx2 := 5, 6

		first_actual := removeElement(input, idx1)
		second_actual := removeElement(input, idx2)

		if !reflect.DeepEqual(first_actual, expected[0]) {
			t.Errorf("expected %v but got %v", expected[0], first_actual)
		}

		if !reflect.DeepEqual(second_actual, expected[1]) {
			t.Errorf("expected %v but got %v", expected[1], second_actual)
		}
	})

	t.Run("Can identify problem indices", func(t *testing.T) {
		input := []int{40, 41, 46, 47, 49, 55}

		_, badIndices := isSafe(input)
		expected := [2]int{1, 2}
		if !reflect.DeepEqual(badIndices, expected) {
			t.Errorf("expected %v but got %v", expected, badIndices)
		}

		first_actual := removeElement(input, expected[0])
		second_actual := removeElement(input, expected[1])

		expected_permutations := [][]int{{40, 46, 47, 49, 55}, {40, 41, 47, 49, 55}}

		if !reflect.DeepEqual(first_actual, expected_permutations[0]) {
			t.Errorf("expected %v but got %v", expected_permutations[0], first_actual)
		}

		if !reflect.DeepEqual(second_actual, expected_permutations[1]) {
			t.Errorf("expected %v but got %v", expected_permutations[1], second_actual)
		}
	})
}
