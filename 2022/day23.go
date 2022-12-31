// This program implements the solution for
// https://adventofcode.com/2022/day/23.
package main

import (
	_ "embed"
	"fmt"
	"math"
	"reflect"
	"strings"
)

//go:embed day23_input.txt
var input string

type coord struct {
	x, y int
}

// parse parses the input of elf positions and returns the position of each elf
// in a hash set of coordinates.
func parse() map[coord]any {
	elves := map[coord]any{}
	for y, row := range strings.Split(input, "\n") {
		for x, col := range row {
			if col == '#' {
				elves[coord{x, y}] = nil
			}
		}
	}
	return elves
}

// proposedCoord tracks the proposed new coordinate an elf wants to move to for
// each round. We need to keep the old coordinate to "revert" the next coord for
// an elf if it conflicts with another elf.
type proposedCoord struct {
	old, new coord
}

// round simulates the elves moving one round following the rules in the problem
// description
func round(roundNum int, elves map[coord]any) map[coord]any {
	// Track the proposed new coordinate each elf wants to move into. Don't worry
	// about conflicts for now.
	var newElves []proposedCoord

elves:
	for elf := range elves {
		newElf := proposedCoord{old: elf}

		// No other elves around this elf, this elf does not move.
		if !hasNeighbors(elf, elves, -1) {
			newElf.new = elf
			newElves = append(newElves, newElf)
			continue
		}

		// The initial direction to test for a valid move for this round increments
		// per round, so just mod it by what round we're on.
		for i, dir := 0, roundNum; ; i, dir = i+1, dir+1 {
			// This elf cannot move in any direction, so it does not move.
			if i == 4 {
				newElf.new = elf
				newElves = append(newElves, newElf)
				continue elves
			}

			// Mod the direction to wrap around (i.e., dir 3's next dir is dir 0).
			dir %= 4
			if hasNeighbors(elf, elves, dir) {
				continue
			}
			switch dir {
			case 0:
				newElf.new = coord{elf.x, elf.y - 1}
			case 1:
				newElf.new = coord{elf.x, elf.y + 1}
			case 2:
				newElf.new = coord{elf.x - 1, elf.y}
			case 3:
				newElf.new = coord{elf.x + 1, elf.y}
			}
			newElves = append(newElves, newElf)
			break
		}
	}

	// Reconcile any conflicts amongst the elves.
	return reconcile(newElves)
}

// reconcile checks to see if more than one elf wanted to move into the same
// coordinate for a given round's proposed coordinates. If they did, those elves
// don't move for that round, so ensure they keep their existing position.
func reconcile(proposedCoords []proposedCoord) map[coord]any {
	// Count how many elves proposed each new coordinate.
	newCoords := map[coord]int{}
	for _, proposedElfCoord := range proposedCoords {
		newCoords[proposedElfCoord.new]++
	}

	elves := map[coord]any{}
	for _, proposedElfCoord := range proposedCoords {
		// If there are conflicts, "revert" that elf's position so they do not move
		// for this round.
		if newCoords[proposedElfCoord.new] > 1 {
			elves[proposedElfCoord.old] = nil
			continue
		}
		elves[proposedElfCoord.new] = nil
	}

	return elves
}

// hasNeighbors checks to see if the specified elf has any neighboring elves in
// the specified direction. A direction of -1 means to check all around the elf.
func hasNeighbors(elf coord, elves map[coord]any, dir int) bool {
	neighbors := []coord{
		{elf.x, elf.y - 1},
		{elf.x + 1, elf.y - 1},
		{elf.x + 1, elf.y},
		{elf.x + 1, elf.y + 1},
		{elf.x, elf.y + 1},
		{elf.x - 1, elf.y + 1},
		{elf.x - 1, elf.y},
		{elf.x - 1, elf.y - 1},
	}
	switch dir {
	case 0: // north
		neighbors = []coord{{elf.x - 1, elf.y - 1}, {elf.x, elf.y - 1}, {elf.x + 1, elf.y - 1}}
	case 1: // south
		neighbors = []coord{{elf.x - 1, elf.y + 1}, {elf.x, elf.y + 1}, {elf.x + 1, elf.y + 1}}
	case 2: // west
		neighbors = []coord{{elf.x - 1, elf.y + 1}, {elf.x - 1, elf.y}, {elf.x - 1, elf.y - 1}}
	case 3: // east
		neighbors = []coord{{elf.x + 1, elf.y - 1}, {elf.x + 1, elf.y}, {elf.x + 1, elf.y + 1}}
	}

	for _, n := range neighbors {
		if _, ok := elves[n]; ok {
			return true
		}
	}
	return false
}

// groundTiles counts the number of ground tiles in the smallest rectangle that
// encompasses the elves.
func groundTiles(elves map[coord]any) int {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for elf := range elves {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}
	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func main() {
	elves := parse()

	newElves := elves
	for i := 0; i < 10; i++ {
		newElves = round(i, newElves)
	}
	fmt.Printf("Part 1: %d\n", groundTiles(newElves))

	newElves = elves
	// Keep simulating rounds until there is no change for a round.
	for i := 0; ; i++ {
		next := round(i, newElves)
		if reflect.DeepEqual(next, newElves) {
			// With valid input, we'll always get an answer. Also rounds start at 1 in
			// the problem description, so add 1 to our answer.
			fmt.Printf("Part 2: %d\n", i+1)
			break
		}
		newElves = next
	}
}
