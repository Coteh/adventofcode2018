package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getValue(valStr string) int {
	val, err := strconv.Atoi(valStr[1:len(valStr)])
	if err != nil {
		return 0
	}
	sign := valStr[0]
	if sign == '+' {
		return val
	} else if sign == '-' {
		return -val
	} else {
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	vals := make([]int, 0)
	totalVal := 0
	recorded := make(map[int]int)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, "Encountered an error")
				os.Exit(1)
			}
			break
		}
		//fmt.Println(input)
		vals = append(vals, getValue(strings.Trim(input, "\n ")))
		//fmt.Println(vals)
		if len(vals) > cap(vals) {
			newVals := make([]int, len(vals), (cap(vals)+1)*2)
			numCopied := copy(newVals, vals)
			if numCopied < len(vals) {
				fmt.Println("An error occurred")
				os.Exit(1)
			}
			vals = newVals
			fmt.Println(vals)
		}
	}
	for {
		for i := 0; i < len(vals); i++ {
			totalVal += vals[i]
			recorded[totalVal] = recorded[totalVal] + 1
			//fmt.Println(vals[i])
			if recorded[totalVal] >= 2 {
				fmt.Println(totalVal)
				os.Exit(0)
			}
		}
	}
	// fmt.Println("No duplicate frequencies found")
}
