// This program implements the solution for
// https://adventofcode.com/2020/day/12.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// moveWithProbableInsts calculates the Manhattan distance between the ship's
// original location and its final location using the navigation rules you
// guessed (part 1).
func moveWithProbableInsts(insts []*actVal) int {
	var dir, ns, ew int
	for _, inst := range insts {
		val := inst.val
		switch inst.act {
		case "N":
			ns += val
		case "S":
			ns -= val
		case "E":
			ew += val
		case "W":
			ew -= val
		case "R":
			val = 360 - val
			fallthrough
		case "L":
			dir = (dir + val) % 360
		case "F":
			switch dir {
			case 0:
				ew += val
			case 90:
				ns += val
			case 180:
				ew -= val
			case 270:
				ns -= val
			}
		}
	}

	return int(math.Abs(float64(ns)) + math.Abs(float64(ew)))
}

// moveWithActualInsts calculates the Manhattan distance between the ship's
// original location and its final location using the actual navigation rules
// (part 2).
func moveWithActualInsts(insts []*actVal) int {
	var ns, ew int
	wayX, wayY := 10, 1
	for _, inst := range insts {
		val := inst.val
		switch inst.act {
		case "N":
			wayY += val
		case "S":
			wayY -= val
		case "E":
			wayX += val
		case "W":
			wayX -= val
		case "R":
			val = 360 - val
			fallthrough
		case "L":
			switch val {
			case 90:
				wayX, wayY = -wayY, wayX
			case 180:
				wayX, wayY = -wayX, -wayY
			case 270:
				wayX, wayY = wayY, -wayX
			}
		case "F":
			ns += val * wayY
			ew += val * wayX
		}
	}

	return int(math.Abs(float64(ns)) + math.Abs(float64(ew)))
}

type actVal struct {
	act string
	val int
}

func main() {
	file, err := os.Open("day12_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var insts []*actVal
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		inst := scanner.Text()
		act := string(inst[0])
		val, err := strconv.Atoi(inst[1:len(inst)])
		if err != nil {
			log.Fatal(err)
		}
		insts = append(insts, &actVal{act: act, val: val})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// NOTE: that there are some important assumptions for this to work that the
	// instructions don't mention. If any of these are no longer true, this code
	// won't work.
	//
	// (1) all rotations are multiples of 90 degrees, this ensures only integer
	//     numbers are used in all calculations.
	// (2) the input rotation values are only 90, 180, and 270, meaning nothing >
	//     360 and no no-ops (i.e., no 0 rotations).

	fmt.Printf("Part 1: Manhattan distance between final location and starting location: %d\n", moveWithProbableInsts(insts))
	fmt.Printf("Part 2: Manhattan distance between final location and starting location: %d\n", moveWithActualInsts(insts))
}
