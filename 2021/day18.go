// This program implements the solution for
// https://adventofcode.com/2021/day/18.
//
// curl -b "$(cat .session)" -o day18_input.txt https://adventofcode.com/2021/day/18/input
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func explode(number string, i int) string {
	closingParenIdx := i + strings.IndexRune(number[i:], ']')
	leftPartNumber := number[:i]
	rightPartNumber := number[closingParenIdx+1:]
	pairParts := strings.Split(number[i+1:closingParenIdx], ",")
	leftVal, _ := strconv.Atoi(pairParts[0])
	rightVal, _ := strconv.Atoi(pairParts[1])

	for j := len(leftPartNumber) - 1; j >= 0; j-- {
		r := leftPartNumber[j]
		if r == '[' || r == ']' || r == ',' {
			continue
		}
		next := leftPartNumber[j-1]
		if next != '[' && next != ']' && next != ',' {
			regularNum, _ := strconv.Atoi(string(next) + string(r))
			leftPartNumber = fmt.Sprintf("%s%d%s", leftPartNumber[:j-1], regularNum+leftVal, leftPartNumber[j+1:])
			break
		}

		regularNum, _ := strconv.Atoi(string(r))
		leftPartNumber = fmt.Sprintf("%s%d%s", leftPartNumber[:j], regularNum+leftVal, leftPartNumber[j+1:])
		break
	}
	for j := 0; j < len(rightPartNumber); j++ {
		r := rightPartNumber[j]
		if r == '[' || r == ']' || r == ',' {
			continue
		}
		next := rightPartNumber[j+1]
		if next != '[' && next != ']' && next != ',' {
			regularNum, _ := strconv.Atoi(string(r) + string(next))
			rightPartNumber = fmt.Sprintf("%s%d%s", rightPartNumber[:j], regularNum+rightVal, rightPartNumber[j+2:])
			break
		}

		regularNum, _ := strconv.Atoi(string(r))
		rightPartNumber = fmt.Sprintf("%s%d%s", rightPartNumber[:j], regularNum+rightVal, rightPartNumber[j+1:])
		break
	}

	return fmt.Sprintf("%s%d%s", leftPartNumber, 0, rightPartNumber)
}

func split(number string, i int) string {
	regularNum, _ := strconv.Atoi(string(number[i]) + string(number[i+1]))
	leftVal := regularNum / 2
	rightVal := int(math.Ceil(float64(regularNum) / float64(2)))
	return fmt.Sprintf("%s[%d,%d]%s", number[:i], leftVal, rightVal, number[i+2:])
}

func reduce(number string) string {
reducing:
	for {
		parens := 0
		for i, r := range number {
			switch r {
			case '[':
				parens++
				if parens <= 4 {
					continue
				}
				number = explode(number, i)
				continue reducing
			case ']':
				parens--
			}
		}
		for i, r := range number {
			if r == '[' || r == ']' || r == ',' {
				continue
			}
			if next := number[i+1]; next == '[' || next == ']' || next == ',' {
				continue
			}
			number = split(number, i)
			continue reducing
		}
		break
	}

	return number
}

func mag(number string) int {
	val, err := strconv.Atoi(number)
	if err == nil {
		return val
	}

	pairMid := 0
	parens := 0
	for i := 1; i < len(number)-1; i++ {
		r := number[i]
		switch r {
		case '[':
			parens++
		case ']':
			parens--
		case ',':
			if parens == 0 {
				pairMid = i
			}
		}
	}

	left := mag(number[1:pairMid])
	right := mag(number[pairMid+1 : len(number)-1])

	return left*3 + 2*right
}

func largestMag(numbers []string) int {
	largest := 0
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if mag := mag(reduce(fmt.Sprintf("[%s,%s]", numbers[i], numbers[j]))); mag > largest {
				largest = mag
			}
			if mag := mag(reduce(fmt.Sprintf("[%s,%s]", numbers[j], numbers[i]))); mag > largest {
				largest = mag
			}
		}
	}
	return largest
}

func parseInput() []string {
	file, _ := os.Open("day18_input.txt")
	defer file.Close()

	var numbers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	return numbers
}

func main() {
	numbers := parseInput()

	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		nextNum := numbers[i]
		sum = fmt.Sprintf("[%s,%s]", sum, nextNum)
		sum = reduce(sum)
	}

	fmt.Printf("Part 1: %d\n", mag(sum))
	fmt.Printf("Part 2: %d\n", largestMag(numbers))
}
