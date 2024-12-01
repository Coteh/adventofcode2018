package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var debugFlag *bool

func parseNumRecipesInput() (numRecipes int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		var err error
		numRecipes, err = strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func parseSequenceInput() (sequence string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		sequence = line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func iterateReceipesRight(recipeArr []int, startingPos []int, checkAfter int) (score string) {
	if len(recipeArr) == 1 || len(startingPos) == 1 {
		fmt.Fprintf(os.Stderr, "Invalid params")
		return
	}

	currPos := startingPos
	for {
		if *debugFlag {
			printRecipeState(recipeArr, currPos)
		}
		sum := 0
		for _, pos := range currPos {
			sum += recipeArr[pos]
		}
		sumStr := strconv.FormatInt(int64(sum), 10)
		arr := strings.Split(sumStr, "")
		for _, val := range arr {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			recipeArr = append(recipeArr, num)
		}
		if len(recipeArr) >= checkAfter+10 {
			break
		}
		for i, pos := range currPos {
			diff := 1 + recipeArr[pos]
			currPos[i] = (currPos[i] + diff) % len(recipeArr)
		}
	}

	var scoreStr []byte
	for _, val := range recipeArr[checkAfter : checkAfter+10] {
		scoreStr = append(scoreStr, byte(val)+48)
	}
	if *debugFlag {
		fmt.Println("scoreStr", scoreStr, string(scoreStr))
	}

	return string(scoreStr)
}

func iterateReceipesLeft(recipeArr []int, startingPos []int, sequence string) (numRecipes int) {
	if len(recipeArr) == 1 || len(startingPos) == 1 {
		fmt.Fprintf(os.Stderr, "Invalid params")
		return
	}

	sequenceNums := make([]int, 0, len(sequence))
	for _, val := range sequence {
		sequenceNums = append(sequenceNums, int(val - '0'))
	}
	if *debugFlag {
		fmt.Println(sequenceNums)
	}

	currPos := startingPos
mainLoop:
	for {
		if *debugFlag {
			printRecipeState(recipeArr, currPos)
		}
		sum := 0
		for _, pos := range currPos {
			sum += recipeArr[pos]
		}
		sumStr := strconv.FormatInt(int64(sum), 10)
		arr := strings.Split(sumStr, "")
		for _, val := range arr {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			recipeArr = append(recipeArr, num)
		}
		// Realized after reading this that the last digit of the input sequence can be missed if the last recipe generated was two digits.
		// https://www.reddit.com/r/adventofcode/comments/a671s8/comment/ebskkoy/
		// To fix this edge case, the starting point of the check needs to start not just from the last digit, but also from the digit before that if 2 digits were generated.
		// This loop will perform a check from each valid starting point (it should only ever be from 1 or 2 digits from the end of scoreboard)
		for k := 1; k <= len(arr); k++ {
			start := len(recipeArr) - k
			j := len(sequence) - 1
			for i := start; i >= 0 && j >= 0; i-- {
				if recipeArr[i] == sequenceNums[j] {
					j--
				} else {
					break
				}
			}
			// If sequence was detected, then count number of digits (recipes) before the sequence and return that
			if (j < 0) {
				numRecipes = len(recipeArr) - k + 1 - len(sequence)
				break mainLoop
			}
		}
		for i, pos := range currPos {
			diff := 1 + recipeArr[pos]
			currPos[i] = (currPos[i] + diff) % len(recipeArr)
		}
	}

	return
}

func printRecipeState(recipeArr []int, currPos []int) {
	for i, recipe := range recipeArr {
		if currPos[0] == i {
			fmt.Printf("(%d)", recipe)
		} else if currPos[1] == i {
			fmt.Printf("[%d]", recipe)
		} else {
			fmt.Printf(" %d ", recipe)
		}
	}
	fmt.Printf("\n")
}

func main() {
	debugFlag = flag.Bool("debug", false, "Turn on debug options")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] <left | right>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	
	direction := flag.Arg(0)
	
	switch (direction) {
		// Part 1
		case "right":
			numRecipes := parseNumRecipesInput()
			if *debugFlag {
				fmt.Println("numRecipes", numRecipes)
			}
			scoreStr := iterateReceipesRight([]int{3, 7}, []int{0, 1}, numRecipes)
			fmt.Println(scoreStr)
		// Part 2
		case "left":
			sequence := parseSequenceInput()
			if *debugFlag {
				fmt.Println("sequence", sequence)
			}
			numRecipes := iterateReceipesLeft([]int{3, 7}, []int{0, 1}, sequence)
			fmt.Println(numRecipes)
		default:
			fmt.Fprintln(os.Stderr, "Please provide direction (\"left\" or \"right\")")
			os.Exit(1)
	}
}
