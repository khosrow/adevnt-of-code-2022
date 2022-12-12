package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func main() {
	args := os.Args
	input, err := os.ReadFile(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := strings.Split(string(input), "\n")

	// part 1
	head := Coord{0, 0}
	tail := Coord{0, 0}
	walk := make(map[Coord]int)

	for _, line := range data {
		if line == "" {
			continue
		}
		items := strings.Split(line, " ")
		instr := items[0]
		steps, err := strconv.Atoi(items[1])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for i := 0; i < steps; i++ {
			switch instr {
			case "R":
				head.x += 1
			case "L":
				head.x -= 1
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			}
			tail = moveTail(head, tail, instr)
			walk[tail] += 1

		}

	}

	fmt.Println("Part 1 - Positions visited:", len(walk))

	fmt.Println("----------------------------")
	// part 2
	rope := make([]Coord, 10)
	walk = make(map[Coord]int)

	for _, line := range data {
		if line == "" {
			continue
		}
		items := strings.Split(line, " ")
		instr := items[0]
		steps, err := strconv.Atoi(items[1])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// fmt.Println("== ", instr, steps, " ==")
		for i := 0; i < steps; i++ {
			switch instr {
			case "R":
				rope[0].x += 1
			case "L":
				rope[0].x -= 1
			case "U":
				rope[0].y += 1
			case "D":
				rope[0].y -= 1
			}
			for j := 1; j < len(rope); j++ {
				rope[j] = moveTail(rope[j-1], rope[j], instr)
				walk[rope[9]] += 1

			}
		}
		fmt.Println(rope)
	}
	// fmt.Println(rope)
	fmt.Println("Part 2 - Positions visited:", len(walk))
}

func moveTail(head Coord, tail Coord, dir string) Coord {
	if tail.y == head.y {
		if math.Abs(float64(tail.x-head.x)) == 2 {
			switch dir {
			case "L":
				tail.x -= 1
			case "R":
				tail.x += 1
			}
		}
	} else if tail.x == head.x {
		if math.Abs(float64(tail.y-head.y)) == 2 {
			switch dir {
			case "U":
				tail.y += 1
			case "D":
				tail.y -= 1
			}
		}
	} else if math.Abs(float64(tail.x-head.x)) == 1 && math.Abs(float64(tail.y-head.y)) == 1 {
		return tail
	} else {
		// Diagonal move
		slope := float64(tail.y-head.y) / float64(tail.x-head.x)
		// fmt.Println("After head:", head, "tail: ", tail)
		// fmt.Println("slope = ", slope)
		if slope >= 0 {
			switch dir {
			case "U", "R":
				tail.x += 1
				tail.y += 1
			case "D", "L":
				tail.x -= 1
				tail.y -= 1
			}
		} else {
			switch dir {
			case "U", "L":
				tail.x -= 1
				tail.y += 1
			case "D", "R":
				tail.x += 1
				tail.y -= 1
			}
		}
	}
	// fmt.Println("After head:", head, "tail: ", tail)

	return tail
}
