package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DEBUG_ON bool = false

type Register [6]int
type Instruction struct {
	Opcode    string
	Arguments [3]int
}

type Device struct {
	InstructionPointer int
	Register           Register
	InstructionSet     []Instruction
}

const (
	OPCODE_addr string = "addr"
	OPCODE_addi string = "addi"
	OPCODE_mulr string = "mulr"
	OPCODE_muli string = "muli"
	OPCODE_banr string = "banr"
	OPCODE_bani string = "bani"
	OPCODE_borr string = "borr"
	OPCODE_bori string = "bori"
	OPCODE_setr string = "setr"
	OPCODE_seti string = "seti"
	OPCODE_gtir string = "gtir"
	OPCODE_gtri string = "gtri"
	OPCODE_gtrr string = "gtrr"
	OPCODE_eqir string = "eqir"
	OPCODE_eqri string = "eqri"
	OPCODE_eqrr string = "eqrr"
)

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

func handleInstruction(reg *Register, opcode string, args [3]int) {
	switch opcode {
	case OPCODE_addr:
		addr(reg, args[0], args[1], args[2])
	case OPCODE_addi:
		addi(reg, args[0], args[1], args[2])
	case OPCODE_mulr:
		mulr(reg, args[0], args[1], args[2])
	case OPCODE_muli:
		muli(reg, args[0], args[1], args[2])
	case OPCODE_banr:
		banr(reg, args[0], args[1], args[2])
	case OPCODE_bani:
		bani(reg, args[0], args[1], args[2])
	case OPCODE_borr:
		borr(reg, args[0], args[1], args[2])
	case OPCODE_bori:
		bori(reg, args[0], args[1], args[2])
	case OPCODE_setr:
		setr(reg, args[0], args[2])
	case OPCODE_seti:
		seti(reg, args[0], args[2])
	case OPCODE_gtir:
		gtir(reg, args[0], args[1], args[2])
	case OPCODE_gtri:
		gtri(reg, args[0], args[1], args[2])
	case OPCODE_gtrr:
		gtrr(reg, args[0], args[1], args[2])
	case OPCODE_eqir:
		eqir(reg, args[0], args[1], args[2])
	case OPCODE_eqri:
		eqri(reg, args[0], args[1], args[2])
	case OPCODE_eqrr:
		eqrr(reg, args[0], args[1], args[2])
	}
}

func processInstructions(device *Device) (answer int) {
	for device.Register[device.InstructionPointer] < len(device.InstructionSet) {
		instructionVal := device.Register[device.InstructionPointer]
		instruction := device.InstructionSet[instructionVal]
		if DEBUG_ON {
			fmt.Printf("Handling instruction %d: %s %v\n", instructionVal, instruction.Opcode, instruction.Arguments)
		}
		handleInstruction(&device.Register, instruction.Opcode, instruction.Arguments)
		if DEBUG_ON {
			fmt.Printf("Instruction pointer after this instruction is: %d\n", device.Register[device.InstructionPointer])
		}
		device.Register[device.InstructionPointer]++
		if DEBUG_ON {
			fmt.Printf("Instruction pointer after incrementing the pointer is: %d\n", device.Register[device.InstructionPointer])
		}
	}
	return device.Register[0]
}

func parseInput() (device Device) {
	scanner := bufio.NewScanner(os.Stdin)

	device = Device{}

	for scanner.Scan() {
		line := scanner.Text()

		splitStr := strings.Split(line, " ")

		if splitStr[0] == "#ip" {
			num, err := strconv.Atoi(splitStr[1])
			if err != nil {
				panic(err)
			}
			device.InstructionPointer = num
			continue
		}

		opcode := splitStr[0]
		args := [3]int{}
		for i := 1; i < len(splitStr); i++ {
			num, err := strconv.Atoi(splitStr[i])
			if err != nil {
				panic(err)
			}
			args[i-1] = num
		}
		instruction := Instruction{
			Opcode:    opcode,
			Arguments: args,
		}
		device.InstructionSet = append(device.InstructionSet, instruction)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	device := parseInput()
	if DEBUG_ON {
		fmt.Println(device)
	}
	answer := processInstructions(&device)
	fmt.Println(answer)
}
