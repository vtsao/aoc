// This program implements the solution for https://adventofcode.com/2022/day/9.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day09_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := rope{
		visitedSecondKnot: map[coord]any{},
		visitedTail:       map[coord]any{},
	}
	for i := 0; i < 10; i++ {
		r.knots = append(r.knots, coord{})
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), " ")
		times, err := strconv.Atoi(lineParts[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < times; i++ {
			r.move(lineParts[0])
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", len(r.visitedSecondKnot))
	fmt.Printf("Part 2: %d\n", len(r.visitedTail))
}

type coord struct {
	x, y int
}

type rope struct {
	knots             []coord
	visitedSecondKnot map[coord]any
	visitedTail       map[coord]any
}

// move moves the head knot and then all its tails.
func (r *rope) move(dir string) {
	switch dir {
	case "U":
		r.knots[0].y--
	case "D":
		r.knots[0].y++
	case "L":
		r.knots[0].x--
	case "R":
		r.knots[0].x++
	}

	for i := 1; i < len(r.knots); i++ {
		r.updateTail(i)
	}
}

// updateTail updates the position of the tail following the rules given in the
// problem description based on how the knot in front of it moved.
//
// Handling each case separately is probably not ideal, there's probably a
// shorter way to do this in some loop.
func (r *rope) updateTail(tailIdx int) {
	headIdx := tailIdx - 1

	switch {
	// Same row.
	case r.knots[headIdx].x == r.knots[tailIdx].x:
		switch {
		case r.knots[headIdx].y-r.knots[tailIdx].y >= 2:
			r.knots[tailIdx].y++
		case r.knots[tailIdx].y-r.knots[headIdx].y == 2:
			r.knots[tailIdx].y--
		}

	// Same col.
	case r.knots[headIdx].y == r.knots[tailIdx].y:
		switch {
		case r.knots[headIdx].x-r.knots[tailIdx].x >= 2:
			r.knots[tailIdx].x++
		case r.knots[tailIdx].x-r.knots[headIdx].x >= 2:
			r.knots[tailIdx].x--
		}

	// Top-right quad.
	case r.knots[headIdx].y < r.knots[tailIdx].y && r.knots[headIdx].x > r.knots[tailIdx].x:
		if (r.knots[tailIdx].y-r.knots[headIdx].y)+(r.knots[headIdx].x-r.knots[tailIdx].x) >= 3 {
			r.knots[tailIdx].y--
			r.knots[tailIdx].x++
		}

	// Top-left quad.
	case r.knots[headIdx].y < r.knots[tailIdx].y && r.knots[headIdx].x < r.knots[tailIdx].x:
		if (r.knots[tailIdx].y-r.knots[headIdx].y)+(r.knots[tailIdx].x-r.knots[headIdx].x) >= 3 {
			r.knots[tailIdx].y--
			r.knots[tailIdx].x--
		}

	// Bottom-left quad.
	case r.knots[headIdx].y > r.knots[tailIdx].y && r.knots[headIdx].x < r.knots[tailIdx].x:
		if (r.knots[headIdx].y-r.knots[tailIdx].y)+(r.knots[tailIdx].x-r.knots[headIdx].x) >= 3 {
			r.knots[tailIdx].y++
			r.knots[tailIdx].x--
		}

	// Bottom-right quad.
	case r.knots[headIdx].y > r.knots[tailIdx].y && r.knots[headIdx].x > r.knots[tailIdx].x:
		if (r.knots[headIdx].y-r.knots[tailIdx].y)+(r.knots[headIdx].x-r.knots[tailIdx].x) >= 3 {
			r.knots[tailIdx].y++
			r.knots[tailIdx].x++
		}
	}

	// Keep track of the coordinates visited for the second knot for part 1 and
	// the tail for part 2.
	if tailIdx == 1 {
		r.visitedSecondKnot[r.knots[tailIdx]] = nil
	}
	if tailIdx == len(r.knots)-1 {
		r.visitedTail[r.knots[tailIdx]] = nil
	}
}
