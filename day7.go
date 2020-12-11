package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for scanner.Scan() {
		scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
