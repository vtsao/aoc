// This program implements the solution for
// https://adventofcode.com/2021/day/15.
//
// curl -b "$(cat .session)" -o day15_input.txt https://adventofcode.com/2021/day/15/input
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type coord struct {
	i, j int
}

type node struct {
	coord coord
	dist  int
}

func min(remainingNodes map[*node]struct{}) *node {
	minNode := &node{dist: math.MaxInt64}
	for node := range remainingNodes {
		if node.dist < minNode.dist {
			minNode = node
		}
	}

	delete(remainingNodes, minNode)
	return minNode
}

func dijkstra(grid [][]int, remainingNodes map[*node]struct{}, coordToNode map[coord]*node, done map[coord]struct{}) {
	for len(remainingNodes) > 0 {
		curNode := min(remainingNodes)
		curCoord, curDist := curNode.coord, curNode.dist

		var adjCoords []coord
		if i, j := curCoord.i-1, curCoord.j; i >= 0 {
			adjCoords = append(adjCoords, coord{i: i, j: j})
		}
		if i, j := curCoord.i+1, curCoord.j; i < len(grid) {
			adjCoords = append(adjCoords, coord{i: i, j: j})
		}
		if i, j := curCoord.i, curCoord.j-1; j >= 0 {
			adjCoords = append(adjCoords, coord{i: i, j: j})
		}
		if i, j := curCoord.i, curCoord.j+1; j < len(grid[0]) {
			adjCoords = append(adjCoords, coord{i: i, j: j})
		}
		for _, adjCoord := range adjCoords {
			if _, ok := done[adjCoord]; !ok {
				if curDist+grid[adjCoord.i][adjCoord.j] < coordToNode[adjCoord].dist {
					coordToNode[adjCoord].dist = curDist + grid[adjCoord.i][adjCoord.j]
				}
			}
		}

		if curCoord.i == len(grid)-1 && curCoord.j == len(grid[0])-1 {
			return
		}
		done[curCoord] = struct{}{}
	}
}

func shortestPath(grid [][]int) int {
	remainingNodes := map[*node]struct{}{}
	coordToNode := map[coord]*node{}
	for i := range grid {
		for j := range grid[0] {
			n := &node{
				coord: coord{i: i, j: j},
				dist:  math.MaxInt64,
			}
			remainingNodes[n] = struct{}{}
			coordToNode[coord{i: i, j: j}] = n
		}
	}
	coordToNode[coord{i: 0, j: 0}].dist = 0
	dijkstra(grid, remainingNodes, coordToNode, map[coord]struct{}{})

	return coordToNode[coord{i: len(grid) - 1, j: len(grid[0]) - 1}].dist
}

func parseInput() ([][]int, error) {
	file, err := os.Open("day15_input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		for _, colStr := range scanner.Text() {
			col, err := strconv.Atoi(string(colStr))
			if err != nil {
				return nil, err
			}
			row = append(row, col)
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func main() {
	grid1, err := parseInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", shortestPath(grid1))

	var grid2 [][]int
	for row := 0; row < 5; row++ {
		for i := 0; i < len(grid1); i++ {
			var newRow []int
			for col := 0; col < 5; col++ {
				for j := 0; j < len(grid1[0]); j++ {
					val := grid1[i][j] + row + col
					if val > 9 {
						val -= 9
					}
					newRow = append(newRow, val)
				}
			}
			grid2 = append(grid2, newRow)
		}
	}
	fmt.Printf("Part 2: %d\n", shortestPath(grid2))
}
