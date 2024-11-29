package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	recipeArr := []int{3, 7}

	var testDebugFlag = false
	debugFlag = &testDebugFlag

	inputs := []int{9, 5, 18, 2018, 409551}
	expecteds := []string{"5158916779", "0124515891", "9251071085", "5941429882", "1631191756"}

	for i, input := range inputs {
		expected := expecteds[i]

		scoreStr := iterateReceipesRight(recipeArr, []int{0, 1}, input)

		if scoreStr != expected {
			t.Fatalf("'%d' case does not match, expected: '%s', actual: '%s'", input, scoreStr, expected)
		}
	}
}

func TestPart2(t *testing.T) {
	t.Skip("TODO: Finish Part 2")
}
