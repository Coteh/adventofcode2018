package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func getValue(valStr string) (int) {
	val, err := strconv.Atoi(valStr[1:len(valStr)])
	if err != nil {
		return 0
	}
	sign := valStr[0]
	if sign == '+' {
		return val
	} else if (sign == '-') {
		return -val
	} else {
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	totalVal := 0
	// TODO delete ended, I never used it
	ended := false
	for !ended {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				// TODO print to stderr
				fmt.Println("Encountered an error")
				os.Exit(1)
			}
			break;
		}
		// TODO no need to trim, just send the whole trimmed input
		splitStr := strings.Split(strings.Trim(input,"\n "), ",")
		for i := 0; i < len(splitStr); i++ {
			totalVal += getValue(strings.TrimLeft(splitStr[i], " "))
		}
	}
	fmt.Println(totalVal)
}
