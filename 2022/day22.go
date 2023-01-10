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

// wrapMap hardcodes the board wrapping for my specific input if the board were
// treated as a cube. This is messy and not ideal.
func wrapMap() map[coordAndDir]coordAndDir {
	edgeMappings := []struct {
		fromEdge, toEdge                     interval
		dir, nextDir, otherDir, nextOtherDir int
	}{
		{
			fromEdge:     interval{coord{50, 0}, coord{50, 49}},
			toEdge:       interval{coord{0, 149}, coord{0, 100}},
			dir:          2,
			nextDir:      0,
			otherDir:     2,
			nextOtherDir: 0,
		},
		{
			fromEdge:     interval{coord{50, 0}, coord{99, 0}},
			toEdge:       interval{coord{0, 150}, coord{0, 199}},
			dir:          3,
			nextDir:      0,
			otherDir:     2,
			nextOtherDir: 1,
		},
		{
			fromEdge:     interval{coord{100, 0}, coord{149, 0}},
			toEdge:       interval{coord{0, 199}, coord{49, 199}},
			dir:          3,
			nextDir:      3,
			otherDir:     1,
			nextOtherDir: 1,
		},
		{
			fromEdge:     interval{coord{149, 0}, coord{149, 49}},
			toEdge:       interval{coord{99, 149}, coord{99, 100}},
			dir:          0,
			nextDir:      2,
			otherDir:     0,
			nextOtherDir: 2,
		},
		{
			fromEdge:     interval{coord{100, 49}, coord{149, 49}},
			toEdge:       interval{coord{99, 50}, coord{99, 99}},
			dir:          1,
			nextDir:      2,
			otherDir:     0,
			nextOtherDir: 3,
		},
		{
			fromEdge:     interval{coord{50, 149}, coord{99, 149}},
			toEdge:       interval{coord{49, 150}, coord{49, 199}},
			dir:          1,
			nextDir:      2,
			otherDir:     0,
			nextOtherDir: 3,
		},
		{
			fromEdge:     interval{coord{0, 100}, coord{49, 100}},
			toEdge:       interval{coord{50, 50}, coord{50, 99}},
			dir:          3,
			nextDir:      0,
			otherDir:     2,
			nextOtherDir: 1,
		},
	}

	wMap := map[coordAndDir]coordAndDir{}
	for _, em := range edgeMappings {
		// Precondition is both intervals are the same length.
		toPoints := em.toEdge.points()
		for i, fromPoint := range em.fromEdge.points() {
			wMap[coordAndDir{coord: fromPoint, dir: em.dir}] = coordAndDir{coord: toPoints[i], dir: em.nextDir}
			wMap[coordAndDir{coord: toPoints[i], dir: em.otherDir}] = coordAndDir{coord: fromPoint, dir: em.nextOtherDir}
		}
	}
	return wMap
}

type coord struct {
	x, y int
}

type coordAndDir struct {
	coord
	dir int
}

// interval represents an interval from a starting coordinate to an end
// coordinate. It's assumed that an interval is a vertical or horizontal line
// (i.e., not a diagonal line). Start and end are both inclusive.
type interval struct {
	start, end coord
}

// points returns all the points along the interval. This only makes sense for
// vertical or horizontal lines, which is a precondition for the interval
// struct.
func (i *interval) points() []coord {
	var p []coord

	// Horizontal line.
	if i.start.y == i.end.y {
		if i.end.x > i.start.x {
			for x := i.start.x; x <= i.end.x; x++ {
				p = append(p, coord{x, i.start.y})
			}
			return p
		}
		for x := i.start.x; x >= i.end.x; x-- {
			p = append(p, coord{x, i.start.y})
		}
		return p
	}

	// Vertical line.
	if i.end.y > i.start.y {
		for y := i.start.y; y <= i.end.y; y++ {
			p = append(p, coord{i.start.x, y})
		}
		return p
	}
	for y := i.start.y; y >= i.end.y; y-- {
		p = append(p, coord{i.start.x, y})
	}
	return p
}

// state holds the board configuration, where we currently are on the board, and
// how the board wrapping works when viewed as a cube.
type state struct {
	board           []string
	posX, posY, dir int
	wrapMap         map[coordAndDir]coordAndDir
}

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
	return state{board, startX, 0, 0, wrapMap()}
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

// move moves us in the direction we are currently facing until we hit a wall.
// Treats the board as 2D with wrap around to the other side.
func (s *state) move(steps int) {
	for i := 0; i < steps; i++ {
		nextX, nextY, nextDir := s.posX, s.posY, s.dir

		switch s.dir {
		case 0:
			nextX++
			if nextX == len(s.board[0]) || s.board[nextY][nextX] == ' ' {
				// Find the first x index of an empty space or a wall in this row.
				nextX = int(math.Min(float64(strings.Index(s.board[s.posY], ".")), float64(strings.Index(s.board[s.posY], "#"))))
			}
		case 1:
			nextY++
			if nextY == len(s.board) || s.board[nextY][nextX] == ' ' {
				// Find the first y index of an empty space or a wall in this column.
				for rowIdx, row := range s.board {
					if col := row[s.posX]; col == '.' || col == '#' {
						nextY = rowIdx
						break
					}
				}
			}
		case 2:
			nextX--
			if nextX == -1 || s.board[nextY][nextX] == ' ' {
				// Find the last x index of an empty space or a wall in this row.
				nextX = int(math.Max(float64(strings.LastIndex(s.board[s.posY], ".")), float64(strings.LastIndex(s.board[s.posY], "#"))))
			}
		case 3:
			nextY--
			if nextY == -1 || s.board[nextY][nextX] == ' ' {
				// Find the last y index of an empty space or a wall in this column.
				for rowIdx := len(s.board) - 1; rowIdx >= 0; rowIdx-- {
					if col := s.board[rowIdx][s.posX]; col == '.' || col == '#' {
						nextY = rowIdx
						break
					}
				}
			}
		}

		switch s.board[nextY][nextX] {
		case '.':
			s.posX = nextX
			s.posY = nextY
			s.dir = nextDir
		case '#':
			return
		}
	}
}

// move moves us in the direction we are currently facing until we hit a wall.
// Treats the board as a cube.
func (s *state) moveForCube(steps int) {
	for i := 0; i < steps; i++ {
		nextX, nextY, nextDir := s.posX, s.posY, s.dir

		// If we're going to "walk off the edge", find the next coordinate we should
		// move to on the cube.
		next, ok := s.wrapMap[coordAndDir{coord: coord{s.posX, s.posY}, dir: s.dir}]
		if ok {
			nextX, nextY, nextDir = next.x, next.y, next.dir
		} else {
			switch s.dir {
			case 0:
				nextX++
			case 1:
				nextY++
			case 2:
				nextX--
			case 3:
				nextY--
			}
		}

		switch s.board[nextY][nextX] {
		case '.':
			s.posX = nextX
			s.posY = nextY
			s.dir = nextDir
		case '#':
			return
		}
	}
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

func (s *state) followPath(path []string, isCube bool) {
	for _, instruction := range path {
		steps, err := strconv.Atoi(instruction)
		if err != nil {
			s.rotate(instruction)
			continue
		}

		if isCube {
			s.moveForCube(steps)
			continue
		}
		s.move(steps)
	}
}

func main() {
	path := parsePath()

	state := parseState()
	state.followPath(path, false)
	fmt.Printf("Part 1: %d\n", state.pwd())

	state = parseState()
	state.followPath(path, true)
	fmt.Printf("Part 2: %d\n", state.pwd())
}
