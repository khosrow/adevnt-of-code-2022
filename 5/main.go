package main

import (
	"fmt"
	"os"
	"strconv"
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

	stack_count := 0
	j := 0
	var tokens []string

	for i, d := range data {
		if d == "" {
			continue
		}

		tokens = strings.Split(strings.TrimLeft(d, " "), "   ")

		// check where the columns start and get column count
		_, err := strconv.Atoi(tokens[0])
		if err == nil {
			stack_count = len(tokens)
			j = i
			break
		}
	}
	fmt.Println("Stack count:", stack_count)

	var stacks, stacks2 [][]string
	stacks = make([][]string, stack_count)
	stacks2 = make([][]string, stack_count)

	// process crates
	for i := j - 1; i >= 0; i -= 1 {
		for s, c := 0, 1; c < len(data[i]); s, c = s+1, c+4 {
			if string(data[i][c]) == " " {
				continue
			}

			stacks[s] = append(stacks[s], string(data[i][c]))
		}
	}

	copy(stacks2, stacks)

	// process moves
	for i := j + 1; i < len(data); i += 1 {
		if data[i] == "" {
			continue
		}
		tokens = strings.Split(data[i], " ")
		count, _ := strconv.Atoi(tokens[1])
		source, _ := strconv.Atoi(tokens[3])
		dest, _ := strconv.Atoi(tokens[5])

		stacks = move(stacks, count, source, dest)
		stacks2 = moveMany(stacks2, count, source, dest)
	}

	fmt.Println("CrateMover 9000 Top:", stackTops(stacks))
	fmt.Println("CrateMover 9001 Top:", stackTops(stacks2))
}

func stackTops(stacks [][]string) string {
	top := ""
	for _, stack := range stacks {
		top += stack[len(stack)-1]
	}
	return top
}

func move(stacks [][]string, count int, source int, dest int) [][]string {
	for i := 0; i < count; i++ {
		crate, s := slicePop(stacks[source-1])
		stacks[source-1] = s
		stacks[dest-1] = append(stacks[dest-1], crate)
	}
	return stacks
}

func moveMany(stacks [][]string, count int, source int, dest int) [][]string {
	fmt.Println("move", count, "from", source, "to", dest)
	crates, s := sliceCut(stacks[source-1], count)
	fmt.Println("Crates: ", crates, "Source: ", s)
	fmt.Println("foo")
	stacks[source-1] = s
	stacks[dest-1] = append(stacks[dest-1], crates...)
	return stacks
}

func sliceCut(s []string, i int) ([]string, []string) {
	if len(s) == i {
		return s, []string{}
	} else {
		return s[len(s)-i:], s[:len(s)-i]
	}

}

func slicePop(s []string) (string, []string) {
	if len(s) == 1 {
		return s[0], []string{}
	} else {
		return s[len(s)-1], s[:len(s)-1]
	}
}
