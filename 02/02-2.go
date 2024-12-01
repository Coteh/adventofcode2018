package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func checkBoxIDs(first string, second string) []string {
	diffCount := 0

	for i, letter := range first {
		if i >= len(second) {
			return nil
		}
		if letter != rune(second[i]) {
			diffCount += 1
			if diffCount > 1 {
				return nil
			}
		}
	}

	commonLetters := make([]string, 0, len(first))

	for i, letter := range first {
		if letter != rune(second[i]) {
			continue
		}
		commonLetters = append(commonLetters, string(letter))
	}

	return commonLetters
}

func getCommonBoxLetters(ids []string) string {
	var result []string

	for i, id := range ids {
		for j, jid := range ids {
			if i != j {
				result = checkBoxIDs(id, jid)
				if result != nil {
					return strings.Join(result, "")
				}
			}
		}
	}

	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ids := make([]string, 0)

	for true {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, "Encountered an error")
				os.Exit(1)
			}
			break
		}
		ids = append(ids, strings.Trim(input, "\n "))
	}
	fmt.Println(getCommonBoxLetters(ids))
}
