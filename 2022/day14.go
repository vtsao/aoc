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

			// "Draw a line" from one coord to the next and store each coord in a map.
			// This represents the locations of each rock.

			// With how the problem is structured, we only deal with horizontal or
			// vertical lines, so we can compare equal x's or y's to determine which
			// one it is.
			if from.x == to.x {
				// The "line" can start from any coord to another, so the "next" coord
				// isn't always larger. Since the direction we "draw" the line from
				// doesn't matter, just go from the smaller one to the larger one.
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

// parse parses a coordinate string like "499,4" into a coord struct.
func parse(s string) coord {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coord{x, y}
}

type coord struct {
	x, y int
}

// sand simulates the falling sand in the cave.
func sand(rocks map[coord]any, floor *int) int {
	restingSand := map[coord]any{}
cycle:
	for {
		cur := coord{500, 0}
		for {
			// We keep simulating each grain of sand that starts from 500,0 until no
			// sand can fall anymore (if there's a floor, part 1), or no sand comes to
			// rest anymore (no floor, part 1).
			if floor != nil {
				if _, ok := restingSand[coord{500, 0}]; ok {
					break cycle
				}
			} else if cur.y > 100000 {
				// This is hacky, we assume if this grain of sand hasn't come to rest
				// for an arbitrary amount of cycles, no more sand can come to rest.
				// I'm sure there's an actual limit we can calculate and target here?
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

// blocked determines whether a coord can have sand in it given the current
// state of the cave. The logic is slightly different depending on whether we're
// using a floor.
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
