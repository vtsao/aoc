// This program implements the solution for https://adventofcode.com/2021/day/5.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	x1, y1, x2, y2 int
}

type point struct {
	x, y int
}

func mark(points map[point]int, x, y int) bool {
	p := point{x: x, y: y}
	if overlaps, ok := points[p]; ok && overlaps == 1 {
		points[p]++
		return true
	}
	points[p]++
	return false
}

func main() {
	file, err := os.Open("day05_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []*line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " -> ")
		start := strings.Split(input[0], ",")
		end := strings.Split(input[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])
		lines = append(lines, &line{x1: x1, y1: y1, x2: x2, y2: y2})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	overlappingPts1, overlappingPts2 := 0, 0
	points1, points2 := map[point]int{}, map[point]int{}
	for _, line := range lines {
		switch {
		// Diagonal line.
		case line.x1 != line.x2 && line.y1 != line.y2:
			switch {
			case line.x1 < line.x2 && line.y1 < line.y2:
				for x, y := line.x1, line.y1; x <= line.x2; {
					if mark(points2, x, y) {
						overlappingPts2++
					}
					x++
					y++
				}
			case line.x1 > line.x2 && line.y1 > line.y2:
				for x, y := line.x2, line.y2; x <= line.x1; {
					if mark(points2, x, y) {
						overlappingPts2++
					}
					x++
					y++
				}
			case line.x1 < line.x2 && line.y1 > line.y2:
				for x, y := line.x1, line.y1; x <= line.x2; {
					if mark(points2, x, y) {
						overlappingPts2++
					}
					x++
					y--
				}
			default:
				for x, y := line.x2, line.y2; x <= line.x1; {
					if mark(points2, x, y) {
						overlappingPts2++
					}
					x++
					y--
				}
			}

		// Horizontal line.
		case line.x1 == line.x2:
			if line.y1 < line.y2 {
				for y := line.y1; y <= line.y2; y++ {
					if mark(points1, line.x1, y) {
						overlappingPts1++
					}
					if mark(points2, line.x1, y) {
						overlappingPts2++
					}
				}
				continue
			}
			for y := line.y2; y <= line.y1; y++ {
				if mark(points1, line.x1, y) {
					overlappingPts1++
				}
				if mark(points2, line.x1, y) {
					overlappingPts2++
				}
			}

		// Vertical line.
		default:
			if line.x1 < line.x2 {
				for x := line.x1; x <= line.x2; x++ {
					if mark(points1, x, line.y1) {
						overlappingPts1++
					}
					if mark(points2, x, line.y1) {
						overlappingPts2++
					}
				}
				continue
			}
			for x := line.x2; x <= line.x1; x++ {
				if mark(points1, x, line.y1) {
					overlappingPts1++
				}
				if mark(points2, x, line.y1) {
					overlappingPts2++
				}
			}
		}
	}
	fmt.Printf("Part 1: %d\n", overlappingPts1)
	fmt.Printf("Part 2: %d\n", overlappingPts2)
}
