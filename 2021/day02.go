// This program implements the solution for https://adventofcode.com/2021/day/2.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cmd struct {
	dir string
	val int
}

func main() {
	file, err := os.Open("day02_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cmds []cmd
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmdStr := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(cmdStr[1])
		if err != nil {
			log.Fatal(err)
		}

		cmds = append(cmds, cmd{dir: cmdStr[0], val: val})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	hPos1, depth1 := 0, 0
	for _, cmd := range cmds {
		switch cmd.dir {
		case "forward":
			hPos1 += cmd.val
		case "up":
			depth1 -= cmd.val
		case "down":
			depth1 += cmd.val
		}
	}
	fmt.Printf("Part 1: %d\n", hPos1*depth1)

	hPos2, depth2, aim := 0, 0, 0
	for _, cmd := range cmds {
		switch cmd.dir {
		case "forward":
			hPos2 += cmd.val
			depth2 += aim * cmd.val
		case "up":
			aim -= cmd.val
		case "down":
			aim += cmd.val
		}
	}
	fmt.Printf("Part 2: %d\n", hPos2*depth2)
}
