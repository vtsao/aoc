// This program implements the solution for
// https://adventofcode.com/2022/day/16.
package main

import (
	"bufio"
	"fmt"
	"os"
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

	valveShortestDists := map[string]map[string]int{}
	for _, v := range valves {
		dists := map[string]int{}
		bfs(valves, v, dists)
		valveShortestDists[v.label] = dists
	}

	maxPressure := dfs("AA", 30, map[string]any{}, 0, valves, valveShortestDists)
	fmt.Printf("Part 1: %d\n", maxPressure)

	maxPressure2 := dfs2("AA", "AA", 26, 26, map[string]any{}, 0, valves, valveShortestDists)
	fmt.Printf("Part 2: %d\n", maxPressure2)
}

type valve struct {
	label    string
	flowRate int
	tunnels  []string
}

func (v *valve) String() string {
	return fmt.Sprintf("{%s, %d, %v}\n", v.label, v.flowRate, v.tunnels)
}

func dfs(cur string, timeLeft int, visited map[string]any, pressure int, valves map[string]valve, valveShortestDists map[string]map[string]int) int {
	if timeLeft <= 0 {
		return pressure
	}

	bestPressure := pressure
	shorestDists := valveShortestDists[cur]
	for v, d := range shorestDists {
		if valves[v].flowRate == 0 {
			continue
		}
		if _, ok := visited[v]; ok {
			continue
		}

		visited[v] = nil
		nextP := (timeLeft - d - 1) * valves[v].flowRate
		p := dfs(v, timeLeft-d-1, visited, pressure+nextP, valves, valveShortestDists)
		delete(visited, v)
		if p > bestPressure {
			bestPressure = p
		}
	}

	return bestPressure
}

func dfs2(cur, eleCur string, timeLeft, eleTimeLeft int, visited map[string]any, pressure int, valves map[string]valve, valveShortestDists map[string]map[string]int) int {
	if timeLeft <= 0 && eleTimeLeft <= 0 {
		return pressure
	}

	bestPressure := pressure
	shorestDists := valveShortestDists[cur]
	eleShorestDists := valveShortestDists[eleCur]
	for v, d := range shorestDists {
		if valves[v].flowRate == 0 {
			continue
		}
		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = nil

		for eleV, eleD := range eleShorestDists {
			if v == eleV {
				continue
			}
			if valves[eleV].flowRate == 0 {
				continue
			}
			if _, ok := visited[eleV]; ok {
				continue
			}

			visited[eleV] = nil
			nextP := (timeLeft - d - 1) * valves[v].flowRate
			if timeLeft <= 0 {
				nextP = 0
			}
			eleNextP := (eleTimeLeft - eleD - 1) * valves[eleV].flowRate
			if eleTimeLeft <= 0 {
				eleNextP = 0
			}
			p := dfs2(v, eleV, timeLeft-d-1, eleTimeLeft-eleD-1, visited, pressure+nextP+eleNextP, valves, valveShortestDists)

			delete(visited, eleV)
			if p > bestPressure {
				bestPressure = p
			}
		}

		delete(visited, v)
	}

	return bestPressure
}

// bfs to get shortest distance from each valve to every other valve.
// Then starting at AA, iterate through all valves AA can get to and find the
// max pressure you can release based on the time left.
// Then set that node to the cur node and continue until either you've opened
// all the valves (keep track of opened valves so you skip them), or you run
// out of time.
func bfs(valves map[string]valve, start valve, dists map[string]int) {
	q := []valve{start}

	dist := 0
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			cur := q[i]
			if _, ok := dists[cur.label]; ok {
				continue
			}

			dists[cur.label] = dist
			for _, t := range cur.tunnels {
				q = append(q, valves[t])
			}
		}
		q = q[l:]
		dist++
	}
}
