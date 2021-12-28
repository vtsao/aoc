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

type coord struct {
	i, j int
}

type amphipod struct {
	color  string
	names  []string
	rooms  []coord
	energy int
}

var amphipods1 = []amphipod{
	{
		color:  "A",
		names:  []string{"A0", "A1"},
		rooms:  []coord{{2, 3}, {3, 3}},
		energy: 1,
	},
	{
		color:  "B",
		names:  []string{"B0", "B1"},
		rooms:  []coord{{2, 5}, {3, 5}},
		energy: 10,
	},
	{
		color:  "C",
		names:  []string{"C0", "C1"},
		rooms:  []coord{{2, 7}, {3, 7}},
		energy: 100,
	},
	{
		color:  "D",
		names:  []string{"D0", "D1"},
		rooms:  []coord{{2, 9}, {3, 9}},
		energy: 1000,
	},
}

var amphipods2 = []amphipod{
	{
		color:  "A",
		names:  []string{"A0", "A1", "A2", "A3"},
		rooms:  []coord{{2, 3}, {3, 3}, {4, 3}, {5, 3}},
		energy: 1,
	},
	{
		color:  "B",
		names:  []string{"B0", "B1", "B2", "B3"},
		rooms:  []coord{{2, 5}, {3, 5}, {4, 5}, {5, 5}},
		energy: 10,
	},
	{
		color:  "C",
		names:  []string{"C0", "C1", "C2", "C3"},
		rooms:  []coord{{2, 7}, {3, 7}, {4, 7}, {5, 7}},
		energy: 100,
	},
	{
		color:  "D",
		names:  []string{"D0", "D1", "D2", "D3"},
		rooms:  []coord{{2, 9}, {3, 9}, {4, 9}, {5, 9}},
		energy: 1000,
	},
}

type state map[string]coord

func stateKey(s state, amphipods []amphipod) string {
	k := ""
	for _, a := range amphipods {
		for _, aName := range a.names {
			k += fmt.Sprintf("%d%d", s[aName].i, s[aName].j)
		}
	}
	return k
}

func occupied(s state, amphipods []amphipod, c coord) bool {
	for _, a := range amphipods {
		for _, n := range a.names {
			if c == s[n] {
				return true
			}
		}
	}
	return false
}

// canMove determines whether the specified amphipod can move to the target
// coordinate. If it can, it returns the energy cost of the move.
func canMove(s state, amphipods []amphipod, aName string, energy int, c coord) (int, bool) {
	og := s[aName]
	defer func() {
		s[aName] = og
	}()

	cost := 0
	if s[aName].i < c.i {
		for s[aName].j < c.j {
			next := coord{s[aName].i, s[aName].j + 1}
			if occupied(s, amphipods, next) {
				return 0, false
			}
			cost += energy
			s[aName] = next
		}
		for s[aName].j > c.j {
			next := coord{s[aName].i, s[aName].j - 1}
			if occupied(s, amphipods, next) {
				return 0, false
			}
			cost += energy
			s[aName] = next
		}
		for s[aName].i < c.i {
			next := coord{s[aName].i + 1, s[aName].j}
			if occupied(s, amphipods, next) {
				return 0, false
			}
			cost += energy
			s[aName] = next
		}

		return cost, true
	}

	for s[aName].i > c.i {
		next := coord{s[aName].i - 1, s[aName].j}
		if occupied(s, amphipods, next) {
			return 0, false
		}
		cost += energy
		s[aName] = next
	}
	for s[aName].j < c.j {
		next := coord{s[aName].i, s[aName].j + 1}
		if occupied(s, amphipods, next) {
			return 0, false
		}
		cost += energy
		s[aName] = next
	}
	for s[aName].j > c.j {
		next := coord{s[aName].i, s[aName].j - 1}
		if occupied(s, amphipods, next) {
			return 0, false
		}
		cost += energy
		s[aName] = next
	}

	return cost, true
}

// subRoomsAfter determines whether all the sub rooms below a given amphipod's
// room are occupied by amphipods of the same color.
func subRoomsAfter(s state, a amphipod, room int) bool {
rooms:
	for i := room + 1; i < len(a.rooms); i++ {
		for _, n := range a.names {
			if a.rooms[i] == s[n] {
				continue rooms
			}
		}
		return false
	}
	return true
}

// finalPos determines if the amphipod is in its final resting place. An
// amphipod is in its final resting place if it's in its own room and all the
// subrooms below it are occupied by other amphipods of the same color.
func finalPos(s state, aName string, a amphipod) bool {
	for i := len(a.rooms) - 1; i >= 0; i-- {
		if s[aName] == a.rooms[i] && subRoomsAfter(s, a, i) {
			return true
		}
	}
	return false
}

type validMove struct {
	c    coord
	cost int
}

// validMoves determines the set of valid moves for the specified amphipod that
// can take given this state.
func validMoves(s state, amphipods []amphipod, aName string, a amphipod, cache map[string][]validMove) []validMove {
	if cached, ok := cache[fmt.Sprintf("%s%d%d", stateKey(s, amphipods), s[aName].i, s[aName].j)]; ok {
		return cached
	}

	// We cache the set of valid moves for this amphipod given this state, since
	// we might arrive at this again.
	var validMoves []validMove
	defer func() {
		cache[fmt.Sprintf("%s%d%d", stateKey(s, amphipods), s[aName].i, s[aName].j)] = validMoves
	}()

	// The amphipod is in the hallway and its room is not occupied (or the room is
	// only occupied by its counterparts), we can try to move it into its room.
	if s[aName].i == 1 {
		for i := len(a.rooms) - 1; i >= 0; i-- {
			if !occupied(s, amphipods, a.rooms[i]) && subRoomsAfter(s, a, i) {
				if cost, ok := canMove(s, amphipods, aName, a.energy, a.rooms[i]); ok {
					validMoves = append(validMoves, validMove{c: a.rooms[i], cost: cost})
				}
			}
		}
		return validMoves
	}

	// The amphipod is in a room, but not its final resting place, we can try to
	// move it into the hallway. Note it can only move into specific spots in the
	// hallway (i.e., not in spots immediately outside a room).
	if !finalPos(s, aName, a) {
		for _, c := range []coord{{1, 1}, {1, 2}, {1, 4}, {1, 6}, {1, 8}, {1, 10}, {1, 11}} {
			if occupied(s, amphipods, c) {
				continue
			}
			if cost, ok := canMove(s, amphipods, aName, a.energy, c); ok {
				validMoves = append(validMoves, validMove{c: c, cost: cost})
			}
		}
	}

	// Otherwise there are no valid moves for the amphipod in this state.
	return validMoves
}

// final checks to see whether all the amphipods are in the organized state.
func final(s state, amphipods []amphipod) bool {
	for _, a := range amphipods {
		for _, aName := range a.names {
			if !finalPos(s, aName, a) {
				return false
			}
		}
	}
	return true
}

// move performs DFS to find the least amount of energy it takes to optimize the
// amphipods.
//
// This only runs in a reasonable time because of all the rules set in the
// problem that allows this DFS to be pruned significantly.
func move(s state, amphipods []amphipod, visited map[string]int, energySoFar int, cache map[string][]validMove) int {
	if final(s, amphipods) {
		return energySoFar
	}

	// If we reach a state we've seen before, but with a higher energy cost, stop
	// searching this path.
	if e, ok := visited[stateKey(s, amphipods)]; ok && energySoFar >= e {
		return math.MaxInt64
	}
	visited[stateKey(s, amphipods)] = energySoFar

	// DFS based on the valid moves we can take at this state.
	energy := math.MaxInt64
	for _, a := range amphipods {
		for _, aName := range a.names {
			for _, validMove := range validMoves(s, amphipods, aName, a, cache) {
				old := s[aName]
				s[aName] = validMove.c
				if e := move(s, amphipods, visited, energySoFar+validMove.cost, cache); e < energy {
					energy = e
				}
				s[aName] = old
			}
		}
	}

	return energy
}

func parseBurrow(burrow []string) state {
	s := state{}
	counts := map[rune]int{'A': 0, 'B': 0, 'C': 0, 'D': 0}
	for i, row := range burrow {
		for j, col := range row {
			if col != ' ' && col != '#' && col != '.' {
				s[fmt.Sprintf("%s%d", string(col), counts[col])] = coord{i, j}
				counts[col]++
			}
		}
	}
	return s
}

func parseInput() (state, state) {
	file, _ := os.Open("day23_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var burrow1 []string
	for scanner.Scan() {
		line := scanner.Text()
		burrow1 = append(burrow1, line)
	}
	var burrow2 []string
	burrow2 = append(burrow2, burrow1[:3]...)
	burrow2 = append(burrow2, "  #D#C#B#A#  ", "  #D#B#A#C#  ")
	burrow2 = append(burrow2, burrow1[3:]...)

	return parseBurrow(burrow1), parseBurrow(burrow2)
}

func main() {
	s1, s2 := parseInput()
	fmt.Printf("Part 1: %d\n", move(s1, amphipods1, map[string]int{}, 0, map[string][]validMove{}))
	fmt.Printf("Part 2: %d\n", move(s2, amphipods2, map[string]int{}, 0, map[string][]validMove{}))
}
