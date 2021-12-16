// This program implements the solution for
// https://adventofcode.com/2021/day/24.
//
// curl -b "$(cat .session)" -o day24_input.txt https://adventofcode.com/2021/day/24/input
package main

import (
	"bufio"
	"log"
	"os"
)

func parseInput() []string {
	file, _ := os.Open("day24_input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func main() {
	lines := parseInput()

	for _, line := range lines {
		log.Println(line)
	}

	// fmt.Printf("Part 1: %d\n", )
	// fmt.Printf("Part 2: %d\n", )
}
