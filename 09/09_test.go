package main

import "testing"

func TestPart1(t *testing.T) {
	testCases := []struct {
		name       string
		numPlayers int
		lastPoint  int
		expected   int
	}{
		// TODO: Move these sample test cases to submodule
		{"sample 1", 9, 25, 32},
		{"sample 2", 10, 1618, 8317},
		{"sample 3", 13, 7999, 146373},
		{"sample 4", 17, 1104, 2764},
		{"sample 5", 21, 6111, 54718},
		{"sample 6", 30, 5807, 37305},
		// TODO: Remove this test case and update test.sh to test using data file from submodule
		{"input", 418, 70769, 402398},
	}

	for _, testCase := range testCases {
		actual := beginGame(testCase.numPlayers, testCase.lastPoint, false)
		if actual != testCase.expected {
			t.Fatalf("case \"%s\" (numPlayers: %d, lastPoint: %d) does not match, expected: %d, actual: %d", testCase.name, testCase.numPlayers, testCase.lastPoint, testCase.expected, actual)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name       string
		numPlayers int
		lastPoint  int
		expected   int
	}{
		// TODO: Remove this test case and update test.sh to test using data file from submodule
		{"input", 418, 70769 * 100, 3426843186},
	}

	for _, testCase := range testCases {
		actual := beginGame(testCase.numPlayers, testCase.lastPoint, false)
		if actual != testCase.expected {
			t.Fatalf("case \"%s\" (numPlayers: %d, lastPoint: %d) does not match, expected: %d, actual: %d", testCase.name, testCase.numPlayers, testCase.lastPoint, testCase.expected, actual)
		}
	}
}
