// This program implements the solution for
// https://adventofcode.com/2022/day/15.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

	intervals := coverage(2000000, math.MinInt, math.MaxInt, sensors)
	covered := 0
	for _, interval := range intervals {
		// End interval is not inclusive, so subtract 1.
		covered += interval.end - interval.start - 1
	}
	fmt.Printf("Part 1: %d\n", covered)

	for y := 0; y < 4000000; y++ {
		intervals := coverage(y, 0, 4000000, sensors)
		// Only one of these rows will have multiple intervals, all other rows
		// should have a single interval from 0 to 4000000 if the input is valid.
		//
		// We also assume the missing beacon isn't at x coordinate 0 or 4000000, but
		// in the middle somewhere.
		if len(intervals) == 2 {
			x := intervals[0].end
			fmt.Printf("Part 2: %d\n", x*4000000+y)
			return
		}
	}
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

type interval struct {
	start, end int
}

// coverage calculates the intervals (x coordinates) covered by all sensors for
// the specified row.
func coverage(y, min, max int, sensors map[coord]sensor) []interval {
	var intervals []interval
	for _, s := range sensors {
		coverageDist := s.dist(s.beacon)
		remaining := coverageDist - s.dist(coord{s.x, y})
		if remaining < 0 {
			continue
		}

		// Interval ends are not inclusive.
		interval := interval{s.x - remaining, s.x + remaining + 1}

		// Cull interval to min/max.
		if interval.start < min {
			interval.start = min
		}
		if interval.end > max {
			interval.end = max
		}
		intervals = append(intervals, interval)
	}

	return merge(intervals)
}

// merge merges any overlapping intervals in the specified intervals.
func merge(intervals []interval) []interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	var merged []interval
	var cur interval
	for i, next := range intervals {
		if i == 0 {
			cur = next
			continue
		}

		if next.start > cur.end {
			merged = append(merged, cur)
			cur = next
			continue
		}

		end := next.end
		if cur.end > end {
			end = cur.end
		}
		cur = interval{cur.start, end}
	}
	merged = append(merged, cur)

	return merged
}
