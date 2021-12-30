// This program implements the solution for
// https://adventofcode.com/2021/day/24.
//
// curl -b "$(cat .session)" -o day24_input.txt https://adventofcode.com/2021/day/24/input
package main

import (
	"fmt"
	"math"
	"strconv"
)

// These are taken from the input file. Each of the 14 instructions in the input
// file are identical except for these variables. We could have parsed this from
// the input file, but it's easier to just enter these by hand.
var zDivs = []int{1, 1, 1, 26, 1, 26, 26, 1, 1, 1, 26, 26, 26, 26}
var xAdds = []int{10, 14, 14, -13, 10, -13, -7, 11, 10, 13, -4, -9, -13, -9}
var yAdds = []int{2, 13, 13, 9, 15, 3, 6, 5, 16, 1, 6, 3, 7, 9}

// exec executes the nth MONAD instruction specified in the input file (there
// are 14 instructions in total). The code here is a reduction/simplification of
// the ALU code from the input file. It was manually inspected and written here.
func exec(z, digit, instIdx int) int {
	if z%26+xAdds[instIdx] == digit {
		return z / zDivs[instIdx]
	}
	return (z/zDivs[instIdx])*26 + digit + yAdds[instIdx]
}

// findMinMaxModelNos recursively runs the MONAD instructions on all 14 digit
// numbers that contain no 0 digits. Normally this would be brute force and
// would not finish in a reasonable time, but intuition tells us that there must
// be some way to prune the recursion tree to make it finish in a reasonable
// time. We figure out how we can prune by examining the input instructions.
func findMinMaxModelNos(z, instIdx int, modelNo string) (int, int) {
	if instIdx == 14 {
		if z == 0 {
			num, _ := strconv.Atoi(modelNo)
			return num, num
		}
		return math.MaxInt64, 0
	}

	smallest, largest := math.MaxInt64, 0
	for i := 1; i < 10; i++ {
		// This prune condition is the key to this entire problem. Reducing the ALU
		// instructions to what we wrote in the exec() function is trivial as is
		// noting the zDivs, xAdds, and yAdds variables that are different between
		// the instructions. But this prune step was based on a guess.
		//
		// Basically exec() has two ways it can modify z and it basically depends on
		// the value in xAdds for the index, since there will always be z values for
		// z%26 that can be equal to the digit. But the xAdds value can make this
		// impossible. Specifically for digits with a positive xAdds value, the
		// first condition in exec() can never be reached. This means z is
		// multiplied by 26 seven times (seven of the instructions are guaranteed
		// to multiply z by 26). There is also some other noise like it adds the
		// digit and value from yAdds, but it turns out those are irrelevant to the
		// pruning.
		//
		// The other seven xAdds values are negative, which means the first
		// condition in exec() CAN be taken (but not always), which divides z by 26
		// (look at how zDivs matches up with xAdds). We assume we WANT these seven
		// cases where the xAdds value is negative to be taken, otherwise z will not
		// be anywhere near 0. So for each instruction where the xAdds value is < 0
		// and we don't take the first condition, then we prune (not recurse).
		//
		// This turns out to prune the recursion tree significantly and it finishes
		// in a second or so.
		if xAdds[instIdx] < 0 && z%26+xAdds[instIdx] != i {
			continue
		}

		newZ := exec(z, i, instIdx)
		small, large := findMinMaxModelNos(newZ, instIdx+1, fmt.Sprintf("%s%d", modelNo, i))
		if large > largest {
			largest = large
		}
		if small < smallest {
			smallest = small
		}
	}

	return smallest, largest
}

func main() {
	smallest, largest := findMinMaxModelNos(0, 0, "")
	fmt.Printf("Part 1: %d\n", largest)
	fmt.Printf("Part 2: %d\n", smallest)
}
