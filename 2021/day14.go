// This program implements the solution for
// https://adventofcode.com/2021/day/14.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type key struct {
	pair           string
	remainingSteps int
}

// insert recursively expands a pair and caches results. We don't need the
// actual string generated by the expansion, only a count of the elements. Note
// this only returns the count of the elements of the expanded string between
// the pair. It does not count the pair itself.
func insert(pair string, insertionRules map[string]string, remainingSteps int, cache map[key]map[string]int) map[string]int {
	if remainingSteps == 0 {
		return map[string]int{}
	}
	if cached, ok := cache[key{pair: pair, remainingSteps: remainingSteps}]; ok {
		return cached
	}

	pair0 := string(pair[0])
	pair1 := string(pair[1])
	insertion := insertionRules[pair0+pair1]
	elemCount := map[string]int{insertion: 1}

	elemCount0 := insert(pair0+insertion, insertionRules, remainingSteps-1, cache)
	for elem, count := range elemCount0 {
		elemCount[elem] += count
	}

	elemCount1 := insert(insertion+pair1, insertionRules, remainingSteps-1, cache)
	for elem, count := range elemCount1 {
		elemCount[elem] += count
	}

	cache[key{pair: pair, remainingSteps: remainingSteps}] = elemCount
	return elemCount
}

func elemCountDiff(template string, insertionRules map[string]string, steps int) int {
	// Go through each of the pairs in the sequence and expand them to get their
	// element count.
	cache := map[key]map[string]int{}
	elemCount := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i]) + string(template[i+1])
		elemCount[string(template[i])]++
		elemCountN := insert(pair, insertionRules, steps, cache)
		for elem, count := range elemCountN {
			elemCount[elem] += count
		}
	}
	elemCount[string(template[len(template)-1])]++

	var counts []int
	for _, count := range elemCount {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	return counts[len(counts)-1] - counts[0]
}

func main() {
	file, err := os.Open("day14_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := scanner.Text()
	scanner.Scan()

	insertionRules := map[string]string{}
	for scanner.Scan() {
		insertionRulesParts := strings.Split(scanner.Text(), " -> ")
		insertionRules[insertionRulesParts[0]] = insertionRulesParts[1]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", elemCountDiff(template, insertionRules, 10))
	fmt.Printf("Part 2: %d\n", elemCountDiff(template, insertionRules, 40))
}