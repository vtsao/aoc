// This program implements the solution for https://adventofcode.com/2022/day/1.
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
	file, err := os.Open("day01_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var elfCals [][]int
	var calories []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elfCals = append(elfCals, calories)
			calories = []int{}
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		calories = append(calories, cal)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	elfCals = append(elfCals, calories)

	eflTotals := totalCals(elfCals)

	fmt.Printf("Part 1: %d\n", eflTotals[len(eflTotals)-1])

	top3 := eflTotals[len(eflTotals)-1] + eflTotals[len(eflTotals)-2] + eflTotals[len(eflTotals)-3]
	fmt.Printf("Part 2: %d\n", top3)
}

func totalCals(elfCals [][]int) []int {
	var eflTotals []int

	for _, elf := range elfCals {
		totalCals := 0
		for _, cals := range elf {
			totalCals += cals
		}
		eflTotals = append(eflTotals, totalCals)
	}

	sort.Ints(eflTotals)
	return eflTotals
}
