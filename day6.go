package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day6_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// For this to work we need to cheat a bit by adding an extra newline at the
	// end of the input so we process the last group.
	yesCount := 0
	group := map[string]interface{}{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		yesses := scanner.Text()

		if len(yesses) == 0 {
			yesCount += len(group)
			group = map[string]interface{}{}
			continue
		}

		for _, q := range yesses {
			group[string(q)] = nil
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: yes count: %d\n", yesCount)
}
