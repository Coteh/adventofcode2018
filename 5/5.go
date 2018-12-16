package main

import (
	"fmt"
	"bufio"
	"sync"
	"regexp"
	"strings"
	"os"
	"log"
	"io"
)

const NumChunks = 8

type PolymerChunkArray struct {
	sync.RWMutex
	values []string
	numChunks int
}

func processChunk(polymerChunk string, chunkIndex int, chunkResults []string, re *regexp.Regexp, wg *sync.WaitGroup) {
	defer wg.Done()

	chunkResults[chunkIndex] = re.ReplaceAllString(polymerChunk, "")
}

func (this *PolymerChunkArray) ProcessChunks(re *regexp.Regexp) string {
	var resultBuilder strings.Builder
	chunkResults := make([]string, this.numChunks)

	var wg sync.WaitGroup

	for i, chunk := range this.values {
		wg.Add(1)
		go processChunk(chunk, i, chunkResults, re, &wg)
	}

	wg.Wait()

	for i := 0; i < this.numChunks; i++ {
		resultBuilder.WriteString(chunkResults[i])
	}

	return resultBuilder.String()
}

func createChunkedArray(input string, numChunks int) *PolymerChunkArray {
	if numChunks <= 0 {
		numChunks = NumChunks
	}
	inputLength := len(input)
	chunkSize := inputLength / numChunks
	remainder := inputLength % numChunks

	chunkSizes := make([]int, numChunks)
	for i := 0; i < numChunks; i++ {
		chunkSizes[i] = chunkSize
	}
	if remainder > 0 {
		for i := 0; i < remainder; i++ {
			chunkSizes[i] += 1
		}
	}

	values := make([]string, numChunks)
	valueIndex := 0
	for i := 0; i < numChunks; i++ {
		values[i] = input[valueIndex:valueIndex + chunkSizes[i]]
		valueIndex += chunkSizes[i]
	}

	return &PolymerChunkArray {
		values: values,
		numChunks: numChunks,
	}
}

func generateRegex() *regexp.Regexp {
	var regexStrBuilder strings.Builder

	for i := 0; i < 26; i++ {
		regexStrBuilder.WriteString(string('A' + i))
		regexStrBuilder.WriteString(string('a' + i))
		regexStrBuilder.WriteString("|")
		regexStrBuilder.WriteString(string('a' + i))
		regexStrBuilder.WriteString(string('A' + i))
		if i < 25 {
			regexStrBuilder.WriteString("|")
		}
	}

	re, err := regexp.Compile(regexStrBuilder.String())
	if err != nil {
		return nil
	}

	return re
}

func testChunks(chunkedArr *PolymerChunkArray, input string) {
	var testBuilder strings.Builder
	
	for i := 0; i < chunkedArr.numChunks; i++ {
		testBuilder.WriteString(chunkedArr.values[i])
	}

	if testBuilder.String() == input {
		fmt.Println("✓ The chunked pieces combine to equal the original input")
	} else {
		log.Fatal("✗ The chunked pieces miss information from original input")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		if err !=  io.EOF {
			log.Fatal("Encountered an error with input")
			os.Exit(1)
		}
	}
	input = strings.TrimRight(input, "\n")

	re := generateRegex()

	// Uncomment to test
	// testChunks(chunkedArr, input)

	prevInput := input
	numChunks := NumChunks

	for true {
		var chunkedArr *PolymerChunkArray

		chunkedArr = createChunkedArray(prevInput, numChunks)

		result := chunkedArr.ProcessChunks(re)

		if result == prevInput {
			if numChunks == 1 {
				// No further actions can be taken
				break
			}
			// Reduce number of chunks
			numChunks -= 1
		}

		prevInput = result
	}

	fmt.Println(len(prevInput))
}
