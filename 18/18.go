package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const DEBUG_ON bool = false

type Acre rune

const (
	OPEN_GROUND Acre = '.'
	TREES       Acre = '|'
	LUMBERYARD  Acre = '#'
)

type Grid [][]Acre

func parseInput() (grid Grid) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		runeArr := []Acre(line)

		grid = append(grid, runeArr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func resolveAcre(grid Grid, xPos, yPos int) (Acre, error) {
	if len(grid) == 0 {
		return 0, fmt.Errorf("grid has no rows")
	}

	if yPos < 0 || yPos >= len(grid) || xPos < 0 || xPos >= len(grid[0]) {
		return 0, fmt.Errorf("invalid position %d,%d", xPos, yPos)
	}

	currAcre := grid[yPos][xPos]
	newAcre := currAcre

	adjacent := make([]Acre, 0, 8)
	if yPos-1 >= 0 {
		if xPos-1 >= 0 {
			adjacent = append(adjacent, grid[yPos-1][xPos-1])
		}
		adjacent = append(adjacent, grid[yPos-1][xPos])
		if xPos+1 < len(grid[0]) {
			adjacent = append(adjacent, grid[yPos-1][xPos+1])
		}
	}
	if xPos-1 >= 0 {
		adjacent = append(adjacent, grid[yPos][xPos-1])
	}
	if xPos+1 < len(grid[0]) {
		adjacent = append(adjacent, grid[yPos][xPos+1])
	}
	if yPos+1 < len(grid[0]) {
		if xPos-1 >= 0 {
			adjacent = append(adjacent, grid[yPos+1][xPos-1])
		}
		adjacent = append(adjacent, grid[yPos+1][xPos])
		if xPos+1 < len(grid[0]) {
			adjacent = append(adjacent, grid[yPos+1][xPos+1])
		}
	}

	switch currAcre {
	case OPEN_GROUND:
		numAdjacentTrees := 0
		for _, adj := range adjacent {
			if adj == TREES {
				numAdjacentTrees++
			}
		}
		if numAdjacentTrees >= 3 {
			newAcre = TREES
		}
	case TREES:
		numAdjacentLumberyards := 0
		for _, adj := range adjacent {
			if adj == LUMBERYARD {
				numAdjacentLumberyards++
			}
		}
		if numAdjacentLumberyards >= 3 {
			newAcre = LUMBERYARD
		}
	case LUMBERYARD:
		numAdjacentTrees := 0
		numAdjacentLumberyards := 0
		for _, adj := range adjacent {
			if adj == TREES {
				numAdjacentTrees++
			}
		}
		for _, adj := range adjacent {
			if adj == LUMBERYARD {
				numAdjacentLumberyards++
			}
		}
		if numAdjacentTrees < 1 || numAdjacentLumberyards < 1 {
			newAcre = OPEN_GROUND
		}
	}

	return newAcre, nil
}

func processGrid(grid Grid) (newGrid Grid) {
	newGrid = deepCopyGrid(grid)

	for y, row := range grid {
		for x, _ := range row {
			newAcre, err := resolveAcre(grid, x, y)
			if err != nil {
				panic(err)
			}
			newGrid[y][x] = newAcre
		}
	}

	return newGrid
}

func performGridGenerations(grid Grid, numGenerations int) (answer int) {
	currGrid := grid
	printGrid(currGrid, 0)
	for i := 0; i < numGenerations; i++ {
		currGrid = processGrid(currGrid)
		printGrid(currGrid, i+1)
	}

	var numTrees, numLumberyards int
	for _, row := range currGrid {
		for _, cell := range row {
			switch cell {
			case TREES:
				numTrees++
			case LUMBERYARD:
				numLumberyards++
			}
		}
	}
	return numTrees * numLumberyards
}

func deepCopyGrid(grid Grid) (newGrid Grid) {
	newGrid = make(Grid, 0, len(grid))
	for _, row := range grid {
		copiedRow := make([]Acre, len(row))
		numCopied := copy(copiedRow, row)
		if numCopied < len(row) {
			panic("copy did not succeed, less elements copied than length of row")
		}
		newGrid = append(newGrid, copiedRow)
	}
	return
}

func printGrid(grid Grid, currMinute int) {
	if !DEBUG_ON {
		return
	}
	if currMinute == 0 {
		fmt.Println("Initial state:")
	} else {
		var extraChr rune
		if currMinute > 1 {
			extraChr = 's'
		}
		fmt.Printf("After %d minute%c:\n", currMinute, extraChr)
	}
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Printf("\n")
	}
}

func main() {
	grid := parseInput()
	if DEBUG_ON {
		fmt.Println(grid)
	}
	answer := performGridGenerations(grid, 10)
	fmt.Println(answer)
}
