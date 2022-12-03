package main

import (
	"fmt"
	"os"
	"strings"
	// "strconv"
	// "sort"
)

func main() {

	AX, AY, AZ, BX, BY, BZ, CX, CY, CZ := 0, 0, 0, 0, 0, 0, 0, 0, 0

	args := os.Args
	input, err := os.ReadFile(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	moves := strings.Split(string(input), "\n")
	fmt.Println(moves)
	for _, move := range moves {
		if move == "" {
			continue
		}

		fmt.Println(move)
		switch move {
		case "A X":
			AX += 4
		case "A Y":
			AY += 8
		case "A Z":
			AZ += 3
		case "B X":
			BX += 1
		case "B Y":
			BY += 5
		case "B Z":
			BZ += 9
		case "C X":
			CX += 7
		case "C Y":
			CY += 2
		case "C Z":
			CZ += 6
		}
	}
	fmt.Printf("AX: %d, AY: %d, AZ: %d BX: %d, BY: %d, BZ: %d, CX: %d, CY: %d, CZ: %d\n",
	AX, AY, AZ, BX, BY, BZ, CX, CY, CZ)

	score := AX + AY + AZ + BX + BY + BZ + CX + CY + CZ
	fmt.Println("score: ", score)

	score = 0
	for _, move := range moves {
		if move == "" {
			continue
		}
		cols := strings.Split(move, " ")
		fmt.Println(cols)
		score += getScore(cols[0], cols[1])
	}
	fmt.Println("score: ", score)

}

func getScore(opponent, outcome string) int {
	var score int
	switch opponent {
	case "A":
		switch outcome {
		case "X":
			score = 3
		case "Y":
			score = 4
		case "Z":
			score = 8
		}
	case "B":
		switch outcome {
		case "X":
			score = 1
		case "Y":
			score = 5
		case "Z":
			score = 9
		}
	case "C":
		switch outcome {
		case "X":
			score = 2
		case "Y":
			score = 6
		case "Z":
			score = 7
		}
	}
	return score
}
