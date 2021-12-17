// This program implements the solution for
// https://adventofcode.com/2021/day/17.
//
// curl -b "$(cat .session)" -o day17_input.txt https://adventofcode.com/2021/day/17/input
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bounds struct {
	lower, upper int
}

func hit(xVel, yVel int, xBounds, yBounds bounds) (int, bool) {
	maxY := 0
	x, y := 0, 0
	for {
		x += xVel
		y += yVel

		if y > maxY {
			maxY = y
		}

		if xVel > 0 {
			xVel--
		} else if xVel < 0 {
			xVel++
		}
		yVel--

		if x >= xBounds.lower && x <= xBounds.upper && y >= yBounds.lower && y <= yBounds.upper {
			return maxY, true
		}
		if x > xBounds.upper || y < yBounds.lower {
			break
		}
	}

	return 0, false
}

func parseInput() (bounds, bounds) {
	file, _ := os.Open("day17_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := strings.TrimLeft(scanner.Text(), "target area: ")
	sParts := strings.Split(s, ", ")

	xBoundsStr := strings.TrimLeft(sParts[0], "x=")
	xBoundsParts := strings.Split(xBoundsStr, "..")
	xBoundsLower, _ := strconv.Atoi(xBoundsParts[0])
	xBoundsUpper, _ := strconv.Atoi(xBoundsParts[1])
	xBounds := bounds{lower: xBoundsLower, upper: xBoundsUpper}

	yBoundsStr := strings.TrimLeft(sParts[1], "y=")
	yBoundsParts := strings.Split(yBoundsStr, "..")
	yBoundsLower, _ := strconv.Atoi(yBoundsParts[0])
	yBoundsUpper, _ := strconv.Atoi(yBoundsParts[1])
	yBounds := bounds{lower: yBoundsLower, upper: yBoundsUpper}

	return xBounds, yBounds
}

func main() {
	xBounds, yBounds := parseInput()

	maxY := 0
	vals := 0
	for x := 0; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			if y, ok := hit(x, y, xBounds, yBounds); ok {
				vals++
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", maxY)
	fmt.Printf("Part 2: %d\n", vals)
}
