package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var adapters []int
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		adapter, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		adapters = append(adapters, adapter)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(adapters)

	// We start off with a diff of 3 having a count of 1, which represents our
	// device's built-in adapter, which always has a rating of 3 higher than the
	// highest adapter.
	diffs := map[int]int{3: 1}
	curJolt := 0
	for _, adapter := range adapters {
		diff := adapter - curJolt
		diffs[diff]++
		curJolt = adapter
	}

	fmt.Printf("Part 1: 1-jolt differences * 3-jolt differences: %d\n", diffs[1]*diffs[3])
}
