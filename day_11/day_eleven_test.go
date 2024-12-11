package day11

import (
	"testing"
)

func TestDayEight(t *testing.T) {

	t.Run("Can blink once and see correct result", func(t *testing.T) {
		stones := []int{0, 1, 10, 99, 999}
		got := CountAfterBlinks(stones, 1)

		expected := 7

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})

	t.Run("Can blink multiple times and see correct result", func(t *testing.T) {
		stones := []int{125, 17}
		got := CountAfterBlinks(stones, 6)

		expected := 22

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})
}
