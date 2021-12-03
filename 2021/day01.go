// This program implements the solution for https://adventofcode.com/2021/day/1.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day01_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var depths []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, depth)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	increaseCnt1 := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increaseCnt1++
		}
	}
	fmt.Printf("Part 1: %d\n", increaseCnt1)

	increaseCnt2 := 0
	prevMeasurement := depths[0] + depths[1] + depths[2]
	for i := 1; i <= len(depths)-3; i++ {
		curMeasurement := depths[i] + depths[i+1] + depths[i+2]
		if curMeasurement > prevMeasurement {
			increaseCnt2++
		}
		prevMeasurement = curMeasurement
	}
	fmt.Printf("Part 2: %d\n", increaseCnt2)
}
