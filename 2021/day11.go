// This program implements the solution for
// https://adventofcode.com/2021/day/11.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	i, j int
}

func flash(i, j int, energyLevels [][]int, flashed map[cell]interface{}, causedByFlash bool) int {
	if _, ok := flashed[cell{i: i, j: j}]; ok {
		return 0
	}
	if i < 0 || i >= len(energyLevels) || j < 0 || j >= len(energyLevels[0]) {
		return 0
	}

	if causedByFlash {
		energyLevels[i][j]++
	}
	if energyLevels[i][j] <= 9 {
		return 0
	}

	flashed[cell{i: i, j: j}] = nil
	flashes := 1
	flashes += flash(i-1, j, energyLevels, flashed, true)
	flashes += flash(i+1, j, energyLevels, flashed, true)
	flashes += flash(i, j-1, energyLevels, flashed, true)
	flashes += flash(i, j+1, energyLevels, flashed, true)
	flashes += flash(i-1, j-1, energyLevels, flashed, true)
	flashes += flash(i-1, j+1, energyLevels, flashed, true)
	flashes += flash(i+1, j-1, energyLevels, flashed, true)
	flashes += flash(i+1, j+1, energyLevels, flashed, true)
	energyLevels[i][j] = 0

	return flashes
}

func main() {
	file, err := os.Open("day11_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var energyLevels [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		for _, energyStr := range strings.Split(scanner.Text(), "") {
			energy, err := strconv.Atoi(energyStr)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, energy)
		}
		energyLevels = append(energyLevels, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	flashes := 0
	for s := 0; ; s++ {
		for i := 0; i < len(energyLevels); i++ {
			for j := 0; j < len(energyLevels[0]); j++ {
				energyLevels[i][j]++
			}
		}
		flashed := map[cell]interface{}{}
		for i := 0; i < len(energyLevels); i++ {
			for j := 0; j < len(energyLevels[0]); j++ {
				flashes += flash(i, j, energyLevels, flashed, false)
			}
		}
		if len(flashed) == 100 {
			fmt.Printf("Part 2: %d\n", s+1)
			// This only works because we assume part 2's answer is > 100.
			return
		}
		if s == 99 {
			fmt.Printf("Part 1: %d\n", flashes)
		}
	}
}
