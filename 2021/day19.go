// This program implements the solution for
// https://adventofcode.com/2021/day/19.
//
// curl -b "$(cat .session)" -o day19_input.txt https://adventofcode.com/2021/day/19/input
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type scanner struct {
	origin coord
	coords map[coord]struct{}
}

type coord struct {
	x, y, z int
}

func roll(s scanner) scanner {
	rolled := scanner{coords: map[coord]struct{}{}}
	for c := range s.coords {
		rolled.coords[coord{x: c.x, y: c.z, z: -c.y}] = struct{}{}
	}
	return rolled
}

func turn(s scanner) scanner {
	turned := scanner{coords: map[coord]struct{}{}}
	for c := range s.coords {
		turned.coords[coord{x: -c.y, y: c.x, z: c.z}] = struct{}{}
	}
	return turned
}

func beacons(scanner1, scanner2 scanner) (scanner, bool) {
	// Scanner 1's coords are always absolute. Scanner 2's are always relative as
	// they have not been corrected yet.
	for coord1 := range scanner1.coords {
		for coord2 := range scanner2.coords {
			// Test to see if this pair of coordinates is overlapping. We know it's
			// overlapping if at least 11 other coordinates overlap with this guess.

			// If we assume this pair of coordinates overlaps, then we can calculate
			// scanner 2's absolute position (b/c scanner 1's position is absolute).
			newScanner2 := scanner{
				origin: coord{
					x: coord1.x - coord2.x,
					y: coord1.y - coord2.y,
					z: coord1.z - coord2.z,
				},
				coords: map[coord]struct{}{},
			}
			// Now using the assumed scanner 2's absolute position, translate all its
			// beacon coordinates from relative positions to absolute positions.
			for c := range scanner2.coords {
				newScanner2.coords[coord{
					x: newScanner2.origin.x + c.x,
					y: newScanner2.origin.y + c.y,
					z: newScanner2.origin.z + c.z,
				}] = struct{}{}
			}
			// Count how many beacons overlap. This is trivial now that all beacon
			// coordinates across scanner 1 and 2 are absolute.
			overlapping := 0
			for c := range scanner1.coords {
				if _, ok := newScanner2.coords[c]; ok {
					overlapping++
				}
			}
			if overlapping >= 12 {
				// If we have at least 12 overlapping beacons, this guess is correct and
				// we have corrected scanner 2.
				return newScanner2, true
			}
		}
	}

	return scanner{}, false
}

func compare(scanner1, scanner2 scanner) (scanner, bool) {
	// Calculate each of the 24 positions scanner 2 could be in. One of them might
	// match scanner 1's positioning system.
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			scanner2 = roll(scanner2)
			if newScanner2, ok := beacons(scanner1, scanner2); ok {
				return newScanner2, true
			}
			for k := 0; k < 3; k++ {
				scanner2 = turn(scanner2)
				if newScanner2, ok := beacons(scanner1, scanner2); ok {
					return newScanner2, true
				}
			}
		}
		scanner2 = roll(turn(roll(scanner2)))
	}

	return scanner{}, false
}

func parseInput() []scanner {
	file, _ := os.Open("day19_input.txt")
	defer file.Close()

	var scanners []scanner
	s := scanner{coords: map[coord]struct{}{}}
	inputReader := bufio.NewScanner(file)
	for inputReader.Scan() {
		line := inputReader.Text()
		switch {
		case line == "":
			scanners = append(scanners, s)
			s = scanner{coords: map[coord]struct{}{}}
		case strings.HasPrefix(line, "---"):
		default:
			coordParts := strings.Split(line, ",")
			x, _ := strconv.Atoi(coordParts[0])
			y, _ := strconv.Atoi(coordParts[1])
			z, _ := strconv.Atoi(coordParts[2])
			s.coords[coord{x: x, y: y, z: z}] = struct{}{}
		}
	}
	scanners = append(scanners, s)

	return scanners
}

func manhattanDist(c1, c2 coord) int {
	return int(math.Abs(float64(c2.x)-float64(c1.x)) + math.Abs(float64(c2.y)-float64(c1.y)) + math.Abs(float64(c2.z)-float64(c1.z)))
}

func main() {
	scanners := parseInput()

	done := map[int]scanner{0: scanners[0]}
	newlyFound := []scanner{scanners[0]}
	for len(newlyFound) > 0 {
		s := newlyFound[0]
		newlyFound = newlyFound[1:]
		for i, remainingScanner := range scanners {
			if _, ok := done[i]; ok {
				continue
			}
			if foundScanner, ok := compare(s, remainingScanner); ok {
				newlyFound = append(newlyFound, foundScanner)
				done[i] = foundScanner
			}
		}
	}
	beacons := map[coord]struct{}{}
	for _, s := range done {
		for c := range s.coords {
			beacons[c] = struct{}{}
		}
	}
	fmt.Printf("Part 1: %d\n", len(beacons))

	largestMDist := 0
	for i := 0; i < len(done); i++ {
		for j := i; j < len(done); j++ {
			if d := manhattanDist(done[i].origin, done[j].origin); d > largestMDist {
				largestMDist = d
			}
		}
	}
	fmt.Printf("Part 2: %d\n", largestMDist)
}
