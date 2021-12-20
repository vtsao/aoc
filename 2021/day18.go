// This program implements the solution for
// https://adventofcode.com/2021/day/18.
//
// curl -b "$(cat .session)" -o day18_input.txt https://adventofcode.com/2021/day/18/input
package main

import (
	"bufio"
	"fmt"
	"log"
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

	// log.Printf("%s, %d\n", number, i)
	// log.Println(closingParenIdx)
	// log.Println(leftPartNumber)
	// log.Println(rightPartNumber)
	// log.Println(leftVal)
	// log.Println(rightVal)

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
		log.Println(number)
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
			case ',':
			default:
				if next := number[i+1]; next == '[' || next == ']' || next == ',' {
					continue
				}
				number = split(number, i)
				continue reducing
			}
		}
		break
	}

	return number
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

	sum := reduce(numbers[0])
	for i := 1; i < len(numbers); i++ {
		nextNum := reduce(numbers[i])
		sum = fmt.Sprintf("[%s,%s]", sum, nextNum)
		sum = reduce(sum)
	}

	fmt.Printf("Part 1: %s\n", sum)
	// fmt.Printf("Part 2: %d\n", )
}
