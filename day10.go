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

	// First load all adapters into memory because we need to sort them.
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

	// Solve Part 1.
	//
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

	// Solve Part 2. Could have done both parts in the same loop, but doing it
	// separately makes this more readable. Performance also doesn't affect our
	// submission on the Advent of Code site.
	//
	// We use DP to solve part 2, because backtracking would be O(2^n). We want to
	// know how many arrangements there are for jolt i. The answer to this is the
	// sum of how many arragements there are for jolts i-1, i-2, and i-3. Not all
	// of these exist due to the input, but at least one of these exist (otherwise
	// we could not solve this puzzle). We do bottom-up DP with the base case that
	// there is 1 way to make 0 jolts, which is just the charging outlet.
	//
	// E.g., to figure out how many arrangments there are for 6 jolts (because you
	// have a 6 jolt adapter), sum the arrangements for 5, 4, and 3 jolts. You
	// would have calculated this from a previous step (if any of them don't
	// exist, their values would be 0, which is fine).
	dp := map[int]int{0: 1}
	for _, adapter := range adapters {
		a := 0
		for i := adapter - 3; i < adapter; i++ {
			a += dp[i]
		}
		dp[adapter] = a
	}
	fmt.Printf("Part 2: arrangements: %d\n", dp[adapters[len(adapters)-1]])
}
