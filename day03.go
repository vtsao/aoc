package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func trees(grid []string, right, down int) int {
	trees := 0
	x, y := right, down
	for y < len(grid) {
		if string(grid[y][x]) == "#" {
			trees++
		}
		x += right
		x %= len(grid[0])
		y += down
	}

	return trees
}

func main() {
	file, err := os.Open("day03_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d trees\n", trees(grid, 3, 1))

	trees2 := trees(grid, 1, 1)
	trees2 *= trees(grid, 3, 1)
	trees2 *= trees(grid, 5, 1)
	trees2 *= trees(grid, 7, 1)
	trees2 *= trees(grid, 1, 2)
	fmt.Printf("Part 2: %d trees\n", trees2)
}
