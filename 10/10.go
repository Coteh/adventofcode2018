package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x int
	y int
}

type LightPoint struct {
	position Vector
	velocity Vector
}

func parseVector(input string, debug bool) Vector {
	newVec := Vector{}

	removedArrows := input[1 : len(input)-1]
	splitArr := strings.Split(removedArrows, ",")

	px, err := strconv.ParseInt(strings.TrimLeft(splitArr[0], " "), 10, 32)
	if err != nil {
		log.Fatal("Could not parse x coord")
	}
	newVec.x = int(px)
	if debug {
		fmt.Println(splitArr)
	}
	py, err := strconv.ParseInt(strings.TrimLeft(splitArr[1], " "), 10, 32)
	if err != nil {
		log.Fatal("Could not parse y coord")
	}
	newVec.y = int(py)

	return newVec
}

func parseLightPoint(input string, debug bool) LightPoint {
	lightPoint := LightPoint{}
	splitArr := strings.Split(input, "=")
	parseState := 0
	for _, val := range splitArr {
		switch parseState {
		case 1:
			// remove velocity portion from split string
			valSplit := strings.Split(val, " v")
			lightPoint.position = parseVector(valSplit[0], debug)
		case 2:
			lightPoint.velocity = parseVector(val, debug)
		default: // initial state, continue to below
		}

		if strings.Contains(val, "position") {
			parseState = 1
		} else if strings.Contains(val, "velocity") {
			parseState = 2
		}
	}

	return lightPoint
}

func printBoard(board [][]rune) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Print("\n")
	}
}

func calculateCenterPoint(workingPoints []Vector, boardLength int) Vector {
	center := Vector{}

	for _, pos := range workingPoints {
		center.x += pos.x
		center.y += pos.y
	}
	center.x /= len(workingPoints)
	center.y /= len(workingPoints)

	return center
}

// Sum up all distances between points and the center point
// I got the idea of checking in at least one dimension
// from here https://stackoverflow.com/a/29729612
func calculateClosenessValue(points []Vector, centerPoint Vector) float64 {
	var totalDist float64
	for _, pos := range points {
		totalDist += math.Abs(float64(pos.x-centerPoint.x)) + math.Abs(float64(pos.y-centerPoint.y))
	}

	return totalDist
}

func runLoop(lightPoints []LightPoint, boardLength int, endTime int) {
	if len(lightPoints) == 0 {
		log.Fatal("No light points")
	}

	board := make([][]rune, boardLength)
	for i := range board {
		board[i] = make([]rune, boardLength)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var x, y, boardX, boardY, originX, originY int
	var center Vector
	workingPoints := make([]Vector, len(lightPoints))
	clearPoints := make([]Vector, len(lightPoints))

	// Records the closest total distance of all points to the center point
	var lowestCloseness float64 = math.MaxFloat64

	for time := 0; time <= endTime; time++ {
		// update working points
		for i, lp := range lightPoints {
			workingPoints[i].x = lp.position.x + (lp.velocity.x * time)
			workingPoints[i].y = lp.position.y + (lp.velocity.y * time)
		}

		// calculate center
		center = calculateCenterPoint(workingPoints, boardLength)

		// calculate relative (0,0) point
		originX = center.x - boardLength/2
		originY = center.y - boardLength/2

		// Once the points start scattering again, print out what the board was at the time,
		// which should have the message
		closeness := calculateClosenessValue(workingPoints, center)
		if closeness < lowestCloseness {
			// Record the new closest total distance
			lowestCloseness = closeness

			// clear board of old points if any
			if time > 0 {
				for _, cp := range clearPoints {
					if cp.x >= 0 && cp.y >= 0 {
						board[cp.y][cp.x] = '.'
					}
				}
			}

			// clear clear points
			for i := range clearPoints {
				clearPoints[i].x = -1
				clearPoints[i].y = -1
			}

			// plot new points
			for i, pos := range workingPoints {
				boardX = -1
				boardY = -1
				x = pos.x
				y = pos.y
				if x >= center.x-boardLength/2 &&
					x < center.x+boardLength/2 &&
					y >= center.y-boardLength/2 &&
					y < center.y+boardLength/2 {
					boardX = x - originX
					boardY = y - originY
					clearPoints[i].x = boardX
					clearPoints[i].y = boardY
				}
				if boardX >= 0 && boardY >= 0 &&
					boardX < boardLength &&
					boardY < boardLength {
					board[boardY][boardX] = '#'
				}
			}
		} else {
			printBoard(board)
			fmt.Println(time - 1)
			break
		}
	}
}

func main() {
	debugFlag := flag.Bool("debug", false, "Turn on debug options")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	lightPoints := make([]LightPoint, 0, 5)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal("Encountered an error with input")
				os.Exit(1)
			}
			break
		}
		input = strings.TrimRight(input, "\n")
		if len(input) == 0 {
			continue
		}

		if *debugFlag {
			fmt.Println(input)
		}

		lightPoint := parseLightPoint(input, *debugFlag)
		if *debugFlag {
			fmt.Println(lightPoint)
		}
		lightPoints = append(lightPoints, lightPoint)
	}

	// I noticed in one of my debugging sessions
	// that there were some points really close to each
	// other within the 10,000 second mark, so
	// I increased end time by 10,000 manually for each run
	// until something printed out. Probably not the best
	// way to do it, but hopefully one day I can learn
	// a better way to detect a plot that can be printed
	// without having to manually set end time.
	runLoop(lightPoints, 100, 20000)
}
