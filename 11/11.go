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

type LargestSquareResults struct {
	x int
	y int
	power int
	size int
	square [][]int
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

func (this *Hologram) MapSquare(x int, y int, size int) [][]int {
	if x >= len(this.board) - (size - 1) ||
		y >= len(this.board) - (size - 1) || 
		size < 1 || size > 300 {
		return nil
	}
	
	mapped := make([][]int, size)
	index := 0
	for i := y; i < y + size; i++ {
		mapped[index] = this.board[i][x:x+size]
		index += 1
	}
	return mapped
}

func (this *Hologram) CheckHighestPower(x int, y int, width int, height int, size int) (int, int, int, [][]int) {
	// Create the arrays
	squares := make([][][]int, height * width)
	for i := y; i < y + height; i++ {
		for j := x; j < x + width; j++ {
			squares[height * i + j] = this.MapSquare(j, i, size)
		}
	}

	// Now check them
	largest := 300 * 300 * -5 // the smallest total power value
	largestX := 0
	largestY := 0

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

func (this *LargestSquareResults) Print() {
	fmt.Println("************")
	fmt.Println(this.x, this.y, this.size, this.power)
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

func findLargestSquare(hologram *Hologram, debug bool) (int, int, int) {
	if debug {
		fmt.Println("************")
	}
	
	// This could be parallelized to improve performance.
	largestArr := make([]*LargestSquareResults, 300)
	for i, _ := range largestArr {
		x, y, highest, square := hologram.CheckHighestPower(0,0,300 - i,300 - i, i + 1)
		largestArr[i] = &LargestSquareResults {
			x: x,
			y: y,
			power: highest,
			size: i + 1,
			square: square,
		}
		if debug {
			largestArr[i].Print()
		}
	}

	largestIndex := -1
	largestPower := 0

	for i, lg := range largestArr {
		if lg.power > largestPower {
			largestPower = lg.power
			largestIndex = i
		}
	}

	if largestIndex == -1 {
		log.Fatal("Unexpected: No largest square found.")
	}

	largest := largestArr[largestIndex]

	if debug {
		fmt.Printf("Highest power is at %d,%d with value %d and size %d\n",
			largest.x,
			largest.y,
			largest.power,
			largest.size)
		printSquare(largest.square)
	}
	
	return largest.x, largest.y, largest.size
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

	// pt1X, pt1Y := findLargestSquare(hologram, *debugFlag)
	// fmt.Println(pt1X, pt1Y)
	pt2X, pt2Y, pt2Size := findLargestSquare(hologram, *debugFlag)
	fmt.Println(pt2X, pt2Y, pt2Size)
}