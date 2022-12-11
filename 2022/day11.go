// This program implements the solution for
// https://adventofcode.com/2022/day/11.
package main

import (
	"fmt"
	"sort"
)

func main() {
	monkeyBusiness1 := rounds(monkeys(), 20, 3)
	monkeyBusiness2 := rounds(monkeys(), 10000, 1)

	fmt.Printf("Part 1: %d\n", monkeyBusiness1)
	fmt.Printf("Part 2: %d\n", monkeyBusiness2)
}

func rounds(monkeys []*monkey, numRounds int, worryManagement int) int {
	for i := 0; i < numRounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				newItem := monkey.op(item)
				newItem /= worryManagement
				monkey.test(newItem, monkeys)
				monkey.numInspects++
			}
			monkey.items = nil
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		// Sort in desc order.
		return monkeys[i].numInspects > monkeys[j].numInspects
	})
	return monkeys[0].numInspects * monkeys[1].numInspects
}

func monkeys() []*monkey {
	return []*monkey{
		{
			items: []int{65, 58, 93, 57, 66},
			op: func(old int) int {
				return old * 7
			},
			test: func(item int, monkeys []*monkey) {
				if item%19 == 0 {
					monkeys[6].items = append(monkeys[6].items, item)
					return
				}
				monkeys[4].items = append(monkeys[4].items, item)
			},
		},
		{
			items: []int{76, 97, 58, 72, 57, 92, 82},
			op: func(old int) int {
				return old + 4
			},
			test: func(item int, monkeys []*monkey) {
				if item%3 == 0 {
					monkeys[7].items = append(monkeys[7].items, item)
					return
				}
				monkeys[5].items = append(monkeys[5].items, item)
			},
		},
		{
			items: []int{90, 89, 96},
			op: func(old int) int {
				return old * 5
			},
			test: func(item int, monkeys []*monkey) {
				if item%13 == 0 {
					monkeys[5].items = append(monkeys[5].items, item)
					return
				}
				monkeys[1].items = append(monkeys[1].items, item)
			},
		},
		{
			items: []int{72, 63, 72, 99},
			op: func(old int) int {
				return old * old
			},
			test: func(item int, monkeys []*monkey) {
				if item%17 == 0 {
					monkeys[0].items = append(monkeys[0].items, item)
					return
				}
				monkeys[4].items = append(monkeys[4].items, item)
			},
		},
		{
			items: []int{65},
			op: func(old int) int {
				return old + 1
			},
			test: func(item int, monkeys []*monkey) {
				if item%2 == 0 {
					monkeys[6].items = append(monkeys[6].items, item)
					return
				}
				monkeys[2].items = append(monkeys[2].items, item)
			},
		},
		{
			items: []int{97, 71},
			op: func(old int) int {
				return old + 8
			},
			test: func(item int, monkeys []*monkey) {
				if item%11 == 0 {
					monkeys[7].items = append(monkeys[7].items, item)
					return
				}
				monkeys[3].items = append(monkeys[3].items, item)
			},
		},
		{
			items: []int{83, 68, 88, 55, 87, 67},
			op: func(old int) int {
				return old + 2
			},
			test: func(item int, monkeys []*monkey) {
				if item%5 == 0 {
					monkeys[2].items = append(monkeys[2].items, item)
					return
				}
				monkeys[1].items = append(monkeys[1].items, item)
			},
		},
		{
			items: []int{64, 81, 50, 96, 82, 53, 62, 92},
			op: func(old int) int {
				return old + 5
			},
			test: func(item int, monkeys []*monkey) {
				if item%7 == 0 {
					monkeys[3].items = append(monkeys[3].items, item)
					return
				}
				monkeys[0].items = append(monkeys[0].items, item)
			},
		},
	}
}

type monkey struct {
	items       []int
	op          func(old int) int
	test        func(item int, monkeys []*monkey)
	numInspects int
}
