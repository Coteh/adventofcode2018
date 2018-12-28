package main

import (
	"fmt"
	"flag"
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"strconv"
)

type Hologram struct {
	board [][]int
}

func initHologram() *Hologram {
	board := make([][]int, 300)

	for i, _ := range board {
		board[i] = make([]int, 300)
	}

	return &Hologram {
		board: board,
	}
}

func (this *Hologram) Print() {
	printSquare(this.board)
}

func (this *Hologram) MapSquare(x int, y int) [][]int {
	if x >= len(this.board) - 2 || y >= len(this.board) - 2 {
		return nil
	}
	
	mapped := make([][]int, 3)
	index := 0
	for i := y; i < y + 3; i++ {
		mapped[index] = this.board[i][x:x+3]
		index += 1
	}
	return mapped
}

func (this *Hologram) CheckHighestPower(x int, y int, width int, height int) (int, int, int, [][]int) {
	// Create the arrays
	squares := make([][][]int, height * width)
	for i := y; i < y + height; i++ {
		for j := x; j < x + width; j++ {
			squares[height * i + j] = this.MapSquare(j, i)
		}
	}

	// Now check them
	largest := 0
	largestX := -1
	largestY := -1

	for i := y; i < y + height; i++ {
		for j := x; j < x + width; j++ {
			cell := checkSquare(squares[i * height + j])
			if cell > largest {
				largest = cell
				largestX = j
				largestY = i
			}
		}
	}

	return largestX, largestY, largest, squares[largestY * height + largestX]
}

func getPowerLevel(x int, y int, serialNum int) int {
	// Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackID := x + 10
	// Begin with a power level of the rack ID times the Y coordinate.
	powerLevel := rackID * y
	// Increase the power level by the value of the grid serial number (your puzzle input).
	powerLevel += serialNum
	// Set the power level to itself multiplied by the rack ID.
	powerLevel *= rackID
	// Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	hundredsDigit := powerLevel / 10 / 10 % 10
	// Subtract 5 from the power level.
	return hundredsDigit - 5
}

func createHologramFromSerialNumber(serialNum int) *Hologram {
	hologram := initHologram()

	var wg sync.WaitGroup

	for i, row := range hologram.board {
		for j, _ := range row {
			wg.Add(1)
			go (func(x int, y int, wg *sync.WaitGroup) {
				defer wg.Done()				
				hologram.board[y][x] = getPowerLevel(x, y, serialNum)
			})(i, j, &wg)
		}
	}

	wg.Wait()

	return hologram
}

func printSquare(square [][]int) {
	fmt.Println("----------------")
	for _, row := range square {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Print("\n")
	}
	fmt.Println("----------------")
}

func checkSquare(square [][]int) int {
	sum := 0

	for _, row := range square {
		for _, cell := range row {
			sum += cell
		}
	}

	return sum
}

func findLargestSquare(hologram *Hologram, debug bool) (int, int) {
	// There are 298 x 298 = 88804 3x3 squares
	// to check if we don't apply any heuristics
	// to narrow down results.
	// This could be parallelized to improve performance.
	x, y, highest, square := hologram.CheckHighestPower(0,0,298,298)

	if debug {
		fmt.Printf("Highest power is at %d,%d with value %d: \n", x, y, highest)
		printSquare(square)
	}
	
	return x, y
}

func testPowerLevel() {
	powerLevel := getPowerLevel(3,5,8)
	if powerLevel != 4 {
		panic("Power level calculation not correct")
	}
}

func main() {
	debugFlag := flag.Bool("debug", false, "Turn on debug options")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	
	input, err := reader.ReadString('\n')
	if err != nil {
		if err != io.EOF {
			log.Fatal("Encountered an error with input")
			os.Exit(1)
		}
	}
	input = strings.TrimRight(input, "\n")
	if input == "" {
		log.Fatal("No input was given.")
	}

	if *debugFlag {
		fmt.Println(input)
		testPowerLevel()
	}

	serialNum, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal("Could not parse serial number")
	}

	hologram := createHologramFromSerialNumber(int(serialNum))

	if *debugFlag {
		hologram.Print()
	}

	pt1X, pt1Y := findLargestSquare(hologram, *debugFlag)
	fmt.Println(pt1X, pt1Y)
}