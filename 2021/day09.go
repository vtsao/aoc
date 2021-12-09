// This program implements the solution for https://adventofcode.com/2021/day/9.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coord struct {
	i, j int
}

func dfs(i, j int, heights [][]int, visited map[coord]interface{}) int {
	if _, ok := visited[coord{i: i, j: j}]; ok {
		return 0
	}
	if i < 0 || i >= len(heights) || j < 0 || j >= len(heights[0]) || heights[i][j] == 9 {
		return 0
	}

	visited[coord{i: i, j: j}] = nil
	size := 1
	size += dfs(i-1, j, heights, visited)
	size += dfs(i+1, j, heights, visited)
	size += dfs(i, j-1, heights, visited)
	size += dfs(i, j+1, heights, visited)

	return size
}

func main() {
	file, err := os.Open("day09_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var heights [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		for _, heightStr := range strings.Split(scanner.Text(), "") {
			h, err := strconv.Atoi(heightStr)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, h)
		}
		heights = append(heights, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sumRiskLevels := 0
	var basinSizes []int
	visited := map[coord]interface{}{}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			// Part 1.
			if i-1 >= 0 && heights[i-1][j] <= heights[i][j] {
				continue
			}
			if i+1 < len(heights) && heights[i+1][j] <= heights[i][j] {
				continue
			}
			if j-1 >= 0 && heights[i][j-1] <= heights[i][j] {
				continue
			}
			if j+1 < len(heights[0]) && heights[i][j+1] <= heights[i][j] {
				continue
			}
			sumRiskLevels += heights[i][j] + 1

			// Part 2.
			basinSize := dfs(i, j, heights, visited)
			if basinSize > 0 {
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}
	sort.Ints(basinSizes)
	fmt.Printf("Part 1: %d\n", sumRiskLevels)
	fmt.Printf("Part 2: %d\n", basinSizes[len(basinSizes)-1]*basinSizes[len(basinSizes)-2]*basinSizes[len(basinSizes)-3])
}
