// This program implements the solution for
// https://adventofcode.com/2022/day/14.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day14_input.txt")

	rocks := map[coord]any{}
	largestY := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := strings.Split(scanner.Text(), " -> ")
		for i := 0; i < len(path)-1; i++ {
			from := parse(path[i])
			to := parse(path[i+1])

			if from.x == to.x {
				smallerY, largerY := from.y, to.y
				if from.y > to.y {
					smallerY, largerY = largerY, smallerY
				}
				for y := smallerY; y <= largerY; y++ {
					rocks[coord{from.x, y}] = nil
				}

				if largerY > largestY {
					largestY = largerY
				}
				continue
			}

			smallerX, largerX := from.x, to.x
			if from.x > to.x {
				smallerX, largerX = largerX, smallerX
			}
			for x := smallerX; x <= largerX; x++ {
				rocks[coord{x, from.y}] = nil
			}
		}
	}

	numResting := sand(rocks, nil)
	fmt.Printf("Part 1: %d\n", numResting)

	floor := largestY + 2
	numRestingWithFloor := sand(rocks, &floor)
	fmt.Printf("Part 2: %d\n", numRestingWithFloor)
}

// parse parses a coordinate string like 499,4 into a coord struct.
func parse(s string) coord {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coord{x, y}
}

type coord struct {
	x, y int
}

func sand(rocks map[coord]any, floor *int) int {
	restingSand := map[coord]any{}
cycle:
	for {
		cur := coord{500, 0}
		for {
			if floor != nil {
				if _, ok := restingSand[coord{500, 0}]; ok {
					break cycle
				}
			} else if cur.y > 100000 {
				break cycle
			}

			if next := (coord{cur.x, cur.y + 1}); !blocked(next, rocks, restingSand, floor) {
				cur = next
				continue
			}
			if next := (coord{cur.x - 1, cur.y + 1}); !blocked(next, rocks, restingSand, floor) {
				cur = next
				continue
			}
			if next := (coord{cur.x + 1, cur.y + 1}); !blocked(next, rocks, restingSand, floor) {
				cur = next
				continue
			}
			restingSand[cur] = nil
			break
		}
	}

	return len(restingSand)
}

func blocked(pos coord, rocks, restingSand map[coord]any, floor *int) bool {
	if floor != nil {
		if pos.y == *floor {
			return true
		}
	}

	if _, ok := rocks[pos]; ok {
		return true
	}
	if _, ok := restingSand[pos]; ok {
		return true
	}
	return false
}
