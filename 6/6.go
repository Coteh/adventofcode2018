package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"io"
	"strings"
	"strconv"
	"flag"
	"math"
)

// Notes
// - If a coordinate has a distance
// on board that is solely closest to
// it and is on edge of board, it's
// an infinite distance
// - We'll mark positions on board with
// the coord id if it's solely closest
// to that coord. If there's a position
// closest to more than one coord, it'll
// be marked with -1

type Coord struct {
	x int
	y int
}

type Board struct {
	values [][]int
	coords []Coord
	width int
	height int
	areas map[int]int
	safeArea int
}

func buildBoard(coords []Coord) *Board {
	// Calculate width and height
	// of observable board based on
	// coords we receive from input
	largestX := 0
	largestY := 0
	for _, value := range coords {
		if value.x > largestX {
			largestX = value.x
		}
		if value.y > largestY {
			largestY = value.y
		}
	}
	width := largestX + 2
	height := largestY + 2

	// Create observable board
	board := make([][]int, width)
	for i := 0; i < width; i++ {
		board[i] = make([]int, height)
	}

	// Create areas hash table
	// which will tally up areas for each
	// infinite and finite areas as we do the next step
	areasMap := make(map[int]int)

	// Calculate Manhattan distances
	// NOTE: This could be a separate
	// function but for brevity it'll be
	// in here.
	// Going to be yucky O(n^3) for now.
	// Also going to calculate safe area in here as well
	// even though that also can be a separate function.
	// Assuming that there is only one safe region on map.
	safeArea := 0
	for i, row := range board {
		for j, _ := range row {
			shortest := width * height
			sum := 0
			var shortestIDs []int
			// NOTE: each coord will have different
			// k value per run since we parallelized
			// the parsing
			for k, coord := range coords {
				manhattan := int(math.Abs(float64(i - coord.x))) + int(math.Abs(float64(j - coord.y)))
				sum += manhattan
				// This is for part 1
				if manhattan < shortest {
					shortestIDs = make([]int, 0, 2)
					shortestIDs = append(shortestIDs, k)
					shortest = manhattan
				} else if manhattan == shortest {
					if shortestIDs == nil {
						log.Fatal("shortestIDs is nil")
					}
					shortestIDs = append(shortestIDs, k)
				}
			}
			// This is for part 2
			if sum < 10000 {
				safeArea += 1
			}
			// This is for part 1
			length := len(shortestIDs)
			if length == 1 {
				id := shortestIDs[0]
				board[i][j] = id
				areasMap[id] += 1
			} else if length > 1 {
				board[i][j] = -1
			}
		}
	}

	return &Board {
		values: board,
		coords: coords,
		width: width,
		height: height,
		areas: areasMap,
		safeArea: safeArea,
	}
}

func determineInfiniteAreas(board *Board) map[int]bool {
	if board == nil || len(board.values) == 0 {
		log.Fatal("Uninitialized board passed into determineInfiniteAreas")
	}
	// Anything not -1 on the borders of the board are infinite areas
	cellMap := make(map[int]bool)

	// Check left side
	for _, cell := range board.values[0] {
		if cell != -1 {
			cellMap[cell] = true
		}
	}
	// Check right side
	for _, cell := range board.values[len(board.values) - 1] {
		if cell != -1 {
			cellMap[cell] = true
		}
	}
	// Check top and bottom sides
	for _, col := range board.values {
		topCell := col[0]
		bottomCell := col[len(col) - 1]
		if topCell != -1 {
			cellMap[topCell] = true
		}
		if bottomCell != -1 {
			cellMap[bottomCell] = true
		}
		
	}

	return cellMap
}

func findLargestFiniteArea(board *Board) int {
	infinites := determineInfiniteAreas(board)
	largestArea := 0

	for key, value := range board.areas {
		if value > largestArea && !infinites[key] {
			largestArea = value
		}
	}

	return largestArea
}

func parseCoord(input string, coordChan chan Coord) {
	splitArr := strings.Split(input, ", ")
	x, err := strconv.Atoi(splitArr[0])
	if err != nil {
		log.Fatal("Could not parse x coord")
	}
	y, err := strconv.Atoi(splitArr[1])
	if err != nil {
		log.Fatal("Could not parse y coord")
	}
	coordChan <- Coord{x, y}
}

func main() {
	debugFlag := flag.Bool("debug", false, "Turn on debug options")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	numCoords := 0
	numCoordsParsed := 0
	coordChan := make(chan Coord)
	
	for true {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal("Encountered an error with input")
				os.Exit(1)
			}
			break;
		}
		input = strings.TrimRight(input, "\n")
		if *debugFlag {
			fmt.Println(input)
		}
		go parseCoord(input, coordChan)
		numCoords += 1
	}

	coords := make([]Coord, 0, numCoords)

	for coord := range coordChan {
		coords = append(coords, coord)
		numCoordsParsed += 1
		if numCoordsParsed >= numCoords {
			break
		}
	}

	if *debugFlag {
		fmt.Println(coords)
		fmt.Println(len(coords))
	}

	board := buildBoard(coords)

	if *debugFlag {
		fmt.Println(board)
	}

	pt1Answer := findLargestFiniteArea(board)
	fmt.Println(pt1Answer)
	pt2Answer := board.safeArea
	fmt.Println(pt2Answer)
}