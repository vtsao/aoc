package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	cmd string
	val int
}

func run(insts []instruction) (int, bool) {
	visited := map[int]interface{}{}
	finished := true
	acc := 0
	for i := 0; i < len(insts); {
		if _, ok := visited[i]; ok {
			finished = false
			break
		}

		visited[i] = nil
		inst := insts[i]
		switch inst.cmd {
		case "acc":
			acc += inst.val
			i++
		case "jmp":
			i += inst.val
		case "nop":
			i++
		}
	}

	return acc, finished
}

func runWithFix(insts []instruction) int {
	// This is horribly inefficient, but whatever, we'll brute force it.

	for i, inst := range insts {
		if inst.cmd != "jmp" {
			continue
		}
		newInsts := make([]instruction, len(insts))
		copy(newInsts, insts)
		newInsts[i].cmd = "nop"
		acc, finished := run(newInsts)
		if finished {
			return acc
		}
	}
	// Since the answer is actually changing a "jmp" to "nop", this for loop is
	// never run for this problem's particular input. But we leave this here for
	// posterity.
	for i, inst := range insts {
		if inst.cmd != "nop" {
			continue
		}
		newInsts := make([]instruction, len(insts))
		copy(newInsts, insts)
		newInsts[i].cmd = "jmp"
		acc, finished := run(newInsts)
		if finished {
			return acc
		}
	}

	return -1
}

func main() {
	file, err := os.Open("day08_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var insts []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instParts := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(instParts[1])
		if err != nil {
			log.Fatal(err)
		}
		insts = append(insts, instruction{
			cmd: instParts[0],
			val: val,
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	acc, _ := run(insts)
	fmt.Printf("Part 1: value of accumulator before infinite loop is: %d\n", acc)

	acc = runWithFix(insts)
	fmt.Printf("Part 2: value of accumulator after fix is: %d\n", acc)
}
