package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	kind     string // "directory", "file"
	size     int
	name     string
	parent   *File
	children []File
}

func main() {
	args := os.Args
	input, err := os.ReadFile(args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := strings.Split(string(input), "\n")

	// setup root directory
	var root File
	root.kind = "directory"
	root.name = "/"
	root.parent = nil
	root.children = make([]File, 0)
	current := &root

	for i := 0; i < len(data); i++ {
		if data[i] == "" {
			continue
		}

		line := data[i]

		// process commands
		if line[0] == '$' {
			command := strings.Split(line, " ")
			switch command[1] {
			case "cd":
				switch command[2] {
				case "/": // special case for root
					continue
				case "..": // go up one level
					current = current.parent
				default: // enter a directory
					for j := 0; j < len(current.children); j++ {
						if current.children[j].name == command[2] {
							current = &current.children[j]
							break
						}
					}
				}
			case "ls": // don't care here, we need the output
				continue
			}
		} else { // process files
			entry := strings.Split(line, " ")
			// entry[0] - filesize
			// entry[1] - filename
			if entry[0] == "dir" {
				addFile(current, entry[1], 0)
			} else {
				size, err := strconv.Atoi(entry[0])
				if err != nil {
					fmt.Println("ERROR", err)
					continue
				}
				addFile(current, entry[1], size)
			}
		}
	}

	// part 1
	disk_used := getSize(&root)

	dir_sizes := getSmallDirs(&root, 100000)
	total := 0
	for _, size := range dir_sizes {
		total += size
	}
	fmt.Println("At most 100KB :", total)

	// part 2
	const DISK_TOTAL = 70000000
	const DISK_NEEDED = 30000000
	disk_to_clean := -1 * (DISK_TOTAL - disk_used - DISK_NEEDED)
	fmt.Println("Disk used     :", disk_used)
	fmt.Println("Disk available:", DISK_TOTAL-disk_used)
	fmt.Println("Disk to clean :", disk_to_clean)

	dir_sizes = getBigDirs(&root, disk_to_clean)

	fmt.Println("Smallest to delete:", getSmallestToRemove(dir_sizes, disk_to_clean))
}

func addFile(root *File, path string, size int) {
	var newFile File
	newFile.name = path
	newFile.parent = root
	if size == 0 {
		newFile.children = make([]File, 0)
		newFile.kind = "directory"

	} else {
		newFile.children = nil
		newFile.kind = "file"
	}

	newFile.size = size
	root.children = append(root.children, newFile)
}

func getSize(f *File) int {
	if f.kind == "file" {
		return f.size
	}
	size := 0
	for i := 0; i < len(f.children); i++ {
		size += getSize(&f.children[i])
	}
	f.size = size
	return size
}

// go through dirs and return the pointer to each small dir
func getSmallDirs(f *File, max int) []int {
	var size []int
	if f.kind == "file" {
		return size
	}
	if f.size < max {
		size = append(size, f.size)
	}
	for i := 0; i < len(f.children); i++ {
		size = append(size, getSmallDirs(&f.children[i], max)...)
	}

	return size
}

// go through dirs and return the pointer to each small dir
func getBigDirs(f *File, max int) []int {
	var size []int
	if f.kind == "file" {
		return size
	}
	if f.size > max {
		size = append(size, f.size)
	}
	for i := 0; i < len(f.children); i++ {
		size = append(size, getBigDirs(&f.children[i], max)...)
	}

	return size
}

func printDir(f File, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	fmt.Println(f.name, f.size)
	for i := 0; i < len(f.children); i++ {
		printDir(f.children[i], depth+1)
	}
}

func getSmallestToRemove(dir_sizes []int, space_req int) int {
	smallest := dir_sizes[0]
	for i := 1; i < len(dir_sizes); i++ {
		if dir_sizes[i] > space_req && dir_sizes[i] < smallest {
			smallest = dir_sizes[i]
		}
	}
	return smallest
}
