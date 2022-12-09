// This program implements the solution for https://adventofcode.com/2022/day/8.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day08_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the input into a 2D array (grid) of tree heights.
	var trees [][]int
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		var row []int
		for _, treeHeightStr := range line {
			treeHeight, _ := strconv.Atoi(string(treeHeightStr))
			row = append(row, treeHeight)
		}
		trees = append(trees, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Check to see if each tree in the grid is visible.
	visibleTrees := 0
	for y, row := range trees {
		for x := range row {
			if isVisible(x, y, trees) {
				visibleTrees++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", visibleTrees)
	// fmt.Printf("Part 2: %d\n", toDelSize)
}

type delta struct {
	x, y int
}

// isVisible checks to see if the specified tree is visible from any direction.
func isVisible(x, y int, trees [][]int) bool {
	for _, d := range []delta{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		if isVisibleDir(x, y, d.x, d.y, trees) {
			return true
		}
	}
	return false
}

// isVisibleDir checks if the specified tree is visible in the direction
// specified by the x and y deltas.
func isVisibleDir(x, y, dx, dy int, trees [][]int) bool {
	treeHeight := trees[y][x]
	for newX, newY := x+dx, y+dy; !(newX < 0 || newX >= len(trees[0]) || newY < 0 || newY >= len(trees)); newX, newY = newX+dx, newY+dy {
		if trees[newY][newX] >= treeHeight {
			return false
		}
	}
	return true
}
