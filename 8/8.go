package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"io"
	"log"
	"strings"
	"strconv"
)

type ParseState int

const (
	CREATE_TREE_NODE	ParseState = iota
	PARSE_CHILD_COUNT 	ParseState = iota
	PARSE_META_COUNT	ParseState = iota
	PARSE_META_DATA		ParseState = iota
	PARSE_COMPLETED		ParseState = iota
)

type TreeNode struct {
	numChildren int
	numMetadata int
	children []*TreeNode
	metadata []int
	parent *TreeNode
	numChildrenRead int
}

func createTreeNode(numChildren int, numMetadata int) *TreeNode {
	return &TreeNode {
		numChildren: numChildren,
		numMetadata: numMetadata,
		children: make([]*TreeNode, 0, numChildren),
		metadata: make([]int, 0, numMetadata),
	}
}

func (this *TreeNode) AddChild(child *TreeNode) {
	this.children = append(this.children, child)
	child.parent = this
	this.numChildrenRead += 1
}

func (this *TreeNode) AddMetadata(metadata int) {
	this.metadata = append(this.metadata, metadata)
}

func (this *TreeNode) GetSum() int {
	total := 0

	for _, val := range this.metadata {
		total += val
	}

	for _, child := range this.children {
		total += child.GetSum()
	}

	return total
}

func (this *TreeNode) Print(level int) {
	fmt.Println("----------------")
	_, err := fmt.Printf("Level %d - This node has %d children and %d metadata\n", level, this.numChildren, this.numMetadata)
	if err != nil {
		log.Fatal("Error printing tree")
	}
	if this.numMetadata == 0 {
		fmt.Println("No metadata")
	} else {
		fmt.Println("Metadata:")
		for i, metadata := range this.metadata {
			_, err := fmt.Printf("%d", metadata)
			if err != nil {
				log.Fatal("Error printing tree metadata")
			}
			if i < this.numMetadata - 1 {
				fmt.Print(", ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("----------------")
	for _, child := range this.children {
		child.Print(level + 1)
	}
}

func parseEnd(tempNode *TreeNode) (parentNode *TreeNode, state ParseState, childCount int, metadataCount int) {
	// If we are in root
	if tempNode.parent == nil {
		state = PARSE_COMPLETED
	} else {
		// Otherwise we're in child,
		// traverse up to its parent
		// and restore child count
		// and metadata count
		tempNode = tempNode.parent
		childCount = tempNode.numChildren
		metadataCount = tempNode.numMetadata
		if tempNode.numChildrenRead >= childCount {
			state = PARSE_META_DATA
		} else {
			state = PARSE_CHILD_COUNT
		}
	}

	return tempNode, state, childCount, metadataCount
}

func parseTree(input string, debug bool) *TreeNode {
	inputArr := strings.Split(input, " ")
	nodeNumArr := make([]int, len(inputArr))
	for i, val := range inputArr {
		num, err := strconv.ParseInt(val, 10, 32)
		nodeNumArr[i] = int(num)
		if err != nil {
			log.Fatal("Error parsing input")
		}
	}

	if debug {
		fmt.Println("Numbers read:")
		for _, val := range nodeNumArr {
			fmt.Println(val)
		}
	}

	var root, tempNode *TreeNode
	state := PARSE_CHILD_COUNT

	childCount := 0
	metadataCount := 0
	numMetadataRead := 0

	length := len(nodeNumArr)
	if debug {
		fmt.Println("number => state value")
	}
	for i := 0; i < length; i++ {
		val := nodeNumArr[i]
		if debug {
			fmt.Printf("%d => %d\n", val, state)
		}
		switch state {
		case PARSE_CHILD_COUNT:
			childCount = val
			state = PARSE_META_COUNT
			break
		case PARSE_META_COUNT:
			metadataCount = val
			// If root node is nil, set new tree node
			// to be root. Otherwise, create new tree node
			// and append it to current node as a child
			if root == nil {
				root = createTreeNode(childCount, metadataCount)
				tempNode = root
			} else {
				newNode := createTreeNode(childCount, metadataCount)
				tempNode.AddChild(newNode)
				tempNode = newNode
			}
			if childCount > 0 {
				state = PARSE_CHILD_COUNT
			} else if metadataCount > 0 {
				state = PARSE_META_DATA
			} else {
				tempNode, state, childCount, metadataCount = parseEnd(tempNode)
				numMetadataRead = 0
			}
			break
		case PARSE_META_DATA:
			if numMetadataRead >= metadataCount {
				tempNode, state, childCount, metadataCount = parseEnd(tempNode)
				numMetadataRead = 0
				i -= 1
			} else {
				tempNode.AddMetadata(val)
				numMetadataRead += 1
			}
			break 
		}
	}

	return root
}

func main() {
	debugFlag := flag.Bool("debug", false, "Turn on debug options")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	var tree *TreeNode

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

		tree = parseTree(input, *debugFlag)
	}

	if (*debugFlag) {
		tree.Print(0)
	}

	pt1Answer := tree.GetSum()
	fmt.Println(pt1Answer)
}
