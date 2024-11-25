package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type PotRulesMap map[string]string
type PotGeneration []int

func parseInput() (initialGeneration PotGeneration, potRulesMap PotRulesMap) {
	scanner := bufio.NewScanner(os.Stdin)

	potRulesMap = make(PotRulesMap)
	
	for scanner.Scan() {
		line := scanner.Text()

		initialStatePrefix := "initial state: "
		if strings.HasPrefix(line, initialStatePrefix) {
			generationStr := line[len(initialStatePrefix):]
			for i, chr := range generationStr {
				if chr == '#' {
					initialGeneration = append(initialGeneration, i)
				}
			}
		} else if line != "" {
			rule := strings.Split(line, " => ")
			// No need to explicitly check the rules that don't produce a plant
			if rule[1] == "#" {
				potRulesMap[rule[0]] = rule[1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func generateNextGeneration(currGeneration PotGeneration, potRulesMap PotRulesMap) (nextGeneration PotGeneration) {
	start := currGeneration[0] - 3
	end := currGeneration[len(currGeneration) - 1] + 3

	potMap := make(map[int]bool)

	for i := 0; i < len(currGeneration); i++ {
		index := currGeneration[i]
		potMap[index] = true
	}

	for i := start; i <= end; i++ {
		var sectionBytes []byte
		if potMap[i - 2] {
			sectionBytes = append(sectionBytes, '#')
		} else {
			sectionBytes = append(sectionBytes, '.')
		}
		if potMap[i - 1] {
			sectionBytes = append(sectionBytes, '#')
		} else {
			sectionBytes = append(sectionBytes, '.')
		}
		if potMap[i] {
			sectionBytes = append(sectionBytes, '#')
		} else {
			sectionBytes = append(sectionBytes, '.')
		}
		if potMap[i + 1] {
			sectionBytes = append(sectionBytes, '#')
		} else {
			sectionBytes = append(sectionBytes, '.')
		}
		if potMap[i + 2] {
			sectionBytes = append(sectionBytes, '#')
		} else {
			sectionBytes = append(sectionBytes, '.')
		}
		section := string(sectionBytes)
		for rule := range potRulesMap {
			if rule == section {
				nextGeneration = append(nextGeneration, i)
			}
		}
	}

	return
}

func main() {
	debugFlag := flag.Bool("debug", false, "Turn on debug options")
	flag.Parse()

	initialGeneration, potRules := parseInput()

	if *debugFlag {
		log.Printf("initial generation: %v", initialGeneration)
		log.Printf("pot rules: %v", potRules)
	}

	var nextGeneration PotGeneration = initialGeneration
	const Part1MaxGenerations = 20
	const Part2MaxGenerations = 50000000000

	for i := 0; i < Part1MaxGenerations; i++ {
		nextGeneration = generateNextGeneration(nextGeneration, potRules)
	}

	var sum = 0
	for _, index := range nextGeneration {
		sum += index
	}
	fmt.Println(sum)

	// TODO: Finish part 2
	// for i := Part1MaxGenerations; i < Part2MaxGenerations; i++ {
	// 	nextGeneration = generateNextGeneration(nextGeneration, potRules)
	// 	fmt.Println(nextGeneration)
	// }

	// sum = 0
	// for _, index := range nextGeneration {
	// 	sum += index
	// }
	// fmt.Println(sum)
}
