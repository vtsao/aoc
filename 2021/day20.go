// This program implements the solution for
// https://adventofcode.com/2021/day/20.
//
// curl -b "$(cat .session)" -o day20_input.txt https://adventofcode.com/2021/day/20/input
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func expand(image [][]string, oob string) [][]string {
	var expanded [][]string
	for i := 0; i < len(image)+2; i++ {
		var row []string
		for j := 0; j < len(image[0])+2; j++ {
			row = append(row, oob)
		}
		expanded = append(expanded, row)
	}

	for i := 1; i < len(expanded)-1; i++ {
		for j := 1; j < len(expanded[0])-1; j++ {
			expanded[i][j] = image[i-1][j-1]
		}
	}

	return expanded
}

func pixelAt(i, j int, image [][]string, oob string) string {
	if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) {
		return oob
	}
	return image[i][j]
}

func decode(idx, algo string) string {
	binary := ""
	for _, r := range idx {
		if r == '.' {
			binary += "0"
			continue
		}
		binary += "1"
	}

	dec, _ := strconv.ParseInt(binary, 2, 64)
	return string(algo[dec])
}

func enhance(image [][]string, algo string, evenStep bool) ([][]string, int) {
	oob := string(algo[0])
	if evenStep {
		oob = "."
	}
	image = expand(image, oob)

	enhanced := make([][]string, len(image))
	for i := range enhanced {
		enhanced[i] = make([]string, len(image[0]))
	}

	lit := 0
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			idx :=
				pixelAt(i-1, j-1, image, oob) +
					pixelAt(i-1, j, image, oob) +
					pixelAt(i-1, j+1, image, oob) +
					pixelAt(i, j-1, image, oob) +
					pixelAt(i, j, image, oob) +
					pixelAt(i, j+1, image, oob) +
					pixelAt(i+1, j-1, image, oob) +
					pixelAt(i+1, j, image, oob) +
					pixelAt(i+1, j+1, image, oob)
			pixel := decode(idx, algo)
			if pixel == "#" {
				lit++
			}
			enhanced[i][j] = pixel
		}
	}

	return enhanced, lit
}

func parseInput() (string, [][]string) {
	file, _ := os.Open("day20_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	algo := scanner.Text()
	scanner.Scan()

	var image [][]string
	for scanner.Scan() {
		var row []string
		for _, pixel := range strings.Split(scanner.Text(), "") {
			row = append(row, pixel)
		}
		image = append(image, row)
	}

	return algo, image
}

func main() {
	algo, image := parseInput()

	lit := 0
	for i := 0; i < 50; i++ {
		image, lit = enhance(image, algo, i%2 == 0)
		if i == 1 {
			fmt.Printf("Part 1: %d\n", lit)
		}
	}

	fmt.Printf("Part 2: %d\n", lit)
}
