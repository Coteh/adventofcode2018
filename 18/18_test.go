package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type RulesExpectation struct {
	grid         Grid
	x            int
	y            int
	ruleName     string
	expectedAcre Acre
}

func Test_AcreRules(t *testing.T) {
	rules := []RulesExpectation{
		{
			grid: Grid{
				[]Acre{'.', '.', '.'},
				[]Acre{'.', '.', '.'},
				[]Acre{'|', '|', '|'},
			},
			x:            1,
			y:            1,
			ruleName:     "open acre becomes trees if 3 or more adjacent trees (case 1)",
			expectedAcre: '|',
		},
		{
			grid: Grid{
				[]Acre{'.', '.', '.'},
				[]Acre{'.', '.', '.'},
				[]Acre{'.', '|', '|'},
			},
			x:            1,
			y:            1,
			ruleName:     "open acre remains the same if less than 3 adjacent trees (case 1)",
			expectedAcre: '.',
		},
		{
			grid: Grid{
				[]Acre{'.', '.', '.'},
				[]Acre{'.', '|', '.'},
				[]Acre{'#', '#', '#'},
			},
			x:            1,
			y:            1,
			ruleName:     "trees becomes lumberyard if 3 or more adjacent lumberyards (case 1)",
			expectedAcre: '#',
		},
		{
			grid: Grid{
				[]Acre{'.', '.', '.'},
				[]Acre{'.', '|', '.'},
				[]Acre{'#', '#', '#'},
			},
			x:            1,
			y:            1,
			ruleName:     "trees remains the same if less than 3 adjacent lumberyards (case 1)",
			expectedAcre: '#',
		},
		{
			grid: Grid{
				[]Acre{'.', '.', '.'},
				[]Acre{'.', '#', '.'},
				[]Acre{'#', '|', '.'},
			},
			x:            1,
			y:            1,
			ruleName:     "lumberyard remains a lumberyard if at least 1 adjacent lumberyard and at least 1 adjacent tree (case 1)",
			expectedAcre: '#',
		},
		{
			grid: Grid{
				[]Acre{'#', '#', '|'},
				[]Acre{'.', '#', '.'},
				[]Acre{'.', '.', '#'},
			},
			x:            1,
			y:            1,
			ruleName:     "lumberyard remains a lumberyard if at least 1 adjacent lumberyard and at least 1 adjacent tree (case 2)",
			expectedAcre: '#',
		},
		{
			grid: Grid{
				[]Acre{'.', '.', '.'},
				[]Acre{'.', '#', '.'},
				[]Acre{'.', '.', '.'},
			},
			x:            1,
			y:            1,
			ruleName:     "lumberyard becomes open acre if no adjacent lumberyards or no adjacent trees (case 1)",
			expectedAcre: '.',
		},
		{
			grid: Grid{
				[]Acre{'#', '.'},
				[]Acre{'#', '|'},
			},
			x:            0,
			y:            0,
			ruleName:     "handles top left corner boundary case",
			expectedAcre: '#',
		},
		{
			grid: Grid{
				[]Acre{'.', '|'},
				[]Acre{'#', '#'},
			},
			x:            1,
			y:            1,
			ruleName:     "handles bottom right corner boundary case",
			expectedAcre: '#',
		},
		{
			grid: Grid{
				[]Acre{'.', '#'},
				[]Acre{'#', '|'},
			},
			x:            1,
			y:            0,
			ruleName:     "handles top right corner boundary case",
			expectedAcre: '#',
		},
		{
			grid: Grid{
				[]Acre{'#', '.'},
				[]Acre{'#', '|'},
			},
			x:            0,
			y:            1,
			ruleName:     "handles bottom left corner boundary case",
			expectedAcre: '#',
		},
	}

	for _, rule := range rules {
		newAcre, err := resolveAcre(rule.grid, rule.x, rule.y)
		require.NoError(t, err)
		assert.Equal(t, rule.expectedAcre, newAcre, fmt.Sprintf("test case failed: '%s': new acre '%c' != expected acre '%c'", rule.ruleName, newAcre, rule.expectedAcre))
	}
}
