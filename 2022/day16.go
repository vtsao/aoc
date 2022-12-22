// This program implements the solution for
// https://adventofcode.com/2022/day/16.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// parse parses a line describing a valve like
// "Valve SY has flow rate=0; tunnels lead to valves GW, LW" into a valve struct
// and stores it in the specified map.
func parse(line string, valves map[string]valve) {
	lineParts := strings.Split(line, "; ")
	valveParts := strings.Split(lineParts[0], " ")
	label := valveParts[1]
	flowRate, _ := strconv.Atoi(strings.TrimPrefix(valveParts[4], "rate="))
	tunnels := strings.Split(strings.Join(strings.Split(lineParts[1], " ")[4:], ""), ",")
	valves[label] = valve{
		label:    label,
		flowRate: flowRate,
		tunnels:  tunnels,
	}
}

func main() {
	file, _ := os.Open("day16_input.txt")

	valves := map[string]valve{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parse(scanner.Text(), valves)
	}

	// We need to turn this problem into a graph problem where we can DFS all the
	// paths for the best pressure.
	allValveDistances := map[string][]valveDistance{}
	for _, v := range valves {
		allValveDistances[v.label] = bfs(valves, v)
	}

	maxPressure := dfs(nil, "AA", 30, 0, map[string]any{}, valves, allValveDistances)
	fmt.Printf("Part 1: %d\n", maxPressure)

	// This part takes a while, it's not optimal.
	maxPressureWithElephant := dfs(map[string]int{}, "AA", 26, 0, map[string]any{}, valves, allValveDistances)
	fmt.Printf("Part 2: %d\n", maxPressureWithElephant)
}

type valve struct {
	label    string
	flowRate int
	tunnels  []string
}

func dfs(elephant map[string]int, cur string, timeLeft, pressure int, visited map[string]any, valves map[string]valve, allValveDistances map[string][]valveDistance) int {
	if timeLeft <= 0 {
		return pressure
	}

	bestPressure := pressure
	valveDistances := allValveDistances[cur]
	// DFS all the neighbors of the current valve, including itself (because you
	// start from "AA", but you don't necessarily open it). Moving to a neighbor
	// means opening it.
	for _, valveDistance := range valveDistances {
		if valves[valveDistance.valveLabel].flowRate == 0 {
			continue
		}
		if _, ok := visited[valveDistance.valveLabel]; ok {
			continue
		}

		// Calculate the final pressure you'll get after the remaining time if you
		// were to open this valve and add it to the total pressure so far. Then
		// DFS. This is basically if you were to open this valve and get the optimal
		// pressure for the remaining valves on your own.
		visited[valveDistance.valveLabel] = nil
		nextP := (timeLeft - valveDistance.distance - 1) * valves[valveDistance.valveLabel].flowRate
		p := dfs(elephant, valveDistance.valveLabel, timeLeft-valveDistance.distance-1, pressure+nextP, visited, valves, allValveDistances)
		delete(visited, valveDistance.valveLabel)

		// If you have an elephant, then get the optimal pressure if the elephant
		// opens this valve and the remaining valves. By doing this you get all
		// combinations of you opening a set of valves and the elephant opening the
		// rest. By "opening" I mean finding the optimal pressure from that set.
		eleP := 0
		if elephant != nil {
			// We need to cache the elephant's solutions, since we will likely have to
			// calculate this many times, since there are many ways to get here.
			// However, this still takes a while, maybe there is more I can cache?
			if cached, ok := elephant[key(visited)]; ok {
				eleP = cached
			} else {
				eleP = dfs(nil, "AA", 26, 0, visited, valves, allValveDistances)
				elephant[key(visited)] = eleP
			}
		}

		// If you stopped trying to open valves at this point and let the elephant
		// do the rest, would your combined pressures be better than if you just
		// tried to do the rest yourself? If so, then the combined pressure is the
		// candidate pressure.
		if eleP+pressure > p {
			p = eleP + pressure
		}
		if p > bestPressure {
			bestPressure = p
		}
	}

	return bestPressure
}

// key turns a visited hash set into a sorted string of valve labels to use as
// a cache key, probably not the most optimal way to do this.
func key(visited map[string]any) string {
	var k []string
	for valveLabel := range visited {
		k = append(k, valveLabel)
	}
	sort.Strings(k)
	return strings.Join(k, "")
}

// valveDistance represents a distance to a valve.
type valveDistance struct {
	valveLabel string
	distance   int
}

// bfs finds the shortest distance from the specified valve to all other valves
// including itself.
func bfs(valves map[string]valve, start valve) []valveDistance {
	var valveDistances []valveDistance

	visited := map[string]any{}
	q := []valve{start}
	dist := 0
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[i]
			if _, ok := visited[cur.label]; ok {
				continue
			}
			visited[cur.label] = nil

			valveDistances = append(valveDistances, valveDistance{cur.label, dist})
			for _, t := range cur.tunnels {
				q = append(q, valves[t])
			}
		}
		q = q[l:]
		dist++
	}

	return valveDistances
}
