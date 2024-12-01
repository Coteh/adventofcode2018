package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"io"
	"strconv"
)

type Claim struct {
	id int
	left int
	top int
	width int
	height int
}

func parseClaim(input string, claimChan chan Claim) {
	if input == "" {
		return
	}

	re, err := regexp.Compile("#([0-9]+)\\s@\\s([0-9]+),([0-9]+):\\s([0-9]+)x([0-9]+)");
	if err != nil {
		return
	}

	matches := re.FindStringSubmatch(input)

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		return
	}
	left, err := strconv.Atoi(matches[2])
	if err != nil {
		return
	}
	top, err := strconv.Atoi(matches[3])
	if err != nil {
		return
	}
	width, err := strconv.Atoi(matches[4])
	if err != nil {
		return
	}
	height, err := strconv.Atoi(matches[5])
	if err != nil {
		return
	}

	claimObj := Claim{id, left, top, width, height}

	claimChan <- claimObj
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	claimChan := make(chan Claim)
	claims := make([]Claim, 0, 1000)
	numClaims := 0
	numClaimsRead := 0
	fabric := make([][]int, 1000)
	for i := range fabric {
		fabric[i] = make([]int, 1000)
	}

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
		go parseClaim(input, claimChan)
		numClaims += 1
	}

	for claim := range claimChan {
		for i := claim.left; i < claim.left + claim.width; i++ {
			for j := claim.top; j < claim.top + claim.height; j++ {
				fabric[i][j] = fabric[i][j] + 1
			}
		}
		// Append to claims array for part 2
		claims = append(claims, claim)
		// Go actually does something like this automatically
		// when you append one more element than slice capacity
		// if len(claims) > cap(claims) {
		// 	newClaims := make([]Claim, len(claims), (cap(claims) + 1) * 2)
		// 	numCopied := copy(newClaims, claims)
		// 	if numCopied < len(claims) {
		// 		fmt.Println("An error occurred")
		// 		os.Exit(1)
		// 	}
		// 	claims = newClaims
		// }
		numClaimsRead += 1
		if numClaimsRead == numClaims {
			break
		}
	}

	answer := 0
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric[i]); j++ {
			if fabric[i][j] > 1 {
				answer += 1
			}
		}
	}

	fmt.Println(answer)
	
	for _, claim := range claims {
		doesOverlap := false
		for i := claim.left; i < claim.left + claim.width; i++ {
			for j := claim.top; j < claim.top + claim.height; j++ {
				 if fabric[i][j] > 1 {
					doesOverlap = true
					break
				 }
			}
			if doesOverlap {
				break
			}
		}
		if !doesOverlap {
			fmt.Println(claim.id)
		}
	}
}