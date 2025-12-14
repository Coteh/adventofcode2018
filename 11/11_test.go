package main

import "testing"

func TestPowerLevel(t *testing.T) {
	testCases := []struct {
		name         string
		serialNumber int
		x            int
		y            int
		expected     int
	}{
		// TODO: Move these sample test cases to submodule
		{"sample 1", 8, 3, 5, 4},
		{"sample 2", 57, 122, 79, -5},
		{"sample 3", 39, 217, 196, 0},
		{"sample 4", 71, 101, 153, 4},
	}

	for _, testCase := range testCases {
		powerLevel := getPowerLevel(testCase.x, testCase.y, testCase.serialNumber)
		if powerLevel != testCase.expected {
			t.Fatalf("case \"%s\" (serial number: %d) does not match, expected: %d, actual: %d", testCase.name, testCase.serialNumber, testCase.expected, powerLevel)
		}
	}
}

func TestPart1(t *testing.T) {
	testCases := []struct {
		name         string
		serialNumber int
		expectedX    int
		expectedY    int
	}{
		// TODO: Move these sample test cases to submodule
		{"sample 5", 18, 33, 45},
		{"sample 6", 42, 21, 61},
		// TODO: Remove this test case and update test.sh to test using data file from submodule
		{"input", 7672, 22, 18},
	}

	for _, testCase := range testCases {
		hologram := createHologramFromSerialNumber(testCase.serialNumber)
		actualX, actualY := findLargest3x3Square(hologram, false)
		if actualX != testCase.expectedX || actualY != testCase.expectedY {
			t.Fatalf("case \"%s\" (serial number: %d) does not match, expected: (%d,%d), actual: (%d,%d)", testCase.name, testCase.serialNumber, testCase.expectedX, testCase.expectedY, actualX, actualY)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name         string
		serialNumber int
		expectedX    int
		expectedY    int
		expectedSize int
	}{
		// TODO: Move these sample test cases to submodule
		{"sample 5", 18, 90, 269, 16},
		{"sample 6", 42, 232, 251, 12},
		// TODO: Remove this test case and update test.sh to test using data file from submodule
		{"input", 7672, 234, 197, 14},
	}

	for _, testCase := range testCases {
		hologram := createHologramFromSerialNumber(testCase.serialNumber)
		actualX, actualY, actualSize := findLargestSquare(hologram, false)
		if actualX != testCase.expectedX || actualY != testCase.expectedY || actualSize != testCase.expectedSize {
			t.Fatalf("case \"%s\" (serial number: %d) does not match, expected: (%d,%d,%d), actual: (%d,%d,%d)", testCase.name, testCase.serialNumber, testCase.expectedX, testCase.expectedY, testCase.expectedSize, actualX, actualY, actualSize)
		}
	}
}
