package day13

import (
	"reflect"
	"testing"
)

func TestDayThirteen(t *testing.T) {

	t.Run("Can load machine data", func(t *testing.T) {

		got := LoadClawMachineData("./test_input.txt")

		expected := []ClawMachine{
			{[2]int{94, 34}, [2]int{22, 67}, [2]int{8400, 5400}},
			{[2]int{26, 66}, [2]int{67, 21}, [2]int{12748, 12176}},
			{[2]int{17, 86}, [2]int{84, 37}, [2]int{7870, 6450}},
			{[2]int{69, 23}, [2]int{27, 71}, [2]int{18641, 10279}}}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	})

	t.Run("Can find minimum cost for machine", func(t *testing.T) {
		machine := ClawMachine{
			[2]int{94, 34}, [2]int{22, 67}, [2]int{8400, 5400}}

		actualCost, found := FindMinCostAfterOffset(machine, 0)

		if !found {
			t.Error("Expected to find a minimum cost")
		}
		expectedCost := 280
		if actualCost != expectedCost {
			t.Errorf("Got cost of %v, expected %v", actualCost, expectedCost)
		}
	})

	t.Run("Can determine when no solution exists", func(t *testing.T) {
		machine := ClawMachine{
			[2]int{26, 66}, [2]int{67, 21}, [2]int{12748, 12176}}

		_, found := FindMinCostAfterOffset(machine, 0)

		if found {
			t.Error("Didn't expect to find a minimum cost")
		}
	})

	t.Run("Can sum all min costs for machines", func(t *testing.T) {
		machines := []ClawMachine{
			{[2]int{94, 34}, [2]int{22, 67}, [2]int{8400, 5400}},
			{[2]int{26, 66}, [2]int{67, 21}, [2]int{12748, 12176}},
			{[2]int{17, 86}, [2]int{84, 37}, [2]int{7870, 6450}},
			{[2]int{69, 23}, [2]int{27, 71}, [2]int{18641, 10279}}}

		got := SumAllCosts(machines, 0)

		expected := 480

		if got != expected {
			t.Errorf("Got %d total cost, expected %d", got, expected)
		}
	})

}

func TestDayThirteenPart2(t *testing.T) {
	t.Run("Can find minimum cost for machine", func(t *testing.T) {
		machine := ClawMachine{[2]int{26, 66}, [2]int{67, 21}, [2]int{12748, 12176}}

		offset := 10000000000000
		actualCost, found := FindMinCostAfterOffset(machine, offset)

		if !found {
			t.Error("Expected to find a minimum cost")
		}

		expectedCost := 459236326669
		if actualCost != expectedCost {
			t.Errorf("Got cost of %v, expected %v", actualCost, expectedCost)
		}
	})
}
