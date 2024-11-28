package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type CartState int
type CartDirection rune

const (
	CartStateTurnLeft   CartState = 0
	CartStateGoStraight CartState = 1
	CartStateTurnRight  CartState = 2

	CartDirectionUp CartDirection = '^'
	CartDirectionDown CartDirection = 'v'
	CartDirectionLeft CartDirection = '<'
	CartDirectionRight CartDirection = '>'

	TrackVertical rune = '|'
	TrackHorizontal rune = '-'
	TrackCurve1 rune = '/'
	TrackCurve2 rune = '\\'
	TrackIntersection rune = '+'

	MaxCartStates int = 3
)

type Position struct {
	X int
	Y int
}

type Cart struct {
	Pos   Position
	State CartState
	Direction CartDirection
}

type CartTrack struct {
	Grid  [][]rune
	OriginalGrid  [][]rune
	Carts []Cart
}

var debugFlag *bool

func parseInput() (cartTrack CartTrack) {
	scanner := bufio.NewScanner(os.Stdin)

	linesRead := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		linesRead = append(linesRead, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if *debugFlag {
		log.Println("linesRead", linesRead)
	}

	cartTrack.Grid = make([][]rune, len(linesRead))
	cartTrack.OriginalGrid = make([][]rune, len(linesRead))
	for i := range cartTrack.Grid {
		cartTrack.Grid[i] = make([]rune, len(linesRead[i]))
		cartTrack.OriginalGrid[i] = make([]rune, len(linesRead[i]))
		for j := range cartTrack.Grid[i] {
			chr := rune(linesRead[i][j])
			cartTrack.Grid[i][j] = chr
			cartTrack.OriginalGrid[i][j] = chr
			if chr == rune(CartDirectionRight) || chr == rune(CartDirectionDown) || chr == rune(CartDirectionLeft) || chr == rune(CartDirectionUp) {
				cartTrack.Carts = append(cartTrack.Carts, Cart{
					Pos: Position{
						X: j,
						Y: i,
					},
					State: CartStateTurnLeft,
					Direction: CartDirection(chr),
				})
				if chr == rune(CartDirectionRight) || chr == rune(CartDirectionLeft) {
					cartTrack.OriginalGrid[i][j] = TrackHorizontal
				} else {
					cartTrack.OriginalGrid[i][j] = TrackVertical
				}
			}
		}
	}

	return
}

func iterateCartTrack(cartTrack CartTrack, stopOnFirstCrash bool) (firstCrash Position, lastPos Position) {
	var crashDetected bool
	if *debugFlag {
		drawBoard(cartTrack)
	}
mainLoop:
	for {
		for i, cart := range cartTrack.Carts {
			var newPos Position
			switch cart.Direction {
			case CartDirectionUp:
				newPos = Position{
					X: cart.Pos.X,
					Y: cart.Pos.Y - 1,
				}
			case CartDirectionDown:
				newPos = Position{
					X: cart.Pos.X,
					Y: cart.Pos.Y + 1,
				}
			case CartDirectionLeft:
				newPos = Position{
					X: cart.Pos.X - 1,
					Y: cart.Pos.Y,
				}
			case CartDirectionRight:
				newPos = Position{
					X: cart.Pos.X + 1,
					Y: cart.Pos.Y,
				}
			}
			if newPos.X >= 0 && newPos.X < len(cartTrack.Grid[0]) && newPos.Y >= 0 && newPos.Y < len(cartTrack.Grid) {
				// If the cart reaches a curve or intersection, it will need to make a turn
				switch (cartTrack.Grid[newPos.Y][newPos.X]) {
				case TrackCurve1:
					switch cart.Direction {
					case CartDirectionUp:
						cart.Direction = CartDirectionRight
					case CartDirectionDown:
						cart.Direction = CartDirectionLeft
					case CartDirectionLeft:
						cart.Direction = CartDirectionDown
					case CartDirectionRight:
						cart.Direction = CartDirectionUp
					}
				case TrackCurve2:
					switch cart.Direction {
					case CartDirectionUp:
						cart.Direction = CartDirectionLeft
					case CartDirectionDown:
						cart.Direction = CartDirectionRight
					case CartDirectionLeft:
						cart.Direction = CartDirectionUp
					case CartDirectionRight:
						cart.Direction = CartDirectionDown
					}
				case TrackIntersection:
					switch cart.State {
					case CartStateTurnLeft:
						switch cart.Direction {
						case CartDirectionUp:
							cart.Direction = CartDirectionLeft
						case CartDirectionDown:
							cart.Direction = CartDirectionRight
						case CartDirectionLeft:
							cart.Direction = CartDirectionDown
						case CartDirectionRight:
							cart.Direction = CartDirectionUp
						}
					case CartStateGoStraight:
						// don't change direction
					case CartStateTurnRight:
						switch cart.Direction {
						case CartDirectionUp:
							cart.Direction = CartDirectionRight
						case CartDirectionDown:
							cart.Direction = CartDirectionLeft
						case CartDirectionLeft:
							cart.Direction = CartDirectionUp
						case CartDirectionRight:
							cart.Direction = CartDirectionDown
						}
					}
					cart.State = (cart.State + 1) % CartState(MaxCartStates)
				}
				cartTrack.Grid[cart.Pos.Y][cart.Pos.X] = cartTrack.OriginalGrid[cart.Pos.Y][cart.Pos.X]
				cart.Pos = newPos
			}
			// fmt.Println("checking here", i, len(cartTrack.Carts))
			cartTrack.Carts[i] = cart
			switch cart.Direction {
			case CartDirectionUp:
				cartTrack.Grid[newPos.Y][newPos.X] = rune(CartDirectionUp)
			case CartDirectionDown:
				cartTrack.Grid[newPos.Y][newPos.X] = rune(CartDirectionDown)
			case CartDirectionLeft:
				cartTrack.Grid[newPos.Y][newPos.X] = rune(CartDirectionLeft)
			case CartDirectionRight:
				cartTrack.Grid[newPos.Y][newPos.X] = rune(CartDirectionRight)
			}
			crashedCarts := make(map[int]bool)
			for j, otherCart := range cartTrack.Carts {
				if i != j && cart.Pos.X == otherCart.Pos.X && cart.Pos.Y == otherCart.Pos.Y {
					crashedCarts[i] = true
					crashedCarts[j] = true
					// part 1 - detecting the first crash
					if (!crashDetected) {
						firstCrash = cart.Pos
						if *debugFlag {
							fmt.Println("first crash", firstCrash)
						}
						crashDetected = true
						if stopOnFirstCrash {
							return
						}
					}
				}
			}
			if len(crashedCarts) > 0 {
				var remainingCarts = make([]Cart, 0, len(cartTrack.Carts) - len(crashedCarts))
				for j, checkCart := range cartTrack.Carts {
					if !crashedCarts[j] {
						remainingCarts = append(remainingCarts, checkCart)
					} else {
						cartTrack.Grid[checkCart.Pos.Y][checkCart.Pos.X] = cartTrack.OriginalGrid[checkCart.Pos.Y][checkCart.Pos.X]
					}
				}
				fmt.Println("before:", cartTrack.Carts)
				cartTrack.Carts = remainingCarts
				fmt.Println("after:", cartTrack.Carts)
				fmt.Println("crash detected, starting back up")
				if len(cartTrack.Carts) == 0 {
					return
				}
				continue mainLoop
			}
		}
		if *debugFlag {
			drawBoard(cartTrack)
			time.Sleep(1 * time.Second)
		}
	}
}

func drawBoard(cartTrack CartTrack) {
	fmt.Println("----------------------")
	for _, row := range cartTrack.Grid {
		fmt.Println(string(row))
	}
	fmt.Println("----------------------")
}

func main() {
	debugFlag = flag.Bool("debug", false, "Turn on debug options")
	flag.Parse()

	cartTrack := parseInput()

	if *debugFlag {
		log.Println("cartTrack", cartTrack)
	}

	firstCrash, _ := iterateCartTrack(cartTrack, true)

	fmt.Printf("%d,%d\n", firstCrash.X, firstCrash.Y)
	// fmt.Printf("%d,%d\n", lastPos.X, lastPos.Y)
}
