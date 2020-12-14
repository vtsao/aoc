package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	entries := map[int]interface{}{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		entries[entry] = nil
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for entry := range entries {
		target := 2020 - entry
		if _, ok := entries[target]; ok {
			fmt.Printf("Part 1: entry: %d, target: %d, answer: %d\n", entry, target, entry*target)
			break
		}
	}

outer:
	for entry1 := range entries {
		target1 := 2020 - entry1
		for entry2 := range entries {
			target2 := target1 - entry2
			if _, ok := entries[target2]; ok {
				fmt.Printf("Part 2: entry1: %d, entry2: %d, target2: %d, answer: %d\n", entry1, entry2, target2, entry1*entry2*target2)
				break outer
			}
		}
	}
}
