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
	"strconv"
)

type StepProgress int32

const (
	NOT_STARTED StepProgress 	= iota
	SCHEDULED	StepProgress	= iota
	STARTED 	StepProgress 	= iota
	COMPLETED 	StepProgress 	= iota
)

type StepNode struct {
	children []*StepNode
	step rune
	prereqs []rune
	progress StepProgress
}

type StepTable struct {
	values map[rune]*StepNode
	stepTracker map[rune]bool
	printed map[rune]bool
}

type StepWorkItem struct {
	workerID int
	endTime int
	stepNode *StepNode
	isWorking bool
}

type StepQueue struct {
	availList []*StepNode
}

func createTreeNode(step rune) *StepNode {
	return &StepNode {
		step: step,
		children: make([]*StepNode, 0, 5),
		prereqs: make([]rune, 0, 5),
		progress: NOT_STARTED,
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

func createQueue(stepTable *StepTable) *StepQueue {
	queue := &StepQueue{
		availList: make([]*StepNode, 0, 5),
	}

	// Poll for initial available nodes that satisfy all prereqs
	// ie. they have no prereqs
	for _, node := range stepTable.values {
		if len(node.prereqs) == 0 {
			queue.availList = append(queue.availList, node)
			node.progress = SCHEDULED
		}
	}

	// Sort the current nodes alphabetically
	queue.Sort()

	return queue
}

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

func (this *StepTable) CheckPrereqCompletion(node *StepNode) bool {
	if node == nil {
		return false
	}
	
	for _, prereq := range node.prereqs {
		if this.values[prereq].progress != COMPLETED {
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

func (this *StepQueue) HasAnymoreSteps() bool {
	return len(this.availList) > 0
}

func (this *StepQueue) Sort() {
	sort.Slice(this.availList, func(i, j int) bool {
		return this.availList[i].step < this.availList[j].step
	})
}

func (this *StepQueue) DequeueNextAvailableStep(stepTable *StepTable) *StepNode {
	nodesAvailable := this.HasAnymoreSteps()
	
	if !nodesAvailable {
		return nil
	}
	
	// Find next available node that 
	// hasn't been started already
	var nextNode *StepNode
	popIndex := -1
	for index, node := range this.availList {
		if node.progress == SCHEDULED && stepTable.CheckPrereqCompletion(node) {
			nextNode = node
			popIndex = index
			break
		}
	}

	if popIndex >= 0 {
		this.availList = append(this.availList[0:popIndex], this.availList[popIndex + 1:]...)
	}

	return nextNode
}

func (this *StepQueue) EnqueueStepChildren(node *StepNode) {
	valuesAdded := false

	for _, child := range node.children {
		if child.progress == NOT_STARTED {
			this.availList = append(this.availList, child)
			child.progress = SCHEDULED
			valuesAdded = true
		}
	}

	if valuesAdded {
		this.Sort()
	}
}

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

func calculateTimeRequired(step rune) int {
	return 60 + (int(step) - int('A') + 1)
}

func workSteps(stepTable *StepTable, numWorkers int, debug bool) int {
	time := 0
	stepsCompleted := false
	workItems := make(map[int]*StepWorkItem)
	queue := createQueue(stepTable)
	workCount := 0

	// Initialize work items
	// (each worker gets a work item - we'll reuse them)
	for i := 0; i < numWorkers; i++ {
		workItems[i] = &StepWorkItem {
			workerID: i,
			endTime: 0,
			stepNode: nil,
			isWorking: false,
		}
	}

	for !stepsCompleted {
		if debug {
			fmt.Println("At " + strconv.Itoa(time) + " seconds")
		}
		// If any workers have completed a step,
		// then take the step off the work item,
		// and enqueue children of the step into queue
		for _, workItem := range workItems {
			if workItem.isWorking && time >= workItem.endTime {
				workItem.stepNode.progress = COMPLETED
				queue.EnqueueStepChildren(workItem.stepNode)
				workItem.stepNode = nil
				workItem.isWorking = false
				workCount -= 1
			}
		}
		// If there are no more available tasks,
		// then set stepsCompleted to true and end loop
		if (!queue.HasAnymoreSteps() && workCount == 0) {
			stepsCompleted = true
			continue
		}
		// Assign unassigned workers to available tasks
		// If there are no workers available atm, then
		// we'll just continue to next iteration
		for _, workItem := range workItems {
			if !workItem.isWorking {
				workItem.stepNode = queue.DequeueNextAvailableStep(stepTable)
				if workItem.stepNode == nil {
					continue
				}
				if debug {
					fmt.Println("The next available work item is task " + string(workItem.stepNode.step) + ".")
				}
				workItem.endTime = time + calculateTimeRequired(workItem.stepNode.step)
				workItem.isWorking = true
				workItem.stepNode.progress = STARTED
				workCount += 1
				if debug {
					fmt.Println("Worker " + strconv.Itoa(workItem.workerID) + " started on task " + string(workItem.stepNode.step) + ": Should be done at " + strconv.Itoa(workItem.endTime) + " seconds.")
				}
			}
		}
		// Increment time by a second
		time += 1
	}

	return time
}

func testSteps() {
	time := 60

	for i := 0; i < 26; i++ {
		expectedTime := (time + 1 + i)
		expectedTimeStr := strconv.Itoa(expectedTime)
		fmt.Println("Step " + string('A' + i) + " should take " + expectedTimeStr + " seconds.")
		if calculateTimeRequired(rune('A' + i)) != expectedTime {
			log.Fatal("Amount of work required for step " + string('A' + i) + " is incorrect.")
		}
	}

	fmt.Println("All steps are correct")
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
			break
		}
		input = strings.TrimRight(input, "\n")
		stepsArr = append(stepsArr, input)
		if *debugFlag {
			fmt.Println(input)
		}
	}

	stepTable = parseSteps(stepsArr, *debugFlag)

	stepTable.Print()

	if *debugFlag {
		testSteps()
	}

	pt2Answer := workSteps(stepTable, 5, *debugFlag)
	fmt.Println(pt2Answer)
}