// This program implements the solution for
// https://adventofcode.com/2022/day/22.
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed day22_input.txt
var input string

// parseState parses the board portion of the input in a state struct. It parses
// the board into a string slice, padding rows as needed so every row is the
// same length, and finds your starting position.
func parseState() state {
	// Find the largest number of columns in the board, since shorter rows are not
	// right padded with spaces.
	var cols int
	lines := strings.Split(input, "\n")
	for _, row := range lines[:len(lines)-2] {
		if l := len(row); l > cols {
			cols = l
		}
	}

	// Store the board as a string slice, padding rows as necessary so they are
	// all the same width.
	var board []string
	for _, row := range lines[:len(lines)-2] {
		board = append(board, row+strings.Repeat(" ", cols-len(row)))
	}

	// Your starting position is the first empty space in the first row.
	startX := strings.Index(lines[0], ".")
	return state{board, startX, 0, 0}
}

// parsePath parses a path description string like "10R5L5R10L4R5L5" into a
// string slice.
func parsePath() []string {
	lines := strings.Split(input, "\n")
	pathStr := lines[len(lines)-1]

	var path []string
	num := ""
	for _, r := range pathStr {
		if r == 'L' || r == 'R' {
			if num != "" {
				path = append(path, num)
				num = ""
			}
			path = append(path, string(r))
			continue
		}
		num += string(r)
	}
	if num != "" {
		path = append(path, num)
	}

	return path
}

func main() {
	path := parsePath()
	state := parseState()

	for _, instruction := range path {
		steps, err := strconv.Atoi(instruction)
		if err == nil {
			state.move(steps)
			continue
		}

		state.rotate(instruction)
	}

	fmt.Printf("Part 1: %d\n", state.pwd())
	// fmt.Printf("Part 2: %d\n", )
}

type state struct {
	board           []string
	posX, posY, dir int
}

func (s *state) move(steps int) {
	for i := 0; i < steps; i++ {
		nextX, nextY := s.posX, s.posY
		switch s.dir {
		case 0:
			nextX++
			if nextX == len(s.board[0]) || s.board[nextY][nextX] == ' ' {
				nextX = s.indexInRow(nextY)
			}
		case 1:
			nextY++
			if nextY == len(s.board) || s.board[nextY][nextX] == ' ' {
				nextY = s.indexInCol(nextX)
			}
		case 2:
			nextX--
			if nextX == -1 || s.board[nextY][nextX] == ' ' {
				nextX = s.lastIndexInRow(nextY)
			}
		case 3:
			nextY--
			if nextY == -1 || s.board[nextY][nextX] == ' ' {
				nextY = s.lastIndexInCol(nextX)
			}
		}

		switch s.board[nextY][nextX] {
		case '.':
			s.posX = nextX
			s.posY = nextY
		case '#':
			return
		}
	}
}

func (s *state) indexInRow(rowIdx int) int {
	return int(math.Min(float64(strings.Index(s.board[rowIdx], ".")), float64(strings.Index(s.board[rowIdx], "#"))))
}

func (s *state) lastIndexInRow(rowIdx int) int {
	return int(math.Max(float64(strings.LastIndex(s.board[rowIdx], ".")), float64(strings.LastIndex(s.board[rowIdx], "#"))))
}

func (s *state) indexInCol(colIdx int) int {
	for rowIdx, row := range s.board {
		if col := row[colIdx]; col == '.' || col == '#' {
			return rowIdx
		}
	}
	// Will never get here with valid input.
	return 0
}

func (s *state) lastIndexInCol(colIdx int) int {
	for rowIdx := len(s.board) - 1; rowIdx >= 0; rowIdx-- {
		if col := s.board[rowIdx][colIdx]; col == '.' || col == '#' {
			return rowIdx
		}
	}
	// Will never get here with valid input.
	return 0
}

func (s *state) rotate(dir string) {
	if dir == "R" {
		s.dir++
		if s.dir > 3 {
			s.dir = 0
		}
		return
	}

	s.dir--
	if s.dir < 0 {
		s.dir = 3
	}
}

func (s *state) pwd() int {
	return (s.posY+1)*1000 + (s.posX+1)*4 + s.dir
}

func (s *state) print() {
	fmt.Printf("pos x: %d, pos y: %d, dir: %d\n", s.posX, s.posY, s.dir)
	for _, row := range s.board {
		fmt.Println(row)
	}
	fmt.Println()
}
