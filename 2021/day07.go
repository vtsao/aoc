// This program implements the solution for https://adventofcode.com/2021/day/7.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day07_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var positions []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, posStr := range strings.Split(scanner.Text(), ",") {
		pos, err := strconv.Atoi(posStr)
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, pos)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(positions)

	// Cache fuel costs for each move value for part 2. E.g., we can look up the
	// fuel cost for x, where x is an arbitrary number of moves.
	fuelDiffs2 := []int{0}
	for i := 1; i <= positions[len(positions)-1]; i++ {
		fuelDiffs2 = append(fuelDiffs2, fuelDiffs2[i-1]+i)
	}

	minFuel1 := math.MaxInt64
	minFuel2 := math.MaxInt64
	for i := 0; i <= positions[len(positions)-1]; i++ {
		fuel1 := 0
		for _, pos := range positions {
			fuel1 += int(math.Abs(float64(pos) - float64(i)))
		}
		if fuel1 < minFuel1 {
			minFuel1 = fuel1
		}

		fuel2 := 0
		for _, pos := range positions {
			fuel2 += fuelDiffs2[int(math.Abs(float64(pos)-float64(i)))]
		}
		if fuel2 < minFuel2 {
			minFuel2 = fuel2
		}
	}
	fmt.Printf("Part 1: %d\n", minFuel1)
	fmt.Printf("Part 2: %d\n", minFuel2)
}
