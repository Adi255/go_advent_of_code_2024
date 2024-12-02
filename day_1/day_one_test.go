package main

import (
	"reflect"
	"testing"
)

func TestDayOne(t *testing.T) {

	t.Run("Can load data from a file", func(t *testing.T) {

		input_one, input_two := LoadData("test_input.txt")
		expected_one := []int{1, 4, 6}
		expected_two := []int{3, 2, 5}

		if !reflect.DeepEqual(input_one, expected_one) {
			t.Errorf("got %v want %v", input_one, expected_one)
		}

		if !reflect.DeepEqual(input_two, expected_two) {
			t.Errorf("got %v want %v", input_one, expected_one)
		}
	})

	t.Run("Can sum distance in two lists", func(t *testing.T) {
		input_one := []int{1, 3, 5}
		input_two := []int{4, 2, 6}

		got := FindDistance(input_one, input_two)
		expected := 3

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

	t.Run("Can calculate similarity score", func(t *testing.T) {
		input_one := []int{3, 4, 2, 1, 3, 3}
		input_two := []int{4, 3, 5, 3, 9, 3}

		got := FindSimilarity(input_one, input_two)
		expected := 31

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})
}
