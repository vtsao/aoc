// This program implements the solution for https://adventofcode.com/2021/day/3.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day03_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var bits [][]rune
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()
		bits = append(bits, []rune{})
		for _, bit := range row {
			bits[i] = append(bits[i], bit)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	gamma := ""
	epsilon := ""
	for j := range bits[0] {
		zeros, ones := 0, 0
		for _, row := range bits {
			if row[j] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			gamma += "0"
			epsilon += "1"
			continue
		}
		gamma += "1"
		epsilon += "0"
	}
	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilonDec, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1: %d\n", gammaDec*epsilonDec)
}
