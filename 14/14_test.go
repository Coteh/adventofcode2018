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
			t.Fatalf("'%d' case does not match, expected: '%s', actual: '%s'", input, expected, scoreStr)
		}
	}
}

func TestPart2(t *testing.T) {
	recipeArr := []int{3, 7}

	var testDebugFlag = false
	debugFlag = &testDebugFlag

	type TestCase struct {
		name string
		input string
		expected int
	}

	testCases := []TestCase{
		{"sample 1", "51589", 9},
		{"sample 2", "01245", 5},
		{"sample 3", "92510", 18},
		{"sample 4", "59414", 2018},
		{"https://www.reddit.com/r/adventofcode/comments/a671s8/2018_day_14_part_2_i_dont_know_why_my_answer_is/", "147061", 20283721},
		{"input", "409551", 20219475},
		{"last digit of input is the second digit of recipe sum", "5158916", 9},
	}

	for _, testCase := range testCases {
		numRecipes := iterateReceipesLeft(recipeArr, []int{0, 1}, testCase.input)

		if numRecipes != testCase.expected {
			t.Fatalf("case \"%s\" ('%s') does not match, expected: %d, actual: %d", testCase.name, testCase.input, testCase.expected, numRecipes)
		}
	}
}
