package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day6_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
