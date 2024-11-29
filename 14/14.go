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

func parseInput() (numRecipes int) {
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
	flag.Parse()

	numRecipes := parseInput()
	if *debugFlag {
		fmt.Println("numRecipes", numRecipes)
	}

	scoreStr := iterateReceipesRight([]int{3, 7}, []int{0, 1}, numRecipes)
	fmt.Println(scoreStr)
}
