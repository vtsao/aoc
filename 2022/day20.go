// This program implements the solution for
// https://adventofcode.com/2022/day/20.
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed day20_input.txt
var input string

// parse parses the encrypted file into a doubly linked list of nodes, keeping
// track of each node by its original index in the file. It also finds the index
// of the value 0 (there is only one 0 value for valid input).
func parse(decryptionKey int) (*node, []*node) {
	var zeroNode *node
	var nodes []*node

	for i, numStr := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(numStr)
		num *= decryptionKey
		n := &node{num: num}
		if i > 0 {
			n.prev = nodes[i-1]
			n.prev.next = n
		}
		nodes = append(nodes, n)

		if num == 0 {
			zeroNode = n
		}
	}
	nodes[0].prev = nodes[len(nodes)-1]
	nodes[0].prev.next = nodes[0]

	return zeroNode, nodes
}

func main() {
	zeroNode, nodes := parse(1)
	mix(nodes)
	fmt.Printf("Part 1: %d\n", sumCoords(zeroNode, nodes))

	zeroNode, nodes = parse(811589153)
	for i := 0; i < 10; i++ {
		mix(nodes)
	}
	fmt.Printf("Part 2: %d\n", sumCoords(zeroNode, nodes))
}

func sumCoords(zeroNode *node, nodes []*node) int {
	sumCoords := 0
	for _, pos := range []int{1000, 2000, 3000} {
		n := zeroNode
		for i := 0; i < pos%len(nodes); i++ {
			n = n.next
		}
		sumCoords += n.num
	}
	return sumCoords
}

type node struct {
	num        int
	next, prev *node
}

func mix(nodes []*node) {
	for i := 0; i < len(nodes); i++ {
		n := nodes[i]

		// Find the node we will be moving our current node to. We will be moving
		// the current node in front of the "move to" node.
		//
		// For example if we have [1, [2], -3, 10, 5, 2, 0] and we are moving [2],
		// the "move to" node will be the node with value 10.
		moveTo := n
		num := mod(n.num, len(nodes))
		for j := 0; j < num; j++ {
			moveTo = moveTo.next
		}
		for j := num; j <= 0; j++ {
			moveTo = moveTo.prev
		}

		// Remove the current node.
		n.prev.next = n.next
		n.next.prev = n.prev

		// Update the current node's prev and next to where it's been moved to.
		n.prev = moveTo
		n.next = moveTo.next

		// Update the nodes surrounding (before and after) where the current node
		// has been moved to, to point to the current node.
		moveTo.next = n
		n.next.prev = n
	}
}

// mod calculates the remainder for the number we need to move a number. This is
// needed for part 2 where the number is huge, but much larger than the number
// of nodes.
//
// We subtract one from the number of nodes because when moving a number that
// wraps around, we don't want to count that number itself when calculating the
// position to move it to. For example, if we have [1 3 1] and we move 3, we
// want to end up with [1 1 3].
func mod(num, numNodes int) int {
	absNum := int(math.Abs(float64(num)))
	remainder := absNum % (numNodes - 1)
	if num < 0 {
		return -remainder
	}
	return remainder
}
