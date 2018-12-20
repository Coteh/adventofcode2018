package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"io"
	"log"
	"strings"
	"regexp"
	"sort"
)

type StepNode struct {
	children []*StepNode
	step rune
	prereqs []rune
}

type StepTable struct {
	values map[rune]*StepNode
	stepTracker map[rune]bool
	printed map[rune]bool
}

// Might need a queue for part 2, so kept
// this stuff commented out for now
// type StepQueueItem struct {
// 	node *StepNode
// 	level int
// }

// type StepQueue struct {
// 	nodeList []*StepQueueItem
// 	existsMap map[rune]bool
// 	highestLevel int
// }

func createTreeNode(step rune) *StepNode {
	return &StepNode {
		step: step,
		children: make([]*StepNode, 0, 5),
		prereqs: make([]rune, 0, 5),
	}
}

func createTree() *StepNode {
	return createTreeNode(0)
}

func createStepTable() *StepTable {
	return &StepTable {
		values: make(map[rune]*StepNode),
		stepTracker: make(map[rune]bool),
		printed: make(map[rune]bool),
	}
}

// func createQueue() *StepQueue {
// 	return &StepQueue{
// 		nodeList: make([]*StepQueueItem, 0, 5),
// 		existsMap: make(map[rune]bool),
// 		highestLevel: -1,
// 	}
// }

func (this *StepNode) AddStepChild(node *StepNode) {
	if this == nil {
		log.Fatal("Error trying to add node as child to null node")
	}

	this.children = append(this.children, node)
	node.AddPrereq(this.step)
}

func (this *StepNode) DebugPrint() {
	fmt.Print("Node Val: " + string(this.step) + ", Children: ")
	for i, child := range this.children {
		fmt.Print(string(child.step))
		if i < len(this.children) - 1 {
			fmt.Print(", ")
		}
	}
	fmt.Print(", Prereqs: ")
	for i, prereq := range this.prereqs {
		fmt.Print(string(prereq))
		if i < len(this.prereqs) - 1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("\n")

	for _, child := range this.children {
		child.DebugPrint()
	}
}

func (this *StepNode) AddPrereq(prereq rune) {
	this.prereqs = append(this.prereqs, prereq)
}

func (this *StepTable) SetStep(step rune, prereq rune) {
	// Creates a node for the step and saves to table
	// if it's not already there
	if this.values[step] == nil {
		this.values[step] = createTreeNode(step)
	}
	// Creates a node for prereq step as well if it's not
	// already there
	if this.values[prereq] == nil {
		this.values[prereq] = createTreeNode(prereq)
	}
	// Set the step node as a child for the prereq node
	this.values[prereq].AddStepChild(this.values[step])
	// Track the step into step tracker, to see
	// which step in the end only shows up as a prereq (like step C in example)
	this.stepTracker[step] = true
}

func (this *StepTable) CheckPrereqs(node *StepNode) bool {
	if node == nil {
		return false
	}
	
	for _, prereq := range node.prereqs {
		if !this.printed[prereq] {
			return false
		}
	}

	return true
}

func (this *StepTable) Print() {
	availList := make([]*StepNode, 0, 5)
	// Poll for initial available nodes that satisfy all prereqs
	// ie. they have no prereqs
	for _, node := range this.values {
		if len(node.prereqs) == 0 {
			availList = append(availList, node)
		}
	}
	nodesAvailable := (len(availList) > 0)
	// If we found some nodes, print them like this:
	for nodesAvailable {
		// leastLetter := 'Z'
		var leastNode *StepNode
		// Now print all nodes that are marked as available
		// For each node available, also:
		// - mark them as printed, so they won't get selected again
		// - find next available nodes from within children of these nodes
		sort.Slice(availList, func(i, j int) bool {
			return availList[i].step < availList[j].step
		})
		leastNode = availList[0]
		fmt.Print(string(leastNode.step))
		this.printed[leastNode.step] = true
		// Create new list, appending remaining letters
		newList := make([]*StepNode, len(availList) - 1)
		// Copy remaining contents of availList over and
		// remove reference to old list
		if len(availList[1:]) > 0 {
			copy(newList, availList[1:])
		}
		availList = newList
		for _, child := range leastNode.children {
			if this.CheckPrereqs(child) && !this.printed[child.step] {
				availList = append(availList, child)
			}
		}
		
		// Any more nodes available?
		nodesAvailable = (len(availList) > 0)
	}
	fmt.Print("\n")
}

// func (this *StepQueue) HasStep(step rune) bool {
// 	return this.existsMap[step]
// }

// func (this *StepQueue) CheckPrereqs(node *StepNode) bool {
// 	for _, prereq := range node.prereqs {
// 		fmt.Println("Checking " + string(node.step) + " prereq, which is " + string(prereq))
// 		if !this.HasStep(prereq) {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (this *StepQueue) EnqueueStep(node *StepNode, level int) {
// 	if !this.CheckPrereqs(node) || this.existsMap[node.step] {
// 		return
// 	}
// 	item := &StepQueueItem {
// 		node: node,
// 		level: level,
// 	}
// 	this.existsMap[node.step] = true

// 	indexToInsert := -1

// 	fmt.Println(level)

// 	// Insert new queue item alphabetically
// 	for i, item := range this.nodeList {
// 		if level < item.level && int(node.step) < int(item.node.step) {
// 			indexToInsert = i
// 			break
// 		}
// 	}

// 	if (indexToInsert == -1) {
// 		this.nodeList = append(this.nodeList, item)
// 	} else {
// 		this.nodeList = append(this.nodeList, nil)
// 		copy(this.nodeList[indexToInsert + 1:], this.nodeList[indexToInsert:])
// 		this.nodeList[indexToInsert] = item
// 	}
// }

// func (this *StepQueue) Print() {
// 	for _, item := range this.nodeList {
// 		fmt.Print(string(item.node.step))
// 	}
// 	fmt.Print("\n")
// }

func parseSteps(stepsArr []string, debug bool) *StepTable {
	// Instantiate our data structures
	table := createStepTable()

	if (debug) {
		fmt.Println(stepsArr)
	}

	re, err := regexp.Compile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")
	if err != nil {
		log.Fatal("Error compiling step parser regex")
	}

	// First we will put all nodes initialized into hash table
	for _, stepStr := range stepsArr {
		matches := re.FindStringSubmatch(stepStr)
		prereq := rune(matches[1][0])
		step := rune(matches[2][0])
		if (debug) {
			fmt.Println("---------\nStep: " + string(step))
			fmt.Println("Prereq: " + string(prereq) + "\n-----------\n")
		}
		table.SetStep(step, prereq)
	}

	if (debug) {
		fmt.Println(table.values)
	}

	return table
}

func main() {
	debugFlag := flag.Bool("debug", false, "Turn on debug options")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	var stepTable *StepTable
	stepsArr := make([]string, 0, 10)

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
		stepsArr = append(stepsArr, input)
		if *debugFlag {
			fmt.Println(input)
		}
	}

	stepTable = parseSteps(stepsArr, *debugFlag)

	stepTable.Print()
}