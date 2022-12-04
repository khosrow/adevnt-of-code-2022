package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// "sort"
)

func main() {

	args := os.Args
	input, err := os.ReadFile(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pairs := strings.Split(string(input), "\n")

	total := 0
	overlap := 0
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		elves := strings.Split(pair, ",")
		elf1 := convertToInt(strings.Split(elves[0], "-"))
		elf2 := convertToInt(strings.Split(elves[1], "-"))

		if contains(elf1, elf2) {
			total += 1
		}

		if overlaps(elf1, elf2) {
			overlap += 1
		}
	}
	fmt.Println("Fully cotained pairs:", total)
	fmt.Println("Overlapping pairs:", overlap)

}

func convertToInt(elf []string) []int {
	var converted []int
	for _, num := range elf {
		i, _ := strconv.Atoi(num)
		converted = append(converted, i)
	}
	return converted
}

func contains(elf1 []int, elf2 []int) bool {
	var large, small []int
	if (elf1[1] - elf1[0]) > (elf2[1] - elf2[0]) {
		large = elf1
		small = elf2
	} else {
		large = elf2
		small = elf1
	}
	if large[0] <= small[0] && large[1] >= small[1] {
		return true
	} else {
		return false
	}
}

func overlaps(elf1 []int, elf2 []int) bool {
	if elf1[0] <= elf2[0] {
		if elf1[1] >= elf2[0] {
			return true
		}
	} else { // elf2[0] < elf1[0]
		if elf2[1] >= elf1[0] {
			return true
		}
	}
	return false
}
