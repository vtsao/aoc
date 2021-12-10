// This program implements the solution for
// https://adventofcode.com/2021/day/10.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("day10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	points1 := 0
	var incompleteStacks [][]rune
outer:
	for _, line := range lines {
		var stack []rune
		for _, r := range line {
			switch {
			case r == '(' || r == '[' || r == '{' || r == '<':
				stack = append(stack, r)
			case r == ')':
				if len(stack) == 0 {
					points1 += 3
					continue outer
				}
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != '(' {
					points1 += 3
					continue outer
				}
			case r == ']':
				if len(stack) == 0 {
					points1 += 57
					continue outer
				}
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != '[' {
					points1 += 57
					continue outer
				}
			case r == '}':
				if len(stack) == 0 {
					points1 += 1197
					continue outer
				}
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != '{' {
					points1 += 1197
					continue outer
				}
			case r == '>':
				if len(stack) == 0 {
					points1 += 25137
					continue outer
				}
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != '<' {
					points1 += 25137
					continue outer
				}
			}
		}
		if len(stack) > 0 {
			incompleteStacks = append(incompleteStacks, stack)
		}
	}
	fmt.Printf("Part 1: %d\n", points1)

	var points2 []int
	for _, stack := range incompleteStacks {
		points := 0
		for i := len(stack) - 1; i >= 0; i-- {
			r := stack[i]
			switch {
			case r == '(':
				points *= 5
				points += 1
			case r == '[':
				points *= 5
				points += 2
			case r == '{':
				points *= 5
				points += 3
			case r == '<':
				points *= 5
				points += 4
			}
		}
		points2 = append(points2, points)
	}
	sort.Ints(points2)
	fmt.Printf("Part 2: %d\n", points2[len(points2)/2])
}
