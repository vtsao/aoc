// This program implements the solution for
// https://adventofcode.com/2021/day/23.
//
// curl -b "$(cat .session)" -o day23_input.txt https://adventofcode.com/2021/day/23/input
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type amphipod struct {
	name         string
	room0, room1 coord
	energy       int
}

var amphipods = []amphipod{
	{
		name:   "A0",
		room0:  coord{i: 2, j: 3},
		room1:  coord{i: 3, j: 3},
		energy: 1,
	},
	{
		name:   "A1",
		room0:  coord{i: 2, j: 3},
		room1:  coord{i: 3, j: 3},
		energy: 1,
	},
	{
		name:   "B0",
		room0:  coord{i: 2, j: 5},
		room1:  coord{i: 3, j: 5},
		energy: 10,
	},
	{
		name:   "B1",
		room0:  coord{i: 2, j: 5},
		room1:  coord{i: 3, j: 5},
		energy: 10,
	},
	{
		name:   "C0",
		room0:  coord{i: 2, j: 7},
		room1:  coord{i: 3, j: 7},
		energy: 100,
	},
	{
		name:   "C1",
		room0:  coord{i: 2, j: 7},
		room1:  coord{i: 3, j: 7},
		energy: 100,
	},
	{
		name:   "D0",
		room0:  coord{i: 2, j: 9},
		room1:  coord{i: 3, j: 9},
		energy: 1000,
	},
	{
		name:   "D1",
		room0:  coord{i: 2, j: 9},
		room1:  coord{i: 3, j: 9},
		energy: 1000,
	},
}

type coord struct {
	i, j int
}

type state map[string]coord

func (s state) key() string {
	k := ""
	for _, a := range amphipods {
		k += fmt.Sprintf("%d%d", s[a.name].i, s[a.name].j)
	}
	return k
}

type burrow struct {
	layout [][]rune
	state  state
}

func (b burrow) occupied(c coord) bool {
	if b.layout[c.i][c.j] == '#' {
		return true
	}
	for _, a := range amphipods {
		if c == b.state[a.name] {
			return true
		}
	}
	return false
}

func final(b burrow) bool {
	for _, a := range amphipods {
		if (b.state[a.name] != a.room0) && (b.state[a.name] != a.room1) {
			return false
		}
	}
	return true
}

func hallway(s state) string {
	for _, a := range amphipods {
		if (s[a.name] == coord{i: 1, j: 3}) || (s[a.name] == coord{i: 1, j: 5}) || (s[a.name] == coord{i: 1, j: 7}) || (s[a.name] == coord{i: 1, j: 9}) {
			return a.name
		}
	}
	return ""
}

func neighbors(c coord) []coord {
	return []coord{{i: c.i - 1, j: c.j}, {i: c.i + 1, j: c.j}, {i: c.i, j: c.j - 1}, {i: c.i, j: c.j + 1}}
}

func other(amphipod string) string {
	if idx := amphipod[1]; idx == '0' {
		return fmt.Sprintf("%s1", string(amphipod[0]))
	}
	return fmt.Sprintf("%s0", string(amphipod[0]))
}

func move(b burrow, visited map[string]struct{}, cache map[string]int) int {
	if final(b) {
		return 0
	}
	if cached, ok := cache[b.state.key()]; ok {
		return cached
	}
	if _, ok := visited[b.state.key()]; ok {
		return math.MaxInt64
	}

	visited[b.state.key()] = struct{}{}
	energy := math.MaxInt64
	hallway := hallway(b.state)
	for _, a := range amphipods {
		if hallway != "" && hallway != a.name {
			continue
		}
		for _, c := range neighbors(b.state[a.name]) {
			if b.occupied(c) {
				continue
			}

			// If the coordinate we're trying to move to isn't in the hallway and
			// isn't the amphipod's own room, that means we're trying to move into
			// another amphipod's room, which isn't allowed.
			if c.i != 1 && (c != a.room0 || c != a.room1) {
				continue
			}

			if c == a.room0 && b.occupied(a.room1) && b.state[other(a.name)] != a.room1 {
				continue
			}

			old := b.state[a.name]
			b.state[a.name] = c
			e := move(b, visited, cache)
			if e != math.MaxInt64 {
				if e = e + a.energy; e < energy {
					energy = e
				}
			}
			b.state[a.name] = old
		}
	}

	cache[b.state.key()] = energy
	return energy
}

func parseInput() burrow {
	file, _ := os.Open("day23_input.txt")
	defer file.Close()

	b := burrow{state: map[string]coord{}}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		var row []rune
		for j, col := range scanner.Text() {
			row = append(row, col)
			if col == '.' || col == '#' {
				continue
			}
			if _, ok := b.state[string(col)+"0"]; !ok {
				b.state[string(col)+"0"] = coord{i: i, j: j}
				continue
			}
			b.state[string(col)+"1"] = coord{i: i, j: j}
		}
		b.layout = append(b.layout, row)
	}

	return b
}

func main() {
	burrow := parseInput()
	energy := move(burrow, map[string]struct{}{}, map[string]int{})
	fmt.Printf("Part 1: %d\n", energy)

	// fmt.Printf("Part 2: %d\n", )
}
