// This program implements the solution for
// https://adventofcode.com/2021/day/24.
//
// curl -b "$(cat .session)" -o day24_input.txt https://adventofcode.com/2021/day/24/input
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func resolve(vars map[string]int, arg string) int {
	val, err := strconv.Atoi(arg)
	if err != nil {
		return vars[arg]
	}
	return val
}

func exec(instructions []string, input string) map[string]int {
	vars := map[string]int{}

	inputIdx := 0
	for _, instStr := range instructions {
		inst := strings.Split(instStr, " ")
		switch inst[0] {
		case "inp":
			vars[inst[1]], _ = strconv.Atoi(string(input[inputIdx]))
			inputIdx++
		case "add":
			vars[inst[1]] = vars[inst[1]] + resolve(vars, inst[2])
		case "mul":
			vars[inst[1]] = vars[inst[1]] * resolve(vars, inst[2])
		case "div":
			vars[inst[1]] = vars[inst[1]] / resolve(vars, inst[2])
		case "mod":
			vars[inst[1]] = vars[inst[1]] % resolve(vars, inst[2])
		case "eql":
			if vars[inst[1]] == resolve(vars, inst[2]) {
				vars[inst[1]] = 1
				continue
			}
			vars[inst[1]] = 0
		}
	}

	return vars
}

func parseInput() []string {
	file, _ := os.Open("day24_input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func main() {
	instructions := parseInput()
	vars := exec(instructions, "3")

	for v, val := range vars {
		log.Printf("%s: %d", v, val)
	}

	// fmt.Printf("Part 1: %d\n", )
	// fmt.Printf("Part 2: %d\n", )
}
