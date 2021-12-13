// This program implements the solution for
// https://adventofcode.com/2021/day/12.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type searcher struct {
	connections      map[string][]string
	pathsThroughCave [][]string
}

func (s *searcher) dfs(curCave string, pathSoFar []string, oneUp int) {
	if strings.ToLower(curCave) == curCave {
		for _, visistedCave := range pathSoFar {
			if curCave == visistedCave {
				switch oneUp {
				case -1:
					return
				case 0:
					oneUp = 1
				case 1:
					return
				}
			}
		}
	}

	pathSoFar = append(pathSoFar, curCave)
	if curCave == "end" {
		s.pathsThroughCave = append(s.pathsThroughCave, pathSoFar)
		return
	}

	for _, connection := range s.connections[curCave] {
		if connection != "start" {
			s.dfs(connection, pathSoFar, oneUp)
		}
	}
}

func main() {
	file, err := os.Open("day12_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	connections := map[string][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := strings.Split(scanner.Text(), "-")
		connections[path[0]] = append(connections[path[0]], path[1])
		connections[path[1]] = append(connections[path[1]], path[0])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s1 := &searcher{connections: connections}
	s1.dfs("start", []string{}, -1)
	fmt.Printf("Part 1: %d\n", len(s1.pathsThroughCave))

	s2 := &searcher{connections: connections}
	s2.dfs("start", []string{}, 0)
	fmt.Printf("Part 2: %d\n", len(s2.pathsThroughCave))
}
