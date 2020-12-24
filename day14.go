// This program implements the solution for
// https://adventofcode.com/2020/day/14.
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day14_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
