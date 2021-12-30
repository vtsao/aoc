// This program implements the solution for
// https://adventofcode.com/2021/day/25.
//
// curl -b "$(cat .session)" -o day25_input.txt https://adventofcode.com/2021/day/25/input
package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInput() [][]rune {
	file, _ := os.Open("day25_input.txt")
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []rune
		for _, col := range scanner.Text() {
			row = append(row, rune(col))
		}
		grid = append(grid, row)
	}

	return grid
}

func print(grid [][]rune) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%s", string(col))
		}
		fmt.Println()
	}
	fmt.Println()
}

func run(grid [][]rune) int {
	steps := 0

	for {
		steps++
		moved := false
		for _, direction := range []string{"east", "south"} {
			nextGrid := make([][]rune, len(grid))
			for i, row := range grid {
				nextGrid[i] = make([]rune, len(row))
			}
			for i := len(grid) - 1; i >= 0; i-- {
				for j := len(grid[i]) - 1; j >= 0; j-- {
					col := grid[i][j]
					if nextGrid[i][j] == 0 {
						nextGrid[i][j] = col
					}
					switch {
					case direction == "east" && col == '>':
						nextJ := j + 1
						if nextJ == len(grid[i]) {
							nextJ = 0
						}
						if grid[i][nextJ] == '.' {
							nextGrid[i][nextJ] = '>'
							nextGrid[i][j] = '.'
							moved = true
						}
					case direction == "south" && col == 'v':
						nextI := i + 1
						if nextI == len(grid) {
							nextI = 0
						}
						if grid[nextI][j] == '.' {
							nextGrid[nextI][j] = 'v'
							nextGrid[i][j] = '.'
							moved = true
						}
					}
				}
			}
			grid = nextGrid
		}

		if !moved {
			break
		}
	}

	return steps
}

func main() {
	grid := parseInput()
	fmt.Printf("Part 1: %d\n", run(grid))
}
