package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	input, err := os.ReadFile(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := strings.Split(string(input), "\n")

	signal := []rune(data[0])
	for i := 0; i < len(signal); i++ {
		if distinctChars(signal, i, 4) > 0 {
			fmt.Println("start-of-packet marker: ", distinctChars(signal, i, 4))
			break
		}
	}

	for i := 0; i < len(signal); i++ {
		if distinctChars(signal, i, 14) > 0 {
			fmt.Println("start-of-message marker: ", distinctChars(signal, i, 14))
			break
		}
	}

}

func distinctChars(data []rune, index int, count int) int {
	m := make(map[rune]int)
	for i := index; i < index+count; i++ {
		// basic error handling for end of slice
		if i == len(data) {
			return -1
		}

		// if the value is already in the map, return 0
		if _, ok := m[data[i]]; ok {
			return 0
		}

		// add the value to the map
		m[data[i]] = 1
	}
	return index + count
}
