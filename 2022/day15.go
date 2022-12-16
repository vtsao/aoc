// This program implements the solution for
// https://adventofcode.com/2022/day/15.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

// parse parses and returns the coordinates of sensor and beacon from a string
// like
// "Sensor at x=2692921, y=2988627: closest beacon is at x=2453611, y=3029623".
func parse(line string) (coord, coord) {
	lineParts := strings.Split(line, ": ")
	sensorCoordStr := strings.TrimPrefix(lineParts[0], "Sensor at ")
	beaconCoordStr := strings.TrimPrefix(lineParts[1], "closest beacon is at ")
	return parseCoord(sensorCoordStr), parseCoord(beaconCoordStr)
}

// parseCoord parses and returns the coordinates from a string like
// "x=2692921, y=2988627".
func parseCoord(c string) coord {
	cParts := strings.Split(c, ", ")
	xStr := strings.Split(cParts[0], "=")
	x, _ := strconv.Atoi(xStr[1])
	yStr := strings.Split(cParts[1], "=")
	y, _ := strconv.Atoi(yStr[1])
	return coord{x, y}
}

func main() {
	file, _ := os.Open("day15_input.txt")

	sensors := map[coord]sensor{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sc, bc := parse(scanner.Text())
		sensors[sc] = sensor{sc, bc}
	}

	// fmt.Printf("Part 1: %d\n", cannotBeBeacons(2000000, sensors))

	//	const max = 4000000
	//	var beacon coord
	//
	// outer:
	//
	//	for y := 0; y < max; y++ {
	//		log.Print(y)
	//		cov := covered(y, 0, max, false, sensors)
	//		if len(cov) < max+1 {
	//			for x := 0; x < max; x++ {
	//				beacon = coord{x, y}
	//				if _, ok := cov[beacon]; !ok {
	//					break outer
	//				}
	//			}
	//		}
	//	}

	mustBeBeacon(0, 4000000, sensors)
	// fmt.Printf("Part 2: %d\n", lightMyWay.x*4000000+lightMyWay.y)
}

func cannotBeBeacons(y int, sensors map[coord]sensor) int {
	count := 0
	for x := -9999999; x < 9999999; x++ {
		if covered(coord{x, y}, false, sensors) {
			count++
		}
	}
	return count
}

func mustBeBeacon(min, max int, sensors map[coord]sensor) {
	var wg sync.WaitGroup

	for y := min; y <= max; y++ {
		y := y
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := min; x <= max; x++ {
				c := coord{x, y}
				if !covered(c, true, sensors) {
					fmt.Printf("Part 2: %d\n", c.x*4000000+c.y)
					os.Exit(0)
				}
			}
		}(y)
	}

	wg.Wait()
}

// for _, dx := range []int{-1, 1} {
// 	for x := s.x; ; x += dx {
// 		cur := coord{x, y}
// 		if s.dist(cur) > maxDist || x < xMin || x > xMax {
// 			break
// 		}
// 		if includeBeacons && cur == s.beacon {
// 			continue
// 		}
// 		cov[cur] = nil
// 	}
// }

func covered(c coord, includeBeacons bool, sensors map[coord]sensor) bool {
	for _, s := range sensors {
		coverageDist := s.dist(s.beacon)
		if c == s.beacon {
			return includeBeacons
		}
		if s.dist(c) <= coverageDist {
			return true
		}
	}
	return false
}

type sensor struct {
	coord
	beacon coord
}

type coord struct {
	x, y int
}

// dist returns the Manhattan distance between the src and the dest coordinates.
func (src *coord) dist(dest coord) int {
	x := int(math.Abs(float64(src.x) - float64(dest.x)))
	y := int(math.Abs(float64(src.y) - float64(dest.y)))
	return x + y
}

// func countCovered(y int, covered map[coord]any) int {
// 	count := 0
// 	for c := range covered {
// 		if c.y == y {
// 			count++
// 		}
// 	}
// 	return count
// }

// func markCovered(sensor, beacon coord, y int, covered map[coord]any) {
// 	maxDist := sensor.dist(beacon)

// 	visited := map[coord]any{}
// 	curDist := 0
// 	queue := []coord{sensor}
// 	for len(queue) > 0 {
// 		curLevelLen := len(queue)
// 		for i := 0; i < curLevelLen; i++ {
// 			cur := queue[i]

// 			if cur.x < 0 || cur.y >= 4000000 {
// 				continue
// 			}

// 			if curDist > maxDist {
// 				continue
// 			}

// 			if _, ok := visited[cur]; ok {
// 				continue
// 			}
// 			visited[cur] = nil
// 			if cur != beacon {
// 				covered[cur] = nil
// 			}

// 			for _, n := range neighbors(cur, y) {
// 				queue = append(queue, n)
// 			}
// 		}
// 		queue = queue[curLevelLen:]
// 		curDist++
// 	}
// }

// func neighbors(cur coord, y int) []coord {
// 	if cur.y == y {
// 		return []coord{
// 			{cur.x - 1, cur.y},
// 			{cur.x + 1, cur.y},
// 		}
// 	}

// 	dir := 1
// 	if cur.y > y {
// 		dir = -1
// 	}
// 	return []coord{
// 		{cur.x, cur.y + dir},
// 	}
// }
