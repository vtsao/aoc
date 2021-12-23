// This program implements the solution for
// https://adventofcode.com/2021/day/23.
//
// curl -b "$(cat .session)" -o day23_input.txt https://adventofcode.com/2021/day/23/input
package main

import (
	"bufio"
	"math"
	"os"
)

type coord struct {
	i, j int
}

func parseInput() ([][]rune, coord, coord, coord, coord, coord, coord, coord, coord) {
	file, _ := os.Open("day23_input.txt")
	defer file.Close()

	var a0, a1, b0, b1, c0, c1, d0, d1 coord
	var burrow [][]rune
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		var row []rune

		for j, col := range scanner.Text() {
			row = append(row, col)
			switch col {
			case 'A':
				if a0.i == 0 && a0.j == 0 {
					a0 = coord{i: i, j: j}
					continue
				}
				a1 = coord{i: i, j: j}
			case 'B':
				if b0.i == 0 && b0.j == 0 {
					b0 = coord{i: i, j: j}
					continue
				}
				b1 = coord{i: i, j: j}
			case 'C':
				if c0.i == 0 && c0.j == 0 {
					c0 = coord{i: i, j: j}
					continue
				}
				c1 = coord{i: i, j: j}
			case 'D':
				if d0.i == 0 && d0.j == 0 {
					d0 = coord{i: i, j: j}
					continue
				}
				d1 = coord{i: i, j: j}
			}
		}

		burrow = append(burrow, row)
	}

	return burrow, a0, a1, b0, b1, c0, c1, d0, d1
}

func occupied(burrow [][]rune, c, a0, a1, b0, b1, c0, c1, d0, d1 coord) bool {
	if burrow[c.i][c.j] == '#' {
		return true
	}

	if c == a0 || c == a1 || c == b0 || c == b1 || c == c0 || c == c1 || c == d0 || c == d1 {
		return true
	}

	return false
}

func move(burrow [][]rune, a0, a1, b0, b1, c0, c1, d0, d1 coord) int {
	if (a0 == coord{i: 2, j: 3} || a0 == coord{i: 3, j: 3}) &&
		(a1 == coord{i: 2, j: 3} || a1 == coord{i: 3, j: 3}) &&
		(b0 == coord{i: 2, j: 5} || b0 == coord{i: 3, j: 5}) &&
		(b1 == coord{i: 2, j: 5} || b1 == coord{i: 3, j: 5}) &&
		(c0 == coord{i: 2, j: 7} || c0 == coord{i: 3, j: 7}) &&
		(c1 == coord{i: 2, j: 7} || c1 == coord{i: 3, j: 7}) &&
		(d0 == coord{i: 2, j: 9} || d0 == coord{i: 3, j: 9}) &&
		(d1 == coord{i: 2, j: 9} || d1 == coord{i: 3, j: 9}) {
		return 0

	}

	// if visited state, return 0

	energy := math.MaxInt64

	for _, c := range []coord{{i: a0.i - 1, j: a0.j}, {i: a0.i + 1, j: a0.j}, {i: a0.i, j: a0.j - 1}, {i: a0.i, j: a0.j + 1}} {
		if occupied(burrow, c, a0, a1, b0, b1, c0, c1, d0, d1) {
			continue
		}

		if (c == coord{i: 1, j: 3}) || (c == coord{i: 1, j: 5}) || (c == coord{i: 1, j: 7}) || (c == coord{i: 1, j: 9}) {
			if e := 1 + move(burrow, c, a1, b0, b1, c0, c1, d0, d1); e < energy {
				energy = e
			}
			continue
		}

		if (c == coord{i: 3, j: 3}) {
			if e := 1 + move(burrow, c, a1, b0, b1, c0, c1, d0, d1); e < energy {
				energy = e
			}
			continue
		}

		if (c == coord{i: 2, j: 3}) &&
			(a1 == coord{i: 3, j: 3}) || !occupied(burrow, coord{i: 3, j: 3}, a0, a1, b0, b1, c0, c1, d0, d1) {
			if e := 1 + move(burrow, c, a1, b0, b1, c0, c1, d0, d1); e < energy {
				energy = e
			}
			continue
		}
	}

	return energy
}

func main() {

	// burrow, a0, a1, b0, b1, c0, c1, d0, d1 := parseInput()

	// fmt.Printf("Part 1: %d\n", )
	// fmt.Printf("Part 2: %d\n", )
}
