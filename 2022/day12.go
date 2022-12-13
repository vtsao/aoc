// This program implements the solution for
// https://adventofcode.com/2022/day/12.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("day12_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var heightmap [][]rune
	var start, end coord
	var startingPoints []coord
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		var row []rune
		for x, height := range scanner.Text() {
			// Purposely don't add the starting point from part 1, because we can just
			// reuse that answer when checking for the min steps from any elevation
			// 'a'.
			if height == 'a' {
				startingPoints = append(startingPoints, coord{x, y})
			}

			if height == 'S' {
				start = coord{x, y}
				height = 'a'
			} else if height == 'E' {
				end = coord{x, y}
				height = 'z'
			}
			row = append(row, height)
		}
		heightmap = append(heightmap, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minSteps1 := bfs(heightmap, start, end)
	fmt.Printf("Part 1: %d\n", minSteps1)

	// I assume the starting point from Part 1 should be included in the candidate
	// set for Part 2, but unlikely to be the same min, since you only have one
	// input set to solve for.
	minSteps2 := minSteps1
	for _, start := range startingPoints {
		if steps := bfs(heightmap, start, end); steps < minSteps2 {
			minSteps2 = steps
		}
	}
	fmt.Printf("Part 2: %d\n", minSteps2)
}

type coord struct {
	x, y int
}

func bfs(heightmap [][]rune, start, end coord) int {
	queue := []coord{start}

	steps := 0
	visited := map[coord]any{}
	for len(queue) > 0 {
		levelLen := len(queue)
		for i := 0; i < levelLen; i++ {
			curPos := queue[i]
			if _, ok := visited[curPos]; ok {
				continue
			}
			visited[curPos] = nil

			if curPos == end {
				return steps
			}

			for _, n := range []coord{
				{curPos.x, curPos.y + 1},
				{curPos.x, curPos.y - 1},
				{curPos.x + 1, curPos.y},
				{curPos.x - 1, curPos.y},
			} {
				// Check for out of bounds.
				if n.x < 0 || n.y < 0 || n.x >= len(heightmap[0]) || n.y >= len(heightmap) {
					continue
				}
				// Check if we're allowed to move to the new height (any -ve height or
				// descent is always allowed).
				if heightmap[n.y][n.x]-heightmap[curPos.y][curPos.x] > 1 {
					continue
				}

				queue = append(queue, n)
			}
		}

		// Remove all the elements from this level in the queue. We do this at the
		// end, since this isn't a real queue. Slightly more efficient, but doesn't
		// matter for this input set.
		queue = queue[levelLen:]
		steps++
	}

	return math.MaxInt
}
