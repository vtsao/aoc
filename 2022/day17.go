// This program implements the solution for
// https://adventofcode.com/2022/day/17.
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed day17_input.txt
var jetPattern string

func main() {
	numRocks := 2022
	floor, _, _ := fall(&numRocks, jetPattern, 0)
	fmt.Printf("Part 1: %d\n", floor)

	floor2 := fallLarge(1000000000000, jetPattern, 50)
	fmt.Printf("Part 2: %d\n", floor2)
}

// fallLarge calculates the height of the rock structure for large numbers of
// rocks. It does so by finding a repeat in the rock structure pattern. Given
// the way the rocks fall follows a specific pattern, it must repeat at some
// point. Once we find a repeat, we can calculate how many times the repeated
// pattern appears for 1000000000000 rocks.
//
// The calculation is a bit tedious. Here's a diagram explanation:
//
// 1. My algorithm finds a repeated pattern from 687 rocks to 2387 rocks.
// Rocks:  0 ... (687  ... 2387) ... (repeated) ... 1000000000000
// Height: 0 ... (1053 ... 3707) ... (repeated) ... 1561176470569
//
// 2. Calculate how many times the repeated rocks can go into
// 1000000000000 - 687:
//
// (1000000000000 - 687) / 1700 == 588235293
// 1700 is the number of rocks in the repeated pattern.
// You need to subtract 687 because the pattern does not repeat from 0 rocks.
//
// 3. Use this to find the height of the repeated rocks:
//
// 588235293 * 2654 == 1561176467622
// 2654 is the height of the repeated pattern.
//
// 4. Now figure out the height of the remaining rocks for the pattern because
// there is a remainder in the number of times the repeated rock section can go
// into 1000000000000 - 687.
//
// 1000000000000 - 999999998100 == 1900
// 999999998100 is from 588235293 * 1700. Technically we are finding only the
// remainder height of the pattern portion, but to do this we need to use the
// height of where the repeated pattern started, since we need to calculate
// this from 0 rocks. So technically we should subtract the 687 from 1900. But
// we don't, since we would need to add that back for the final height
// calculation anyways.
//
// So reuse the small fall algorithm to calculate the height after 1900 rocks,
// which is 2947.
//
// 5. Finally, add everything up:
//
// 1561176467622 + 2947 == 1561176470569
func fallLarge(numRocks int, jetPattern string, threshold int) int {
	_, start, end := fall(nil, jetPattern, threshold)

	numRocksPerRepeat := end.rocks - start.rocks
	heightPerRepeat := end.height - start.height

	timesRepeated := (numRocks - start.rocks) / numRocksPerRepeat
	heightOfRepeated := timesRepeated * heightPerRepeat

	rocksRepeated := numRocksPerRepeat * timesRepeated
	remainingRocks := 1000000000000 - rocksRepeated
	heightOfRemainingRocks, _, _ := fall(&remainingRocks, jetPattern, threshold)

	return heightOfRepeated + heightOfRemainingRocks
}

type coord struct {
	x, y int
}

type rocksAndHeight struct {
	rocks, height int
}

// fall simulates the rocks falling in the cave. If a number of rocks is
// specified, it simulates the rock fall pattern for that number of rocks. If
// a number of rocks is not specified, it finds a repeat in the pattern.
//
// The threshold is used when finding a repeated rock pattern. We only keep
// track of the last threshold rock coordinates, culling everything else. This
// is needed because I don't know how to calculate what threshold will work. So
// it was a bit of trial and error to find a threshold that works for my input.
func fall(numRocks *int, jetPattern string, threshold int) (int, rocksAndHeight, rocksAndHeight) {
	floor := 0
	shapeType := 0
	rocks := map[coord]any{}
	jetIdx := 0
	hashes := map[string]rocksAndHeight{}
	for i := 0; ; i++ {
		if numRocks != nil && i == *numRocks {
			break
		}

		coords := shape(floor, shapeType)
		top, jIdx := move(floor, coords, jetIdx, jetPattern, rocks)
		if top > floor {
			floor = top
		}

		if numRocks == nil {
			key := cullAndKey(floor, shapeType, jetIdx, threshold, rocks)
			if floor >= threshold {
				if startRocksAndHeight, ok := hashes[key]; ok {
					return 0, startRocksAndHeight, rocksAndHeight{i, floor}
				}
				hashes[key] = rocksAndHeight{i, floor}
			}
		}

		jetIdx = jIdx
		shapeType++
		if shapeType == 5 {
			shapeType = 0
		}
	}

	return floor, rocksAndHeight{}, rocksAndHeight{}
}

// cullAndKey removes any rock coordinates below the specified threshold from
// the current floor. E.g., if the floor is 100 and the threshold is 10, then
// any rocks under a y coord of 90 will be removed. Essentially we only keep
// track of the last threshold rows. It also calculates a hash for the last
// threshold rows that we use to find when the pattern repeats.
//
// The hash is not optimal, but fast enough for us and works.
func cullAndKey(floor, shapeType, jetIdx, threshold int, rocks map[coord]any) string {
	keyArr := []string{strconv.Itoa(shapeType), strconv.Itoa(jetIdx)}
	for rock := range rocks {
		if rock.y > floor-threshold {
			keyArr = append(keyArr, fmt.Sprintf("%d,%d", rock.x, rock.y-(floor-8)))
		}
		if rock.y <= floor-threshold {
			delete(rocks, rock)
		}
	}

	sort.Strings(keyArr)
	return strings.Join(keyArr, ";")
}

// shape returns the coordinates where the next shape will appear depending on
// the shape type and current floor.
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

// move simulates the movement of the specified shape using the specified jet
// pattern and current jet index until the shape comes to a stop.
func move(floor int, shape []coord, jetIdx int, jetPattern string, rocks map[coord]any) (int, int) {
	cur := shape
	jet := true
	for {
		var next []coord
		for _, c := range cur {
			if jet {
				if jetPattern[jetIdx] == '>' {
					nextC := coord{c.x + 1, c.y}
					// If we can't move this shape left or right, stop here, but say we've
					// "moved" the entire shape so we can still try to move it down.
					if _, ok := rocks[nextC]; ok || nextC.x >= 7 {
						next = cur
						break
					}
					next = append(next, nextC)
					continue
				}

				nextC := coord{c.x - 1, c.y}
				if _, ok := rocks[nextC]; ok || nextC.x < 0 {
					// If we can't move this shape left or right, stop here, but say we've
					// "moved" the entire shape so we can still try to move it down.
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

		// If we weren't able to move the entire shape down b/c it is blocked, then
		// the shape has reached its final resting place.
		if len(next) != len(cur) {
			break
		}
		cur = next
	}

	// Calculate the top y coord of the shape.
	top := 0
	for _, c := range cur {
		rocks[c] = nil
		if c.y > top {
			top = c.y
		}
	}
	return top, jetIdx
}
