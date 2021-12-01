// This program implements the solution for
// https://adventofcode.com/2020/day/11.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// seatOccupied returns whether the seat is occupied.
func seatOccupied(i, j int, seats []string) bool {
	if i < 0 || i >= len(seats) || j < 0 || j >= len(seats[0]) {
		return false
	}
	return string(seats[i][j]) == "#"
}

// firstVisibleSeatOccupied returns whether the first visible seat in a specific
// direction is occupied.
func firstVisibleSeatOccupied(i, j, iDir, jDir int, seats []string) bool {
	if i < 0 || i >= len(seats) || j < 0 || j >= len(seats[0]) {
		return false
	}
	if string(seats[i][j]) == "#" {
		return true
	}
	if string(seats[i][j]) == "L" {
		return false
	}
	return firstVisibleSeatOccupied(i+iDir, j+jDir, iDir, jDir, seats)
}

func countOccupied(seats []string) int {
	count := 0
	for _, row := range seats {
		for _, seat := range row {
			if string(seat) == "#" {
				count++
			}
		}
	}

	return count
}

// Run changes the seat states according to the rules until there are two
// consecutive state changes with no diff.
func run(seats []string, nearsighted bool) []string {
	for {
		stateChanged := false
		nextState := make([]string, len(seats))
		for i, row := range seats {
			for j, seat := range row {
				var nextSeatState string
			outer:
				switch string(seat) {

				case "L":
					for _, adjI := range []int{-1, 0, 1} {
						for _, adjJ := range []int{-1, 0, 1} {
							if adjI == 0 && adjJ == 0 {
								continue
							}

							// If nearsighted, empty seats only get taken if all their
							// adjacent seats are empty.
							if nearsighted {
								if seatOccupied(i+adjI, j+adjJ, seats) {
									break outer
								}
								continue
							}

							// Otherwise empty seats only get taken if one cannot see another
							// first visible occupied seat in the 8 directions around this
							// seat.
							if firstVisibleSeatOccupied(i+adjI, j+adjJ, adjI, adjJ, seats) {
								break outer
							}
						}
					}
					nextSeatState = "#"

				case "#":
					count := 0
					for _, adjI := range []int{-1, 0, 1} {
						for _, adjJ := range []int{-1, 0, 1} {
							if adjI == 0 && adjJ == 0 {
								continue
							}

							// If nearsighted, occupied seats only become empty if at least 4
							// of their adjacent seats are occupied.
							if nearsighted {
								if seatOccupied(i+adjI, j+adjJ, seats) {
									count++
									if count >= 4 {
										nextSeatState = "L"
										break outer
									}
								}
								continue
							}

							// Otherwise occupied seats only become empty if one can see 5 or
							// more first visible occupied seats in the 8 directions around
							// this seat.
							if firstVisibleSeatOccupied(i+adjI, j+adjJ, adjI, adjJ, seats) {
								count++
								if count >= 5 {
									nextSeatState = "L"
									break outer
								}
							}
						}
					}
				}

				// If the state of this seat didn't change, make sure it gets its
				// previous state.
				if nextSeatState == "" {
					nextState[i] += string(seat)
					continue
				}
				nextState[i] += nextSeatState
				stateChanged = true
			}
		}

		// Finish when there's no diff, otherwise we update the state to the new
		// state we just calculated.
		if !stateChanged {
			break
		}
		seats = nextState
	}

	return seats
}

func main() {
	file, err := os.Open("day11_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Load in all the seats into memory.
	var seats []string
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		seats = append(seats, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d occupied seats\n", countOccupied(run(seats, true)))
	fmt.Printf("Part 2: %d occupied seats\n", countOccupied(run(seats, false)))
}
