// This program implements the solution for
// https://adventofcode.com/2022/day/17.
package main

import (
	"bufio"
	"os"
)

func main() {
	file, _ := os.Open("day17_input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}
}