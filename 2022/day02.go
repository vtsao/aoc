// This program implements the solution for https://adventofcode.com/2022/day/2.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day02_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalScore1 := 0
	totalScore2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		totalScore1 += score1(round[0], round[1])
		totalScore2 += score2(round[0], round[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", totalScore1)
	fmt.Printf("Part 2: %d\n", totalScore2)
}

func score1(opponent, you string) int {
	switch opponent {
	case "A":
		switch you {
		case "X":
			return 1 + 3
		case "Y":
			return 2 + 6
		case "Z":
			return 3 + 0
		}
	case "B":
		switch you {
		case "X":
			return 1 + 0
		case "Y":
			return 2 + 3
		case "Z":
			return 3 + 6
		}
	case "C":
		switch you {
		case "X":
			return 1 + 6
		case "Y":
			return 2 + 0
		case "Z":
			return 3 + 3
		}
	}

	// Will never get here for valid input.
	return 0
}

func score2(opponent, you string) int {
	switch opponent {
	case "A":
		switch you {
		case "X":
			return 3 + 0
		case "Y":
			return 1 + 3
		case "Z":
			return 2 + 6
		}
	case "B":
		switch you {
		case "X":
			return 1 + 0
		case "Y":
			return 2 + 3
		case "Z":
			return 3 + 6
		}
	case "C":
		switch you {
		case "X":
			return 2 + 0
		case "Y":
			return 3 + 3
		case "Z":
			return 1 + 6
		}
	}

	// Will never get here for valid input.
	return 0
}
