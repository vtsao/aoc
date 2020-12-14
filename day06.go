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

	// For this to work we need to cheat a bit by adding an extra newline at the
	// end of the input so we process the last group.
	anyYesCount := 0
	allYesCount := 0
	curGroup := map[string]int{}
	curGroupLen := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		yesses := scanner.Text()

		if len(yesses) == 0 {
			anyYesCount += len(curGroup)

			for _, count := range curGroup {
				if count == curGroupLen {
					allYesCount++
				}
			}

			curGroup = map[string]int{}
			curGroupLen = 0
			continue
		}

		curGroupLen++
		for _, q := range yesses {
			curGroup[string(q)]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: any yes count: %d\n", anyYesCount)
	fmt.Printf("Part 2: all yes count: %d\n", allYesCount)
}
