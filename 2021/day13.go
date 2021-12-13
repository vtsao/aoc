// This program implements the solution for
// https://adventofcode.com/2021/day/13.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	width  = 1311
	height = 957
)

type fold struct {
	dir string
	val int
}

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, col := range row {
			if col == true {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func countDots(grid [][]bool) int {
	cnt := 0
	for _, row := range grid {
		for _, col := range row {
			if col == true {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	file, err := os.Open("day13_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	var folds []*fold
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "fold along"):
			foldParts := strings.Split(strings.Split(line, " ")[2], "=")
			foldVal, err := strconv.Atoi(foldParts[1])
			if err != nil {
				log.Fatal(err)
			}
			f := &fold{
				dir: foldParts[0],
				val: foldVal,
			}
			folds = append(folds, f)
		default:
			coord := strings.Split(line, ",")
			x, err := strconv.Atoi(coord[0])
			if err != nil {
				log.Fatal(err)
			}
			y, err := strconv.Atoi(coord[1])
			if err != nil {
				log.Fatal(err)
			}
			grid[y][x] = true
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, fold := range folds {
		var newGrid [][]bool
		switch fold.dir {
		case "y":
			if fold.val < len(grid)-fold.val-1 {
				copyStartRow := fold.val + fold.val + 1
				mergeStartRow := fold.val + 1
				endRow := len(grid)
				for i := mergeStartRow; i < copyStartRow; i++ {
					var newRow []bool
					for j := 0; j < len(grid[0]); j++ {
						if grid[i][j] == true || grid[fold.val-(i-fold.val)][j] == true {
							newRow = append(newRow, true)
							continue
						}
						newRow = append(newRow, false)
					}
					newGrid = append(newGrid, newRow)
				}
				for i := copyStartRow; i < endRow; i++ {
					newGrid = append(newGrid, grid[i])
				}
				break
			}
			copyStartRow := 0
			mergeStartRow := fold.val - (len(grid) - fold.val - 1)
			endRow := fold.val
			for i := copyStartRow; i < mergeStartRow; i++ {
				newGrid = append(newGrid, grid[i])
			}
			for i := mergeStartRow; i < endRow; i++ {
				var newRow []bool
				for j := 0; j < len(grid[0]); j++ {
					if grid[i][j] == true || grid[fold.val+(fold.val-i)][j] == true {
						newRow = append(newRow, true)
						continue
					}
					newRow = append(newRow, false)
				}
				newGrid = append(newGrid, newRow)
			}
		case "x":
			if fold.val < len(grid[0])-fold.val-1 {
				copyStartCol := fold.val + fold.val + 1
				mergeStartCol := fold.val + 1
				endCol := len(grid[0])
				for i := 0; i < len(grid); i++ {
					var newRow []bool
					for j := mergeStartCol; j < copyStartCol; j++ {
						if grid[i][j] == true || grid[i][fold.val-(j-fold.val)] == true {
							newRow = append(newRow, true)
							continue
						}
						newRow = append(newRow, false)
					}
					newGrid = append(newGrid, newRow)
				}
				for i := 0; i < len(grid); i++ {
					for j := copyStartCol; j < endCol; j++ {
						newGrid[i] = append(newGrid[i], grid[i][j])
					}
				}
				break
			}
			copyStartCol := 0
			mergeStartCol := fold.val - (len(grid[0]) - fold.val - 1)
			endCol := fold.val
			for i := 0; i < len(grid); i++ {
				var newRow []bool
				for j := copyStartCol; j < mergeStartCol; j++ {
					newRow = append(newRow, grid[i][j])
				}
				newGrid = append(newGrid, newRow)
			}
			for i := 0; i < len(grid); i++ {
				for j := mergeStartCol; j < endCol; j++ {
					if grid[i][j] == true || grid[i][fold.val+(fold.val-j)] == true {
						newGrid[i] = append(newGrid[i], true)
						continue
					}
					newGrid[i] = append(newGrid[i], false)
				}
			}
		}

		grid = newGrid
		if i == 0 {
			fmt.Printf("Part 1: %d\n", countDots(grid))
		}
	}
	fmt.Println("Part 2:")
	printGrid(grid)
}
