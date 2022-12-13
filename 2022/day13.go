// This program implements the solution for
// https://adventofcode.com/2022/day/13.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("day13_input.txt")
	defer file.Close()

	// Keep track of the sum of indexes of the packets in the right order for part
	// 1.
	sumRightOrderIdxs := 0

	// Keep track of all packets for part 2.
	div1 := parse("[[2]]")
	div1.isDiv = true
	div2 := parse("[[6]]")
	div2.isDiv = true
	packets := []*packet{div1, div2}

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		left := parse(scanner.Text())
		packets = append(packets, left)

		scanner.Scan()
		right := parse(scanner.Text())
		packets = append(packets, right)

		if isRightOrder(left, right) == 1 {
			sumRightOrderIdxs += i
		}

		i++
		scanner.Scan()
	}

	fmt.Printf("Part 1: %d\n", sumRightOrderIdxs)

	// Part 2 is just using the comparison from part 1 to do a sort with a custom
	// comparator.
	sort.Slice(packets, func(i, j int) bool {
		return isRightOrder(packets[i], packets[j]) == 1
	})
	// Find the indexes of the divider packets after all packets are sorted.
	decoderKey := 1
	for i, p := range packets {
		if p.isDiv {
			decoderKey *= i + 1
		}
	}
	fmt.Printf("Part 2: %d\n", decoderKey)
}

// packetStr parses a packet in string format into a packet struct. We use a
// stack to keep track of which packet to add values to.
func parse(packetStr string) *packet {
	var root *packet

	stack := []*packet{}
	for i := 0; i < len(packetStr); i++ {
		switch packetStr[i] {
		// This is a new value, add it to the packet that's at the top of the stack
		// if there is a packet at the top. Then always append this new packet to
		// the top of the stack.
		case '[':
			newPacket := &packet{}
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				top.list = append(top.list, newPacket)
			}
			stack = append(stack, newPacket)
		// Simply pop the packet stack and keep track of the last one popped as
		// that's the root packet.
		case ']':
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		// Just ignore commas.
		case ',':
		// Assuming valid input, everything else is an integer.
		default:
			num := string(packetStr[i])
			// There are double digit values, so look ahead until we find all
			// consecutive integers.
			for {
				if _, err := strconv.Atoi(string(packetStr[i+1])); err != nil {
					break
				}
				num += string(packetStr[i+1])
				i++
			}
			top := stack[len(stack)-1]
			v, _ := strconv.Atoi(num)
			top.list = append(top.list, &packet{val: v, isVal: true})
		}
	}

	return root
}

type packet struct {
	val   int
	list  []*packet
	isVal bool
	isDiv bool
}

func isRightOrder(left, right *packet) int {
	switch {
	case left.isVal && right.isVal:
		if left.val < right.val {
			return 1
		}
		if left.val > right.val {
			return -1
		}
		return 0
	case !left.isVal && !right.isVal:
		for i, lp := range left.list {
			if i >= len(right.list) {
				break
			}
			o := isRightOrder(lp, right.list[i])
			if o == -1 {
				return -1
			}
			if o == 1 {
				return 1
			}
		}
		if len(left.list) < len(right.list) {
			return 1
		}
		if len(left.list) > len(right.list) {
			return -1
		}
		return 0
	case left.isVal && !right.isVal:
		lp := &packet{
			isVal: false,
			list: []*packet{
				{
					val:   left.val,
					isVal: true,
				},
			},
		}
		return isRightOrder(lp, right)
	case !left.isVal && right.isVal:
		rp := &packet{
			isVal: false,
			list: []*packet{
				{
					val:   right.val,
					isVal: true,
				},
			},
		}
		return isRightOrder(left, rp)
	}

	// Will never reach here with valid input.
	return 0
}
