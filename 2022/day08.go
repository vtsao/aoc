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

	// For each tree in the grid, check to see if it's visible and calculate its
	// scenic score. Keep track of the best scenic score.
	visibleTrees := 0
	maxScenicScore := 0
	for y, row := range trees {
		for x := range row {
			scenicScore, ok := isVisible(x, y, trees)
			if ok {
				visibleTrees++
			}
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Printf("Part 1: %d\n", visibleTrees)
	fmt.Printf("Part 2: %d\n", maxScenicScore)
}

type delta struct {
	x, y int
}

// isVisible checks to see if the specified tree is visible from any direction
// and calculates its scenic score.
func isVisible(x, y int, trees [][]int) (int, bool) {
	visible := false
	scenicScore := 1
	for _, d := range []delta{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		viewingDist, ok := isVisibleDir(x, y, d.x, d.y, trees)
		if ok {
			visible = true
		}
		scenicScore *= viewingDist
	}
	return scenicScore, visible
}

// isVisibleDir checks if the specified tree is visible in the direction
// specified by the x and y deltas and calculates its viewing distance.
func isVisibleDir(x, y, dx, dy int, trees [][]int) (int, bool) {
	treeHeight := trees[y][x]
	viewingDist := 0
	for newX, newY := x+dx, y+dy; !(newX < 0 || newX >= len(trees[0]) || newY < 0 || newY >= len(trees)); newX, newY = newX+dx, newY+dy {
		viewingDist++
		if trees[newY][newX] >= treeHeight {
			return viewingDist, false
		}
	}
	return viewingDist, true
}
