package day5

import (
	"reflect"
	"testing"
)

func TestDayFive(t *testing.T) {

	testRules := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}

	testUpdates := []string{
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}

	t.Run("Can load data", func(t *testing.T) {

		rules, updates := LoadPrintInstructions("./test_input.txt")

		expectedRules := testRules

		expectedUpdates := testUpdates

		if !reflect.DeepEqual(rules, expectedRules) {
			t.Errorf("expected %v but got %v", expectedRules, rules)
		}

		if !reflect.DeepEqual(updates, expectedUpdates) {
			t.Errorf("expected %v but got %v", expectedUpdates, updates)
		}

	})

	t.Run("Can parse rules", func(t *testing.T) {
		got := ParseRules(testRules)

		expected := map[int][]int{
			29: {13},
			47: {53, 13, 61, 29},
			53: {29, 13},
			61: {13, 53, 29},
			75: {29, 53, 47, 61, 13},
			97: {13, 61, 47, 29, 53, 75},
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can parse updates", func(t *testing.T) {
		got := ParseUpdates(testUpdates)

		expected := [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can determine a valid update is valid", func(t *testing.T) {
		rules := ParseRules(testRules)
		update := []int{75, 47, 61, 53, 29}

		got := ValidUpdate(update, rules)

		expected := true

		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("Can determine an invalid update is invalid", func(t *testing.T) {
		rules := ParseRules(testRules)
		update := []int{61, 13, 29}

		got := ValidUpdate(update, rules)

		expected := false

		if got != expected {
			t.Errorf("expected %t but got %t", expected, got)
		}
	})

	t.Run("Can sum middle pages of valid updates", func(t *testing.T) {
		rules := ParseRules(testRules)
		updates := ParseUpdates(testUpdates)

		got, _ := SumMiddlePages(rules, updates)

		expected := 143

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

	t.Run("Can reorder invalid update", func(t *testing.T) {
		invalidUpdate := []int{75, 97, 47, 61, 53}
		rules := ParseRules(testRules)

		got := ReorderUpdate(invalidUpdate, rules)

		expected := []int{97, 75, 47, 61, 53}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can reorder another invalid update", func(t *testing.T) {
		invalidUpdate := []int{97, 13, 75, 29, 47}
		rules := ParseRules(testRules)

		got := ReorderUpdate(invalidUpdate, rules)

		expected := []int{97, 75, 47, 29, 13}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can sum middle pages of invalid updates", func(t *testing.T) {
		rules := ParseRules(testRules)
		updates := ParseUpdates(testUpdates)

		_, badUpdateSum := SumMiddlePages(rules, updates)

		expected := 123

		if badUpdateSum != expected {
			t.Errorf("expected %d but got %d", expected, badUpdateSum)
		}
	})

}
