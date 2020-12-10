package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func bs(start, end int, ranges string) int {
	if len(ranges) == 1 {
		if r := string(ranges[0]); r == "F" || r == "L" {
			return start
		}
		return end - 1
	}

	mid := start + (end-start)/2
	if r := string(ranges[0]); r == "F" || r == "L" {
		return bs(start, mid, ranges[1:])
	}
	return bs(mid, end, ranges[1:])
}

func main() {
	file, err := os.Open("day5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	maxSeatID := 0
	seats := map[int]interface{}{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seat := scanner.Text()
		row := bs(0, 128, seat[:7])
		col := bs(0, 8, seat[7:])
		seatID := row*8 + col
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
		seats[seatID] = nil
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d is the max seat ID\n", maxSeatID)

	for row := 0; row < 128; row++ {
		for col := 0; col < 8; col++ {
			seatID := row*8 + col
			if _, ok := seats[seatID]; !ok {
				_, ok1 := seats[seatID+1]
				_, ok2 := seats[seatID-1]
				if ok1 && ok2 {
					fmt.Printf("Part 2: %d is my seat ID\n", seatID)
				}
			}
		}
	}
}
