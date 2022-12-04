// This program implements the solution for https://adventofcode.com/2022/day/4.
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
	file, err := os.Open("day04_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	overlapsCompletelyCount := 0
	overlapsCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		elf1Sections := parsePair(pair[0])
		elf2Sections := parsePair(pair[1])
		if overlapsCompletely(elf1Sections, elf2Sections) || overlapsCompletely(elf2Sections, elf1Sections) {
			overlapsCompletelyCount++
		}
		if overlaps(elf1Sections, elf2Sections) || overlaps(elf2Sections, elf1Sections) {
			overlapsCount++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", overlapsCompletelyCount)
	fmt.Printf("Part 2: %d\n", overlapsCount)
}

type section struct {
	start, end int
}

func parsePair(pair string) section {
	sections := strings.Split(pair, "-")
	start, _ := strconv.Atoi(sections[0])
	end, _ := strconv.Atoi(sections[1])
	return section{start: start, end: end}
}

// overlapsCompletely returns whether the first section is completely contained
// within the second section.
func overlapsCompletely(elf1Sections, elf2Sections section) bool {
	return elf1Sections.start >= elf2Sections.start && elf1Sections.end <= elf2Sections.end
}

// overlaps returns whether the first section overlaps with the second section.
func overlaps(elf1Sections, elf2Sections section) bool {
	return !(elf1Sections.end < elf2Sections.start || elf1Sections.start > elf2Sections.end)
}
