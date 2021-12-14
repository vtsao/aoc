// This program implements the solution for
// https://adventofcode.com/2021/day/15.
//
// curl -b "$(cat .session)" -o day15_input.txt https://adventofcode.com/2021/day/15/input
package main

import (
	"bufio"
	"log"
	"os"
)

func parseInput() ([]string, error) {
	file, err := os.Open("day15_input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	lines, err := parseInput()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		log.Println(line)
	}

	// fmt.Printf("Part 1: %d\n", )
	// fmt.Printf("Part 2: %d\n", )
}
