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
	totalVal := 0
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, "Encountered an error")
				os.Exit(1)
			}
			break
		}
		// TODO no need to trim, just send the whole trimmed input
		splitStr := strings.Split(strings.Trim(input, "\n "), ",")
		for i := 0; i < len(splitStr); i++ {
			totalVal += getValue(strings.TrimLeft(splitStr[i], " "))
		}
	}
	fmt.Println(totalVal)
}
