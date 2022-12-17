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
func parse(line string, valves map[string]*valve) {
	lineParts := strings.Split(line, "; ")
	valveParts := strings.Split(lineParts[0], " ")
	label := valveParts[1]
	flowRate, _ := strconv.Atoi(strings.TrimPrefix(valveParts[4], "rate="))
	tunnels := strings.Split(strings.Join(strings.Split(lineParts[1], " ")[4:], ""), ",")
	valves[label] = &valve{
		label:    label,
		flowRate: flowRate,
		tunnels:  tunnels,
	}
}

func main() {
	file, _ := os.Open("day16_input.txt")

	valves := map[string]*valve{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parse(scanner.Text(), valves)
	}

	// fmt.Printf("Part 1: %d\n", )
	// fmt.Printf("Part 2: %d\n", )
}

type valve struct {
	label    string
	flowRate int
	tunnels  []string
}

func (v *valve) String() string {
	return fmt.Sprintf("{%s, %d, %v}\n", v.label, v.flowRate, v.tunnels)
}
