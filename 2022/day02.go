// This program implements the solution for https://adventofcode.com/2022/day/2.
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day02_input.txt")
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

	// fmt.Printf("Part 1: %d\n", eflTotals[len(eflTotals)-1])
	// fmt.Printf("Part 2: %d\n", top3)
}
