// This program implements the solution for https://adventofcode.com/2021/day/5.
package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type line struct {
	x1, y1, x2, y2 int
}

func main() {
	file, err := os.Open("day05_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []*line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " -> ")
		start := strings.Split(input[0], ",")
		end := strings.Split(input[1], ",")
		// x1, _ := strconv.Atoi(start[0])
		// lines = append(lines, &line{
		// 	x1: start[0],
		// 	y1: start[1],
		// 	x2: end[0],
		// 	y2: end[1],
		// })
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Part 1: %d\n", )
}
