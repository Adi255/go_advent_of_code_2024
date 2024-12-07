package day6

import (
	"reflect"
	"testing"
)

var testMap = [][]byte{
	{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
	{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
}

func TestDaySix(t *testing.T) {

	t.Run("Can load data", func(t *testing.T) {
		got := LoadGuardMap("./test_input.txt")
		expected := testMap

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Guard can step forward", func(t *testing.T) {
		input := [][]byte{
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '^', '.'}}

		expected := [][]byte{
			{'.', '.', '.'},
			{'.', '^', '.'},
			{'.', 'X', '.'}}

		got, _ := GuardStep(input, Heading{[2]int{2, 1}, guardUp})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

	})

	t.Run("Guard can turn from up to right", func(t *testing.T) {
		input := [][]byte{
			{'.', '.', '.'},
			{'.', '#', '.'},
			{'.', '^', '.'}}

		expected := [][]byte{
			{'.', '.', '.'},
			{'.', '#', '.'},
			{'.', '>', '.'}}

		got, _ := GuardStep(input, Heading{[2]int{2, 1}, guardUp})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

	})

	t.Run("Guard can turn from right to down", func(t *testing.T) {
		input := [][]byte{
			{'.', '.', '.'},
			{'>', '#', '.'},
			{'.', '.', '.'}}

		expected := [][]byte{
			{'.', '.', '.'},
			{'v', '#', '.'},
			{'.', '.', '.'}}

		got, _ := GuardStep(input, Heading{[2]int{1, 0}, guardRight})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

	})

	t.Run("Guard can turn from down to left", func(t *testing.T) {
		input := [][]byte{
			{'.', 'v', '.'},
			{'.', '#', '.'},
			{'.', '.', '.'}}

		expected := [][]byte{
			{'.', '<', '.'},
			{'.', '#', '.'},
			{'.', '.', '.'}}

		got, _ := GuardStep(input, Heading{[2]int{0, 1}, guardDown})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

	})

	t.Run("Guard can turn from left to up", func(t *testing.T) {
		input := [][]byte{
			{'.', '.', '.'},
			{'.', '#', '<'},
			{'.', '.', '.'}}

		expected := [][]byte{
			{'.', '.', '.'},
			{'.', '#', '^'},
			{'.', '.', '.'}}

		got, _ := GuardStep(input, Heading{[2]int{1, 2}, guardLeft})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

	})

	t.Run("Guard can exit the map upwards", func(t *testing.T) {
		input := [][]byte{
			{'.', '^', '.'},
			{'.', '#', '.'},
			{'.', '.', '.'}}

		expected := [][]byte{
			{'.', 'X', '.'},
			{'.', '#', '.'},
			{'.', '.', '.'}}

		startHeading := Heading{[2]int{0, 1}, guardUp}
		got, endHeading := GuardStep(input, startHeading)

		if !reflect.DeepEqual(startHeading, endHeading) {
			t.Error("Should be finished")
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}

	})

	t.Run("Guard can exit the map downwards", func(t *testing.T) {
		input := [][]byte{
			{'.', '.', '.'},
			{'.', '#', '.'},
			{'.', 'v', '.'}}

		expected := [][]byte{
			{'.', '.', '.'},
			{'.', '#', '.'},
			{'.', 'X', '.'}}

		startHeading := Heading{[2]int{2, 1}, guardDown}
		got, endHeading := GuardStep(input, startHeading)

		if !reflect.DeepEqual(startHeading, endHeading) {
			t.Error("Should be finished")
		}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Guard can take multiple steps", func(t *testing.T) {
		input := [][]byte{
			{'.', '.', '#'},
			{'.', '.', '.'},
			{'.', '#', '<'}}

		expected := [][]byte{
			{'.', '.', '#'},
			{'.', '.', 'X'},
			{'.', '#', 'X'}}

		startHeading := Heading{[2]int{2, 2}, guardLeft}

		stepOne, nextHeading := GuardStep(input, startHeading)
		stepTwo, nextHeading := GuardStep(stepOne, nextHeading)
		stepThree, nextHeading := GuardStep(stepTwo, nextHeading)
		stepFour, finalHeading := GuardStep(stepThree, nextHeading)

		if !reflect.DeepEqual(finalHeading, nextHeading) {
			t.Error("Should be finished")
		}

		if !reflect.DeepEqual(stepFour, expected) {
			t.Errorf("expected %v but got %v", expected, stepFour)
		}

	})

	t.Run("Can count visited positions", func(t *testing.T) {
		input := [][]byte{
			{'.', 'X', '#'},
			{'.', 'X', 'X'},
			{'.', '#', 'X'}}

		got := FindVisitedPositions(input)

		expected := [][2]int{
			{0, 1},
			{1, 1},
			{1, 2},
			{2, 2}}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	})

	t.Run("Can detect closed loop", func(t *testing.T) {
		obstaclePosition := [2]int{7, 6}
		addedObstacleMap := AddObstacleToMap(testMap, obstaclePosition)
		startHeading := Heading{[2]int{6, 4}, guardUp}

		closedLoop := ClosedLoop(addedObstacleMap, startHeading, obstaclePosition)

		if !closedLoop {
			t.Error("Should be a closed loop with provided obstacle")
		}

	})

	t.Run("Can count good obstacle positions", func(t *testing.T) {
		exitedMap := DeepCopyMap(testMap)
		heading := Heading{[2]int{6, 4}, guardUp}
		var nextHeading Heading
		for {
			exitedMap, nextHeading = GuardStep(exitedMap, heading)
			if reflect.DeepEqual(heading, nextHeading) {
				break
			}
			heading = nextHeading
		}
		visitedPositions := FindVisitedPositions(exitedMap)
		startHeading := CurrentGuardHeading(testMap)
		got := CountGoodObstaclePositions(testMap, startHeading, visitedPositions)
		expected := 6

		if got != expected {
			t.Errorf("expected %d possible obstacle positions but got %d", expected, got)
		}
	})
}
