package main

import (
	"testing"
)

func TestExpectedTimeRequiredForSteps(t *testing.T) {
	fixedTimeAmount := 60

	for i := 0; i < 26; i++ {
		expectedTime := (fixedTimeAmount + 1 + i)
		actualTime := calculateTimeRequired(rune('A'+i), fixedTimeAmount)
		if actualTime != expectedTime {
			t.Fatalf("Amount of work required for step %c is incorrect. expected: %d, actual: %d", rune('A'+i), expectedTime, actualTime)
		}
	}
}
