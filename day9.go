package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func valid(nums []int, num int) bool {
	// Probably don't actually need to optimize into a set like this, since there
	// are only ever 25 numbers to look through. This makes it O(n) instead of
	// O(n^2).
	numsSet := map[int]interface{}{}
	for _, n := range nums {
		numsSet[n] = nil
	}

	for n := range numsSet {
		target := num - n
		if _, ok := numsSet[target]; ok {
			return true
		}
	}

	return false
}

func findWeakness(nums []int, target int) int {
	sumsToIdx := map[int]int{}
	curSum := 0

	// Keep a rolling sum and a map of the rolling sums to their indexes so we can
	// do this in O(n) time.
	//
	// We want (j > i and we're on j):
	//   rolling sum at j - rolling sum at i == target
	//   rolling sum at i == target - rolling sum at j
	//
	// Since we've kept track of all prev rolling sums in a map, we see if we can
	// find rolling sum at i for the cur rolling sum at j.
	var start, end int
	for i, num := range nums {
		curSum += num
		if s, ok := sumsToIdx[curSum-target]; ok {
			start = s + 1
			end = i
			break
		}
		sumsToIdx[curSum] = i
	}

	// Note this assumes there is an answer, otherwise it will return an incorrect
	// results.
	min, max := math.MaxInt64, 0
	for i := start; i < end+1; i++ {
		num := nums[i]
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min + max
}

func main() {
	file, err := os.Open("day9_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nums []int
	var invalidNum int
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)

		// There is a preamble of 25 numbers, so don't validate the first 25 nums.
		if i < 25 {
			continue
		}

		// We need all the nums for part 2, that's why we keep going and store all
		// the nums instead of only keeping 25.
		if !valid(nums[len(nums)-26:], num) {
			// Note this assumes there is only one invalid number, otherwise this will
			// overwrite it.
			invalidNum = num
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: first number that is not the sum of two of the 25 numbers before it: %d\n", invalidNum)

	w := findWeakness(nums, invalidNum)
	fmt.Printf("Part 2: encryption weakness: %d\n", w)
}
