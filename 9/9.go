package main

import (
	"fmt"
	"strings"
	"strconv"
	"flag"
	"log"
	"os"
	"bufio"
	"io"
)

type MarbleNode struct {
	pointValue int
	next *MarbleNode
	prev *MarbleNode
	end bool
}

type MarbleList struct {
	head *MarbleNode
	curr *MarbleNode
}

func createMarbleNode(pointValue int) *MarbleNode {
	return &MarbleNode {
		pointValue: pointValue,
	}
}

func createMarbleList() *MarbleList {
	return &MarbleList {}
}

func (this *MarbleNode) SetNext(node *MarbleNode) {
	oldNext := this.next
	this.next = node
	
	node.next = oldNext
	node.prev = this

	node.next.prev = node

	node.end = this.end
	this.end = false
}

func (this *MarbleNode) Unlink() {
	this.prev.next = this.next
	this.next.prev = this.prev
	this.prev.end = this.end
}

func (this *MarbleList) SetHead(node *MarbleNode) {
	this.head = node
}

func (this *MarbleList) SetCurrentMarble(node *MarbleNode) {
	this.curr = node
}

func (this *MarbleList) SetWrap(node *MarbleNode) {
	node.next = this.head
	this.head.prev = node.next
	node.end = true
}

func (this *MarbleList) PlaceMarbleClockwise(node *MarbleNode) {
	// When placing marble 0 down
	if this.head == nil {
		this.SetHead(node)
		this.SetWrap(node)
		return
	}

	if this.curr == nil {
		log.Fatal("Unexpected: Current node set to nil")
	} else if this.curr.next == nil {
		log.Fatal("Unexpected: Current node's next node set to nil")
	}

	this.curr.next.SetNext(node)
	// If next node is end node, this one will be the end node
	// if this.curr.next.end {
		// this.SetWrap(node)
	// }

	this.SetCurrentMarble(node)
}

func (this *MarbleList) Print() {
	tempNode := this.head
	for tempNode != nil {
		if tempNode == this.curr {
			fmt.Print("(")
		}
		fmt.Printf("%d", tempNode.pointValue)
		if tempNode == this.curr {
			fmt.Print(")")
		}
		fmt.Print(" ")
		if tempNode.end {
			tempNode = nil
		} else {
			tempNode = tempNode.next		
		}
	}
	fmt.Print("\n")
}

func beginGame(numPlayers int, lastPoint int, debug bool) int {
	marbleList := createMarbleList()
	scoreboard := make([]int, numPlayers)
	var tempNode *MarbleNode
	var playerIndex, pointValue int

	for i := 0; i <= lastPoint; i++ {
		playerIndex = i % numPlayers
		tempNode = createMarbleNode(i)
		if marbleList.head == nil {
			marbleList.SetHead(tempNode)
			marbleList.SetWrap(tempNode)
			marbleList.SetCurrentMarble(tempNode)
		} else {
			if i != 0 && i % 23 == 0 {
				// Save the point value for this
				// node to give player below
				pointValue = tempNode.pointValue
				// Set tempNode to current node
				tempNode = marbleList.curr
				// Go back 7 nodes counter-clockwise
				for j := 0; j < 7; j++ {
					if debug {
						fmt.Println(tempNode.pointValue)
					}
					tempNode = tempNode.prev
					if tempNode == nil {
						log.Fatal("Unexpected: tempNode is null")
					}
				}
				// Save node before removing it
				removedNode := tempNode
				// Grab reference to removed node's next node
				nextNode := removedNode.next
				// Remove the node by unlinking it
				removedNode.Unlink()
				// Set the current node to next node clockwise
				marbleList.SetCurrentMarble(nextNode)
				// Reward player points from this
				// removed node as well as node they
				// were going to place
				scoreboard[playerIndex] += removedNode.pointValue + pointValue
			} else {
				marbleList.PlaceMarbleClockwise(tempNode)
			}
		}
		if debug {
			fmt.Printf("***Player %d's turn****\n", playerIndex + 1)
			marbleList.Print()
			fmt.Println("Scoreboard:")
			for i, score := range scoreboard {
				fmt.Printf("[%d]-%d ", i + 1, score)
			}
			fmt.Print("\n")
			fmt.Println("**********************")
		}
	}

	maxPoints := 0
	for _, value := range scoreboard {
		if value > maxPoints {
			maxPoints = value
		}
	}
	return maxPoints
}

func getNumberOfPlayers(playersStr string) int {
	spaceIndex := strings.IndexRune(playersStr, ' ')
	if spaceIndex == -1 {
		log.Fatal("Number of players string invalid")
	}
	result, err := strconv.ParseInt(playersStr[:spaceIndex], 10, 32)
	if err != nil {
		log.Fatal("Could not parse number of players")
	}
	return int(result)
}

func getLastMarblePoint(lastMarbleStr string) int {
	cutStr := strings.TrimPrefix(lastMarbleStr, " last marble is worth ")
	pointsIndex := strings.LastIndex(cutStr, " points")
	if pointsIndex == -1 {
		log.Fatal("Last marble point string invalid")
	}
	result, err := strconv.ParseInt(cutStr[:pointsIndex], 10, 32)
	if err != nil {
		log.Fatal("Could not parse last marble point")
	}
	return int(result)
}

func parseMarbleGame(input string, debug bool) (int, int) {
	inputArr := strings.Split(input, ";")
	numPlayers := getNumberOfPlayers(inputArr[0])
	lastPoint := getLastMarblePoint(inputArr[1])
	if debug {
		fmt.Println(numPlayers)
		fmt.Println(lastPoint)
	}
	return numPlayers, lastPoint
}

func main() {
	debugFlag := flag.Bool("debug", false, "Turn on debug options")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	var numPlayers, lastPoint int

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
		if len(input) == 0 {
			continue
		}
		
		if *debugFlag {
			fmt.Println(input)
		}

		numPlayers, lastPoint = parseMarbleGame(input, *debugFlag)
	}

	pt1Answer := beginGame(numPlayers, lastPoint, *debugFlag)
	fmt.Println(pt1Answer)
}