// This program implements the solution for https://adventofcode.com/2022/day/3.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day03_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sumPerSack := 0
	sumPerGroup := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// This is hacky for the groups, could probably have done this in a loop,
		// but whatever.

		rucksack := scanner.Text()
		c := commonItemPerSack(rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:])
		sumPerSack += priority(c)
		g1 := rucksack

		scanner.Scan()
		rucksack = scanner.Text()
		c = commonItemPerSack(rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:])
		sumPerSack += priority(c)
		g2 := rucksack

		scanner.Scan()
		rucksack = scanner.Text()
		c = commonItemPerSack(rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:])
		sumPerSack += priority(c)
		g3 := rucksack

		c = commonItemPerGroup(g1, g2, g3)
		fmt.Printf("%q\n", c)
		sumPerGroup += priority(c)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", sumPerSack)
	fmt.Printf("Part 2: %d\n", sumPerGroup)
}

func commonItemPerSack(comp1, comp2 string) rune {
	comp1Set := map[rune]any{}
	for _, r := range comp1 {
		comp1Set[r] = nil
	}

	for _, r := range comp2 {
		if _, ok := comp1Set[r]; ok {
			return r
		}
	}

	// Will never get here with valid input.
	return 0
}

func commonItemPerGroup(g1, g2, g3 string) rune {
	g1Set := map[rune]any{}
	for _, r := range g1 {
		g1Set[r] = nil
	}
	g2Set := map[rune]any{}
	for _, r := range g2 {
		g2Set[r] = nil
	}

	for _, r := range g3 {
		if _, ok := g1Set[r]; !ok {
			continue
		}
		if _, ok := g2Set[r]; !ok {
			continue
		}
		return r
	}

	// Will never get here with valid input.
	return 0
}

func priority(item rune) int {
	if item >= 'a' {
		return int(item-'a') + 1
	}
	return int(item-'A') + 26 + 1
}
