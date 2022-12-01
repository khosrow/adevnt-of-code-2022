package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {

	args := os.Args
	input, err := os.ReadFile(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Part 1
	calories := strings.Split(string(input), "\n")
	var elves []int

	elf := 0
	for _, calorie := range calories {
		if calorie == "" {
			elves = append(elves, elf)
			elf = 0
		} else {
			c, _ := strconv.Atoi(calorie)
			elf += c
		}
	}

	sort.Sort(sort.IntSlice(elves))
	fmt.Println("Highest:", elves[len(elves)-1])
	fmt.Println("Top Three:", elves[len(elves)-1], elves[len(elves)-2], elves[len(elves)-3])
	fmt.Println("Top Three Total:", elves[len(elves)-1]+elves[len(elves)-2]+elves[len(elves)-3])
}


