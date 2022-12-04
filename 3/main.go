package main

import (
	"fmt"
	"os"
	"strings"
	// "strconv"
	// "sort"
)

func main() {

	args := os.Args
	input, err := os.ReadFile(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sacks := strings.Split(string(input), "\n")

	var items []int
	for _, sack := range sacks {
		if sack == "" {
			continue
		}

		//split the sack into two parts
		part1 := sack[:len(sack)/2]
		part2 := sack[len(sack)/2:]

		out:
		for _, c := range part1 {
			for _, d := range part2 {
				if c == d {
					items = append(items, getPriority(int(c)))
					break out
				}
			}
		}
		// fmt.Println(part1, " ", part2)
	}

	// fmt.Println(items)
	fmt.Println("item sum: ", getSum(items))

	var badges []int
	// iterate all scaks in groups of three
	for i := 0; i < len(sacks); i += 3 {
		outer:
		for _, s1 := range sacks[i] {
			for _, s2 := range sacks[i+1] {
				if s1 == s2 {
					for _, s3 := range sacks[i+2] {
						if s1 == s3 {
							badges = append(badges, getPriority(int(s1)))
							break outer
						}
					}
				}
			}
		}
	}
	fmt.Println("badge sum: ", getSum(badges))
}

func getPriority(c int) int {
	priority := 0
	if c >= 65 && c <= 90 {
		priority = c - 64 + 26
	} else {
		priority = c - 96
	}
	return priority
}

func getSum(items []int) int {
	sum := 0
	for _, item := range items {
		sum += item
	}
	return sum
}
