// This program implements the solution for https://adventofcode.com/2022/day/6.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day06_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := scanner.Text()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", marker(data, 4))
	fmt.Printf("Part 2: %d\n", marker(data, 14))
}

func marker(data string, distinctChars int) int {
data:
	for i := 0; i <= len(data)-distinctChars; i++ {
		four := data[i : i+distinctChars]

		charsCount := map[rune]int{}
		for _, char := range four {
			charsCount[char]++
		}

		for _, count := range charsCount {
			if count > 1 {
				continue data
			}
		}

		return i + distinctChars
	}

	// Will never happen for valid input.
	return 0
}
