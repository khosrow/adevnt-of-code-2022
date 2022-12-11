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

	// populate trees
	trees := make([][]int, len(data)-1, len(data)-1)

	for i := 0; i < len(data); i++ {
		line := data[i]
		if line == "" {
			continue
		}
		for j := 0; j < len(line); j++ {
			tree, err := strconv.Atoi(string(line[j]))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			trees[i] = append(trees[i], tree)
		}
	}

	// part 1
	visibleTrees := (len(trees) - 1) * 4 // all trees are visible from the top and left
	fmt.Println("Dimensions: ", len(trees), "x", len(trees[0]))
	fmt.Println("Trees on the edge: ", visibleTrees)

	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			if isVisible(trees, i, j) {
				visibleTrees += 1
			}
		}
	}
	fmt.Println("visible trees: ", visibleTrees)

	// part 2
	score := 0
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			s := scenicScore(trees, i, j)
			if s > score {
				score = s
			}
		}
	}
	fmt.Println("scenic score: ", score)
}

func isVisible(trees [][]int, row int, col int) bool {
	visible := false
	for j := 0; j < col; j++ {
		if trees[row][j] < trees[row][col] {
			visible = true
		} else {
			visible = false
			break
		}
	}
	if visible {
		return visible
	}

	for j := col + 1; j < len(trees[row]); j++ {
		if trees[row][j] < trees[row][col] {
			visible = true
		} else {
			visible = false
			break
		}
	}
	if visible {
		return visible
	}

	for i := 0; i < row; i++ {
		if trees[i][col] < trees[row][col] {
			visible = true
		} else {
			visible = false
			break
		}
	}
	if visible {
		return visible
	}

	for i := row + 1; i < len(trees); i++ {
		if trees[i][col] < trees[row][col] {
			visible = true
		} else {
			visible = false
			break
		}
	}
	return visible
}

func scenicScore(trees [][]int, row int, col int) int {
	left, right, up, down := 0, 0, 0, 0

	for j := col - 1; j >= 0; j-- {
		left += 1
		if trees[row][j] >= trees[row][col] {
			break
		}
	}

	for j := col + 1; j < len(trees[row]); j++ {
		right += 1
		if trees[row][j] >= trees[row][col] {
			break
		}
	}

	for i := row - 1; i >= 0; i-- {
		up += 1
		if trees[i][col] >= trees[row][col] {
			break
		}
	}

	for i := row + 1; i < len(trees); i++ {
		down += 1
		if trees[i][col] >= trees[row][col] {
			break
		}
	}

	return left * right * up * down
}
