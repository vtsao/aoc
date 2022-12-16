// This program implements the solution for
// https://adventofcode.com/2022/day/16.
package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("day16_input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}

	// fmt.Printf("Part 1: %d\n", )
	// fmt.Printf("Part 2: %d\n", )
}
