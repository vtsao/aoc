// This program implements the solution for
// https://adventofcode.com/2020/day/13.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// earliestBusAfterTS returns the earliest bus * the time you need to wait for
// that bus, that you can take after the earliest timestamp you can depart at.
func earliestBusAfterTS(earliestTS int, busses map[int]int) int {
	earliest := math.MaxInt64
	var busIDTimesWait int
	for bus := range busses {
		wait := bus - earliestTS%bus
		if wait < earliest {
			earliest = wait
			busIDTimesWait = wait * bus
		}
	}
	return busIDTimesWait
}

// earliestTSForSubseqDeparts returns earliest timestamp such that the first bus
// departs at that time and each subsequent listed bus departs at that
// subsequent minute.
//
// This solution requires the use of the Chinese Remainder Theorem (CRT)
// (https://en.wikipedia.org/wiki/Chinese_remainder_theorem). I'm not sure it
// can be done without it. Honestly I did not take the time to understand it. I
// adapted an impl of the CRT from
// https://golangnews.org/2020/12/computing-the-chinese-remainder-theorem/.
func earliestTSForSubseqDeparts(bs map[int]int) *big.Int {
	busses := map[*big.Int]*big.Int{}
	for b, targetMod := range bs {
		busses[big.NewInt(int64(b))] = big.NewInt(int64(targetMod))
	}

	// Compute N, which is the product of all the busses.
	N := big.NewInt(1)
	for bus := range busses {
		N.Mul(N, bus)
	}

	// x is the accumulated answer.
	x := new(big.Int)

	for bus, targetMod := range busses {
		// Nk = N/bus.
		Nk := new(big.Int).Div(N, bus)

		// N'k (Nkp) is the multiplicative inverse of Nk modulo bus.
		Nkp := new(big.Int)
		if Nkp.ModInverse(Nk, bus) == nil {
			return big.NewInt(-1)
		}

		// x += targetMod*Nk*Nkp.
		x.Add(x, Nkp.Mul(targetMod, Nkp.Mul(Nkp, Nk)))
	}

	return x.Mod(x, N)
}

func main() {
	file, err := os.Open("day13_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// There are only two lines for this input, so no need to loop through the
	// input file.
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	earliestTS, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	scanner.Scan()
	busSchedule := strings.Split(scanner.Text(), ",")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Parse out only the active busses. Map each bus ID to the mod result we
	// want, which is the earliest timestamp for part 2 mod the bus ID.
	busses := map[int]int{}
	for i, b := range busSchedule {
		if b == "x" {
			continue
		}
		bus, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal(err)
		}
		// E.g., let T be the earliest timestamp we want to find for part 2. Then if
		// bus ID is 59 with an offset of 4 (from the first bus), then we want the
		// mod result of T % 59 = (59 - 4).
		busses[bus] = (bus - i)
	}

	fmt.Printf("Part 1: earliest bus you can take * minutes you need to wait for that bus: %d\n", earliestBusAfterTS(earliestTS, busses))
	fmt.Printf("Part 2: earliest timestamp such that the first bus ID departs at that time and each subsequent listed bus ID departs at that subsequent minute: %d\n", earliestTSForSubseqDeparts(busses))
}
