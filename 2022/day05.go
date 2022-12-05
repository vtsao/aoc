// This program implements the solution for https://adventofcode.com/2022/day/5.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day05_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	crates1 := crates()
	crates2 := crates()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		if !strings.HasPrefix(instruction, "move ") {
			continue
		}

		// E.g., "move 6 from 9 to 3".
		instructionParts := strings.Split(instruction, " ")
		num, _ := strconv.Atoi(instructionParts[1])
		from, _ := strconv.Atoi(instructionParts[3])
		to, _ := strconv.Atoi(instructionParts[5])
		crates1 = move(crates1, num, from, to)
		crates2 = move2(crates2, num, from, to)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %q\n", topCrates(crates1))
	fmt.Printf("Part 2: %q\n", topCrates(crates2))
}

func topCrates(crates []string) string {
	top := ""
	for _, crate := range crates {
		top += string(crate[len(crate)-1])
	}
	return top
}

func move(crates []string, num, from, to int) []string {
	for i := 0; i < num; i++ {
		fromCrates := crates[from-1]
		fromTop := fromCrates[len(fromCrates)-1]

		crates[to-1] += string(fromTop)

		fromCrates = fromCrates[:len(fromCrates)-1]
		crates[from-1] = fromCrates
	}
	return crates
}

func move2(crates []string, num, from, to int) []string {
	fromCrates := crates[from-1]
	fromCratesToMove := fromCrates[len(fromCrates)-num:]

	crates[to-1] += fromCratesToMove

	fromCratesToKeep := fromCrates[:len(fromCrates)-num]
	crates[from-1] = fromCratesToKeep

	return crates
}

func crates() []string {
	return []string{
		"GFVHPS",
		"GJFBVDZM",
		"GMLJN",
		"NGZVDWP",
		"VRCB",
		"VRSMPWLZ",
		"THP",
		"QRSNCHZV",
		"FLGPVQJ",
	}
}
