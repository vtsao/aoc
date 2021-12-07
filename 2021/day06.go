// This program implements the solution for https://adventofcode.com/2021/day/6.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type key struct {
	timer, days int
}

var cache = map[key]int{}

func computeOffspring(timer, days int) int {
	if cached, ok := cache[key{timer: timer, days: days}]; ok {
		return cached
	}
	if timer > days || days <= 0 {
		return 0
	}

	remDays := days - timer
	offspring := 1 + computeOffspring(9, remDays)
	for remDays >= 7 {
		remDays -= 7
		offspring += 1 + computeOffspring(9, remDays)
	}

	cache[key{timer: timer, days: days}] = offspring
	return offspring
}

func main() {
	file, err := os.Open("day06_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var fishes []int
	for _, fish := range strings.Split(scanner.Text(), ",") {
		timer, err := strconv.Atoi(fish)
		if err != nil {
			log.Fatal(err)
		}
		fishes = append(fishes, timer+1)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalFishes1 := 0
	for _, fish := range fishes {
		totalFishes1 += 1 + computeOffspring(fish, 80)
	}
	fmt.Printf("Part 1: %d\n", totalFishes1)

	totalFishes2 := 0
	for _, fish := range fishes {
		totalFishes2 += 1 + computeOffspring(fish, 256)
	}
	fmt.Printf("Part 2: %d\n", totalFishes2)
}
