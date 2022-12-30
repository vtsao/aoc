// This program implements the solution for
// https://adventofcode.com/2022/day/21.
package main

import (
	"fmt"
	"strings"

	_ "embed"
)

//go:embed day21_input.txt
var input string

// parse parses the input file of job strings to job structs.
func parse() map[string]job {
	jobs := map[string]job{}
	for _, jobStr := range strings.Split(input, "\n") {
		// Handle number-yelling monkeys like "sllz: 4".
		var monkey string
		var num int
		if _, err := fmt.Sscanf(jobStr, "%s %d", &monkey, &num); err == nil {
			monkey = strings.TrimRight(monkey, ":")
			jobs[monkey] = job{num: num}
			continue
		}

		// Handle math operation monkeys like "sjmn: drzm * dbpl".
		var left, right, op string
		if _, err := fmt.Sscanf(jobStr, "%s %s %s %s", &monkey, &left, &op, &right); err == nil {
			monkey = strings.TrimRight(monkey, ":")
			jobs[monkey] = job{left: left, op: op, right: right}
		}
	}

	return jobs
}

func main() {
	jobs := parse()
	fmt.Printf("Part 1: %d\n", *number("root", jobs, false))
	fmt.Printf("Part 2: %d\n", humn(jobs))
}

type job struct {
	num             int
	left, right, op string
}

// number calculates the number a monkey will yell. If unknownHumn is set to
// true, it will return nil if a calculation requires the value of "humn" to be
// known.
func number(monkey string, jobs map[string]job, unknownHumn bool) *int {
	if unknownHumn && monkey == "humn" {
		return nil
	}

	// If this monkey knows its number, just return it.
	job := jobs[monkey]
	if job.left == "" {
		return &job.num
	}

	// Otherwise, recursively calculate the number this monkey will yell.
	left := number(job.left, jobs, unknownHumn)
	right := number(job.right, jobs, unknownHumn)
	if left == nil || right == nil {
		return nil
	}
	var result int
	switch job.op {
	case "+":
		result = *left + *right
	case "-":
		result = *left - *right
	case "*":
		result = *left * *right
	case "/":
		result = *left / *right
	}

	return &result
}

// humn calculates the value that "humn" needs to yell to make the left and
// right numbers of "root" equal.
func humn(jobs map[string]job) int {
	monkey := "root"
	var wantNum int

	// Work backwards from the target number to find what "humn" should be.
	//
	// For example:
	//
	// root: pppw == sjmn (150), so pppw is also 150
	// pppw (150) == cczh / lfqf (4), so pppw (150) * lfqf (4) == cczh (600)
	// cczh (600) == sllz (4) + lgvd, so cczh (600) - sllz (4) == lgvd (596)
	// lgvd (596) == ljgn (2) * ptdq, so lgvd (596) / ljgn (2) == ptdq (298)
	// ptdq (298) == humn - dvpt (3) --> ptdq (298) + dvpt (3) == humn (301)
	for monkey != "humn" {
		j := jobs[monkey]
		// "root" is now equality.
		if monkey == "root" {
			j = job{left: jobs["root"].left, right: jobs["root"].right, op: "="}
		}
		// One of left or right is always calculatable, meaning it does not depend
		// on "humn". So we use this invariant to allow us to work backwards.
		if n := number(j.left, jobs, true); n != nil {
			wantNum = inverse(wantNum, *n, j.op, true)
			monkey = j.right
			continue
		}
		wantNum = inverse(wantNum, *number(j.right, jobs, true), j.op, false)
		monkey = j.left
	}

	return wantNum
}

// inverse inverts a math operation that is in the form of
// (num == left <op> right) to calculate the unknown variable, given that we
// know the value of num and the value of exactly one of left or right.
func inverse(wantNum, knownNum int, op string, knownNumIsLeft bool) int {
	switch op {

	// root: cczh = lfqf
	// This is a special case just for "root", no other job should use this op.
	case "=":
		return knownNum

	// pppw == cczh + lfqf
	case "+":
		// pppw - cczh == lfqf
		// pppw - lfqf == cczh
		return wantNum - knownNum

	// pppw == cczh - lfqf
	case "-":
		// cczh - pppw == lfqf
		if knownNumIsLeft {
			return knownNum - wantNum
		}
		// pppw + lfqf == cczh
		return wantNum + knownNum

	// pppw == cczh * lfqf
	case "*":
		// pppw / cczh == lfqf
		// pppw / lfqf == cczh
		return wantNum / knownNum

	// pppw == cczh / lfqf
	case "/":
		// cczh / pppw == lfqf
		if knownNumIsLeft {
			return knownNum / wantNum
		}
		// pppw * lfqf== cczh
		return wantNum * knownNum
	}

	// Will never arrive here with valid input.
	return 0
}
