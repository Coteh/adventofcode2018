package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strings"
)

type letterResult struct {
	hasExactlyTwo bool
	hasExactlyThree bool
}

func getLetterResults (id string) (result letterResult) {
	letterMap := make(map[int32]int)

	result = letterResult{false, false}

	for _, letter := range id {
		letterMap[letter] = letterMap[letter] + 1
	}

	for _, value := range letterMap {
		if value == 2 {
			result.hasExactlyTwo = true
		} else if value == 3 {
			result.hasExactlyThree = true
		}
	}

	return result
}

func getChecksum (ids []string) (int) {
	twos, threes := 0, 0

	for _, id := range ids {
		results := getLetterResults(id)
		if results.hasExactlyTwo {
			twos += 1
		}
		if results.hasExactlyThree {
			threes += 1
		}
	}

	return twos * threes
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ids := make([]string, 0)

	for true {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				// TODO print to stderr
				fmt.Println("Encountered an error")
				os.Exit(1)
			}
			break;
		}
		ids = append(ids, strings.Trim(input, "\n "))
	}
	fmt.Println(getChecksum(ids))
}