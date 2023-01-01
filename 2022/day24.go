// This program implements the solution for
// https://adventofcode.com/2022/day/24.
package main

import (
	"fmt"
	"reflect"
	"strings"

	_ "embed"
)

//go:embed day24_input.txt
var input string

type valley struct {
	start, end           coord
	width, height        int
	precomputedBlizzards []blizzards
}

type coord struct {
	x, y int
}

// blizzard represents a blizzard in the valley. A blizzard has a direction and
// a current position in the valley.
type blizzard struct {
	dir int
	pos coord
}

// blizzards represents the current state of all blizzards in the valley. We use
// a map of coordinates to blizzards at a coordinate, because more than one
// blizzard may occupy a given coordinate.
type blizzards map[coord][]blizzard

// parse parses the text input that describes the valley and initial state of
// the blizzard into a valley struct. It precomputes all states of the blizzard.
func parse() valley {
	initialBlizzards := blizzards{}

	inputParts := strings.Split(input, "\n")
	for y, row := range inputParts {
		for x, col := range row {
			dir := 0
			switch col {
			case '^':
			case '>':
				dir = 1
			case 'v':
				dir = 2
			case '<':
				dir = 3
			default:
				continue
			}
			initialBlizzards[coord{x, y}] = append(initialBlizzards[coord{x, y}], blizzard{dir, coord{x, y}})
		}
	}

	w := len(inputParts[0])
	h := len(inputParts)
	return valley{
		start:                coord{strings.Index(inputParts[0], "."), 0},
		end:                  coord{strings.Index(inputParts[len(inputParts)-1], "."), len(inputParts) - 1},
		width:                w,
		height:               h,
		precomputedBlizzards: precomputeBlizzards(initialBlizzards, w, h),
	}
}

// precomputeBlizzards precomputes all states of a blizzards pattern given its
// initial state and the width and height of the valley the blizzards are in.
func precomputeBlizzards(initial blizzards, valleyWidth, valleyHeight int) []blizzards {
	precomputed := []blizzards{initial}
	cur := initial
	for {
		cur = next(cur, valleyWidth, valleyHeight)
		// This will work for valid input. Eventually the blizzard's pattern will
		// repeat (wrap), so we only need to store unique states in order.
		if reflect.DeepEqual(cur, precomputed[0]) {
			break
		}
		precomputed = append(precomputed, cur)
	}
	return precomputed
}

// next computes the next state of a set of blizzards in a valley of the
// specified width and height. It just simulates their movement pattern.
func next(cur blizzards, valleyWidth, valleyHeight int) blizzards {
	nextBlizzards := blizzards{}
	for _, blizzardsAtPos := range cur {
		for _, bliz := range blizzardsAtPos {
			dx, dy := delta(bliz.dir)
			newBliz := blizzard{dir: bliz.dir, pos: coord{bliz.pos.x + dx, bliz.pos.y + dy}}
			newBliz.pos = wrap(newBliz.pos, valleyWidth, valleyHeight)
			nextBlizzards[newBliz.pos] = append(nextBlizzards[newBliz.pos], newBliz)
		}
	}
	return nextBlizzards
}

// delta computes the x and y delta for a blizzard given a direction.
func delta(dir int) (int, int) {
	switch dir {
	case 0:
		return 0, -1
	case 1:
		return 1, 0
	case 2:
		return 0, 1
	case 3:
		return -1, 0
	}
	return 0, 0 // Will never get here with valid input.
}

// wrap wraps the blizzard's specified position if needed, otherwise it just
// returns the position.
func wrap(pos coord, valleyWidth, valleyHeight int) coord {
	if pos.x == 0 {
		pos.x = valleyWidth - 2
	}
	if pos.x == valleyWidth-1 {
		pos.x = 1
	}
	if pos.y == 0 {
		pos.y = valleyHeight - 2
	}
	if pos.y == valleyHeight-1 {
		pos.y = 1
	}
	return pos
}

// bfs find the shortest path from the valley's start point to its end point.
func bfs(valley valley, curTime int) int {
	q := []coord{valley.start}

	// This is the only tricky part of this BFS, we have a third dimension to
	// track visited nodes, which is the time. Normally for a 2D BFS it's just x
	// and y.
	type key struct {
		time int
		pos  coord
	}
	visited := map[key]any{}
	for time := curTime; len(q) > 0; time++ {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[i]
			if cur == valley.end {
				return time
			}
			if _, ok := visited[key{time, cur}]; ok {
				continue
			}
			visited[key{time, cur}] = nil

			// Determine all the next positions we can visit from our current
			// position. Valid next positions are the four directions around us and
			// waiting (not moving).
			for _, n := range []coord{{cur.x, cur.y}, {cur.x, cur.y - 1}, {cur.x + 1, cur.y}, {cur.x, cur.y + 1}, {cur.x - 1, cur.y}} {
				// Check OOB, but only if the next position is not the start or end.
				// Not the start because we may need to wait at the start before moving
				// out of the start if a blizzard is blocking us. Not the end because we
				// want to move into the end position.
				if n != valley.start && n != valley.end && (n.x <= 0 || n.x >= valley.width-1 || n.y <= 0 || n.y >= valley.height-1) {
					continue
				}
				// Precomputed blizzards only contains unique states in order, so we
				// need to mod our current time by its length to determine which
				// blizzards state we're on. Also time+1 because we are considering the
				// next blizzards state here.
				if _, ok := valley.precomputedBlizzards[(time+1)%len(valley.precomputedBlizzards)][n]; ok {
					continue
				}

				q = append(q, n)
			}
		}
		q = q[l:]
	}

	return 0 // Will never get here with valid input.
}

func main() {
	valley := parse()
	time := bfs(valley, 0)
	fmt.Printf("Part 1: %d\n", time)

	// For this, just swap the valley's start and end, then start the BFS at a
	// specified time. Then do it again.
	valley.start, valley.end = valley.end, valley.start
	time = bfs(valley, time)
	valley.start, valley.end = valley.end, valley.start
	time = bfs(valley, time)
	fmt.Printf("Part 2: %d\n", time)
}
