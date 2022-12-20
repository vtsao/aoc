// This program implements the solution for
// https://adventofcode.com/2022/day/17.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("day17_input.txt")

	var jetPattern string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jetPattern = scanner.Text()
	}

	floor := fall(jetPattern)
	fmt.Printf("Part 1: %d\n", floor)
}

type coord struct {
	x, y int
}

func fall(jetPattern string) int {
	floor := 0
	shapeType := 0
	rocks := map[coord]any{}
	jetIdx := 0
	for i := 0; i < 20220; i++ {
		if shapeType == 0 && jetIdx == 0 {
			log.Print(floor)
		}
		coords := shape(floor, shapeType)
		top, jIdx := move(floor, coords, jetIdx, jetPattern, rocks)
		if top > floor {
			floor = top
		}

		jetIdx = jIdx
		shapeType++
		if shapeType == 5 {
			shapeType = 0
		}

		var rockC []coord
		for r := range rocks {
			rockC = append(rockC, r)
		}
		// draw(rockC)
	}

	return floor
}

func shape(floor, shapeType int) []coord {
	switch shapeType {
	case 0:
		return []coord{
			{2, floor + 4},
			{3, floor + 4},
			{4, floor + 4},
			{5, floor + 4},
		}
	case 1:
		return []coord{
			{3, floor + 6},
			{2, floor + 5},
			{3, floor + 5},
			{4, floor + 5},
			{3, floor + 4},
		}
	case 2:
		return []coord{
			{4, floor + 6},
			{4, floor + 5},
			{2, floor + 4},
			{3, floor + 4},
			{4, floor + 4},
		}
	case 3:
		return []coord{
			{2, floor + 7},
			{2, floor + 6},
			{2, floor + 5},
			{2, floor + 4},
		}
	case 4:
		return []coord{
			{2, floor + 4},
			{2, floor + 5},
			{3, floor + 5},
			{3, floor + 4},
		}
	}

	// Will never get here with valid input.
	return nil
}

func move(floor int, shape []coord, jetIdx int, jetPattern string, rocks map[coord]any) (int, int) {
	cur := shape
	jet := true
	for {
		// draw(cur)

		var next []coord
		for _, c := range cur {
			if jet {
				if jetPattern[jetIdx] == '>' {
					nextC := coord{c.x + 1, c.y}
					if _, ok := rocks[nextC]; ok || nextC.x >= 7 {
						next = cur
						break
					}
					next = append(next, nextC)
					continue
				}

				nextC := coord{c.x - 1, c.y}
				if _, ok := rocks[nextC]; ok || nextC.x < 0 {
					next = cur
					break
				}
				next = append(next, nextC)
				continue
			}

			nextC := coord{c.x, c.y - 1}
			if _, ok := rocks[nextC]; ok || nextC.y <= 0 {
				next = nil
				break
			}
			next = append(next, nextC)
		}

		if jet {
			jetIdx++
			if jetIdx == len(jetPattern) {
				jetIdx = 0
			}
		}
		jet = !jet

		if len(next) != len(cur) {
			break
		}
		cur = next
	}

	// draw(cur)

	top := 0
	for _, c := range cur {
		rocks[c] = nil
		if c.y > top {
			top = c.y
		}
	}
	return top, jetIdx
}

func draw(coords []coord) {
	cave := make([][]string, 10)
	for i := range cave {
		cave[i] = make([]string, 7)
	}

	for _, c := range coords {
		cave[c.y][c.x] = "#"
	}

	for i := len(cave) - 1; i >= 0; i-- {
		row := ""
		for _, col := range cave[i] {
			if col == "#" {
				row += col
				continue
			}
			row += "."
		}
		fmt.Println(row)
	}
	fmt.Println()
}
