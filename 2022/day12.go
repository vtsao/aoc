// This program implements the solution for
// https://adventofcode.com/2022/day/12.
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day12_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Part 1: %d\n", )
	// fmt.Printf("Part 2: %d\n", )
}
