// This program implements the solution for
// https://adventofcode.com/2022/day/18.
package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed day18_input.txt
var input string

// parse parses a string that represents a cube, like "5,6,15" into a cube
// struct.
func parse(cubeStr string) cube {
	c := strings.Split(cubeStr, ",")
	x, _ := strconv.Atoi(c[0])
	y, _ := strconv.Atoi(c[1])
	z, _ := strconv.Atoi(c[2])
	return cube{x, y, z}
}

func main() {
	cubes := map[cube]any{}
	for _, cubeStr := range strings.Split(input, "\n") {
		cubes[parse(cubeStr)] = nil
	}

	surfaceArea, extSurfaceArea := exposed(cubes)
	fmt.Printf("Part 1: %d\n", surfaceArea)
	fmt.Printf("Part 2: %d\n", extSurfaceArea)
}

type cube struct {
	x, y, z int
}

func (c cube) neighbors() []cube {
	return []cube{
		{c.x - 1, c.y, c.z},
		{c.x + 1, c.y, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y + 1, c.z},
		{c.x, c.y, c.z - 1},
		{c.x, c.y, c.z + 1},
	}
}

// exposed returns the surface area of the specified cubes that are exposed (the
// number of cube surfaces with no adjacent cube). The second return value
// returns only the exterior exposed surface area.
//
// In order to do this we calculate the neighbors of each cube and see if the
// coordinates of the neighbor is another cube. If it's not, then that
// particular surface is exposed. It's easier to visualize this in 2D, and then
// just apply it to 3D.
//
// For example:
//
//	       (2, 4)
//	(1, 3)        (3, 3)
//	(1, 2)        (3, 2)
//	       (2, 1)
func exposed(cubes map[cube]any) (int, int) {
	// count keeps track of how many exposed cube surfaces there are.
	count := 0
	// potentiallyTrapped tracks coordinates that are potentially part of a
	// trapped area inside the lava droplet. Candidates are coordinates beside a
	// cube's surface that are exposed.
	potentiallyTrapped := map[cube]any{}

	for c := range cubes {
		for _, n := range c.neighbors() {
			// If there is no cube beside this particular surface, then the surface is
			// exposed. This coordinate is also a candidate of a potentially trapped
			// area inside the lava droplet.
			if _, ok := cubes[n]; !ok {
				count++
				potentiallyTrapped[n] = nil
			}
		}
	}

	trapped := 0
	visited := map[cube]any{}
	// For each potentially trapped coordinate, BFS it to see if the BFS ends
	// without "timing out". If it ends w/o timing out, then this is a trapped
	// area.
	for c := range potentiallyTrapped {
		trapped += bfs(c, cubes, visited, 50)
	}

	return count, count - trapped
}

// bfs does BFS from a specified coordinate until it either determines that the
// start coordinate is part of a trapped area inside the lava droplet, and then
// returns the trapped surface area.
//
// Because coordinates on the exterior of the lava droplet will never end on a
// BFS, we have a threshold that acts as a "timeout". If BFS reaches this
// timeout, then we assume we're not BFS an interior area. This is a heuristic
// that we trialed and errored to find (i.e., kept increasing until the answer
// no longer changed). I'm not sure how to do it w/o this heuristic/hack.
func bfs(start cube, cubes, visited map[cube]any, threshold int) int {
	q := []cube{start}

	trapped := 0
	for i := 0; len(q) > 0; i++ {
		if i == threshold {
			return 0
		}

		l := len(q)
		for j := 0; j < l; j++ {
			c := q[j]
			if _, ok := visited[c]; ok {
				continue
			}
			visited[c] = nil

			for _, n := range c.neighbors() {
				if _, ok := cubes[n]; ok {
					trapped++
					continue
				}
				q = append(q, n)
			}
		}

		q = q[l:]
	}

	return trapped
}
