package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Register [4]int

type Input struct {
	BeforeRegister Register
	AfterRegister Register
	Instructions [4]int
}

// add register
func addr(reg *Register, aReg int, bReg int, cReg int) {
	reg[cReg] = reg[aReg] + reg[bReg]
}

// add immediate
func addi(reg *Register, aReg int, bVal int, cReg int) {
	reg[cReg] = reg[aReg] + bVal
}

// multiply register
func mulr(reg *Register, aReg int, bReg int, cReg int) {
	reg[cReg] = reg[aReg] * reg[bReg]
}

// multiply immediate
func muli(reg *Register, aReg int, bVal int, cReg int) {
	reg[cReg] = reg[aReg] * bVal
}

// bitwise AND register
func banr(reg *Register, aReg int, bReg int, cReg int) {
	reg[cReg] = reg[aReg] & reg[bReg]
}

// bitwise AND immediate
func bani(reg *Register, aReg int, bVal int, cReg int) {
	reg[cReg] = reg[aReg] & bVal
}

// bitwise OR register
func borr(reg *Register, aReg int, bReg int, cReg int) {
	reg[cReg] = reg[aReg] | reg[bReg]
}

// bitwise OR immediate
func bori(reg *Register, aReg int, bVal int, cReg int) {
	reg[cReg] = reg[aReg] | bVal
}

// set register
func setr(reg *Register, aReg int, cReg int) {
	reg[cReg] = reg[aReg]
}

// set immediate
func seti(reg *Register, aVal int, cReg int) {
	reg[cReg] = aVal
}

// greater-than immediate/register
func gtir(reg *Register, aVal int, bReg int, cReg int) {
	if aVal > reg[bReg] {
		reg[cReg] = 1
	} else {
		reg[cReg] = 0
	}
}

// greater-than register/immediate
func gtri(reg *Register, aReg int, bVal int, cReg int) {
	if reg[aReg] > bVal {
		reg[cReg] = 1
	} else {
		reg[cReg] = 0
	}
}

// greater-than register/register
func gtrr(reg *Register, aReg int, bReg int, cReg int) {
	if reg[aReg] > reg[bReg] {
		reg[cReg] = 1
	} else {
		reg[cReg] = 0
	}
}

// equal immediate/register
func eqir(reg *Register, aVal int, bReg int, cReg int) {
	if aVal == reg[bReg] {
		reg[cReg] = 1
	} else {
		reg[cReg] = 0
	}
}

// equal register/immediate
func eqri(reg *Register, aReg int, bVal int, cReg int) {
	if reg[aReg] == bVal {
		reg[cReg] = 1
	} else {
		reg[cReg] = 0
	}
}

// equal register/register
func eqrr(reg *Register, aReg int, bReg int, cReg int) {
	if reg[aReg] == reg[bReg] {
		reg[cReg] = 1
	} else {
		reg[cReg] = 0
	}
}

func compareRegisters(reg1 Register, reg2 Register) bool {
	if len(reg1) != len(reg2) {
		return false
	}
	for i, reg1Val := range reg1 {
		reg2Val := reg2[i]
		if reg1Val != reg2Val {
			return false
		}
	}
	return true
}

func determineNumberOfValidInstructions(input Input) (numResults int) {
	instructions := input.Instructions

	for i := 0; i < 16; i++ {
		var reg = input.BeforeRegister
		switch (i) {
		case 0:
			addr(&reg, instructions[1], instructions[2], instructions[3])
		case 1:
			addi(&reg, instructions[1], instructions[2], instructions[3])
		case 2:
			mulr(&reg, instructions[1], instructions[2], instructions[3])
		case 3:
			muli(&reg, instructions[1], instructions[2], instructions[3])
		case 4:
			banr(&reg, instructions[1], instructions[2], instructions[3])
		case 5:
			bani(&reg, instructions[1], instructions[2], instructions[3])
		case 6:
			borr(&reg, instructions[1], instructions[2], instructions[3])
		case 7:
			bori(&reg, instructions[1], instructions[2], instructions[3])
		case 8:
			setr(&reg, instructions[1], instructions[3])
		case 9:
			seti(&reg, instructions[1], instructions[3])
		case 10:
			gtir(&reg, instructions[1], instructions[2], instructions[3])
		case 11:
			gtri(&reg, instructions[1], instructions[2], instructions[3])
		case 12:
			gtrr(&reg, instructions[1], instructions[2], instructions[3])
		case 13:
			eqir(&reg, instructions[1], instructions[2], instructions[3])
		case 14:
			eqri(&reg, instructions[1], instructions[2], instructions[3])
		case 15:
			eqrr(&reg, instructions[1], instructions[2], instructions[3])
		}
		if compareRegisters(reg, input.AfterRegister) {
			numResults++
		}
	}

	return
}

const (
	SCAN_MODE_NONE = ""
	SCAN_MODE_PART1 = "part1_input"
	SCAN_MODE_PART2 = "part2_input"
)

func parseInputArray(line string) ([4]int) {
	var myArr [4]int
	err := json.Unmarshal([]byte(line), &myArr)
	if err != nil {
		panic(err)
	}
	return myArr
}

func parseInput() (inputSet []Input) {
	scanner := bufio.NewScanner(os.Stdin)

	currMode := SCAN_MODE_NONE

	var currBefore Register
	var currAfter Register
	var currInstructions [4]int

	for scanner.Scan() {
		line := scanner.Text()

		switch currMode {
		case SCAN_MODE_PART1:
			var arrStr string
			if strings.Contains(line, "After:  ") {
				arrStr = line[len("After:  "):]
			} else if line != "" {
				modifiedLine := strings.Replace(line, " ", ",", -1)
				arrStr = fmt.Sprintf("[%s]", modifiedLine)
			} else {
				inputSet = append(inputSet, Input{
					BeforeRegister: currBefore,
					AfterRegister: currAfter,
					Instructions: currInstructions,
				})
				currMode = SCAN_MODE_NONE
				break
			}
			myArr := parseInputArray(arrStr)
			if strings.Contains(line, "After:  ") {
				currAfter = myArr
			} else {
				currInstructions = myArr
			}
		case SCAN_MODE_PART2:
			// ignore part 2 for now
			// fmt.Printf("Ignoring '%s' for part 1\n", line)
		case SCAN_MODE_NONE:
			if strings.Contains(line, "Before: ") {
				currMode = SCAN_MODE_PART1
				arrStr := line[len("Before: "):]
				myArr := parseInputArray(arrStr)
				currBefore = myArr
			} else {
				currMode = SCAN_MODE_PART2
			}
		}
	}

	if currMode == SCAN_MODE_PART1 {
		inputSet = append(inputSet, Input{
			BeforeRegister: currBefore,
			AfterRegister: currAfter,
			Instructions: currInstructions,
		})
		currMode = SCAN_MODE_NONE
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// TODO: Also parse the instructions and make a second return value for those

	return
}

func main() {
	inputSet := parseInput()

	part1Answer := 0
	for _, input := range inputSet {
		numResults := determineNumberOfValidInstructions(input)
		if numResults >= 3 {
			part1Answer += 1
		}
	}
	fmt.Println(part1Answer)
}
