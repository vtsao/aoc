// This program implements the solution for https://adventofcode.com/2021/day/3.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func oxygen(bits [][]rune, col int) [][]rune {
	if len(bits) == 1 {
		return bits
	}

	var zeroRows, oneRows [][]rune
	zeroes, ones := 0, 0

	for _, row := range bits {
		if row[col] == '0' {
			zeroes++
			zeroRows = append(zeroRows, row)
			continue
		}
		ones++
		oneRows = append(oneRows, row)
	}

	if zeroes > ones {
		return oxygen(zeroRows, col+1)
	}
	return oxygen(oneRows, col+1)
}

func co2(bits [][]rune, col int) [][]rune {
	if len(bits) == 1 {
		return bits
	}

	var zeroRows, oneRows [][]rune
	zeroes, ones := 0, 0

	for _, row := range bits {
		if row[col] == '0' {
			zeroes++
			zeroRows = append(zeroRows, row)
			continue
		}
		ones++
		oneRows = append(oneRows, row)
	}

	if zeroes <= ones {
		return co2(zeroRows, col+1)
	}
	return co2(oneRows, col+1)
}

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
				continue
			}
			ones++
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

	oxygenDec, err := strconv.ParseInt(string(oxygen(bits, 0)[0]), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2Dec, err := strconv.ParseInt(string(co2(bits, 0)[0]), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 2: %d\n", oxygenDec*co2Dec)
}
