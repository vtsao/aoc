// This program implements the solution for
// https://adventofcode.com/2022/day/25.
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed day25_input.txt
var input string

// toBase10 translates a SNAFU number to a base 10 number.
func toBase10(snafu string) int {
	var num int
	for i, s := range snafu {
		var digit int
		switch s {
		case '-':
			digit = -1
		case '=':
			digit = -2
		default:
			digit, _ = strconv.Atoi(string(s))
		}
		num += digit * int(math.Pow(float64(5), float64(len(snafu)-1-i)))
	}
	return num
}

// sumFuelReq sums all the fuel requirements in the input file and returns the
// sum as a base 10 number.
func sumFuelReq() int {
	sum := 0
	for _, fuelReq := range strings.Split(input, "\n") {
		sum += toBase10(fuelReq)
	}
	return sum
}

// toBase5 translates a base 10 number to base 5 and returns its digits in a
// string slice.
func toBase5(num int) []string {
	numDigits := int(math.Log(float64(num)) / math.Log(float64(5)))
	var digits []string
	for i := numDigits; i >= 0; i-- {
		placeValue := int(math.Pow(float64(5), float64(i)))
		digit := num / placeValue
		digits = append(digits, strconv.Itoa(digit))
		num -= digit * placeValue
	}
	return digits
}

// toSNAFU translates a base 10 number to a SNAFU number.
func toSNAFU(num int) string {
	// First translate the number to base 5.
	snafu := toBase5(num)

	// Now we "fix" the base 5 number by taking care of any digits that are
	// greater than 2. Digits that are greater than 2 basically need to be
	// replaced with "-" or "=" and incrementing the previous digit, this gives us
	// the equivalent number in SNAFU.
	//
	// We continuously fix one digit at a time until the entire number is a valid
	// SNAFU number.
outer:
	for {
		for i, curDigit := range snafu {
			var newCurDigit string
			switch curDigit {
			case "4":
				newCurDigit = "-"
			case "3":
				newCurDigit = "="
			default:
				continue
			}

			snafu[i] = newCurDigit
			if i-1 >= 0 {
				snafu[i-1] = increment(snafu[i-1])
				continue outer
			}
			snafu = append([]string{"1"}, snafu...)

			continue outer
		}
		break
	}

	return strings.Join(snafu, "")
}

func increment(digit string) string {
	switch digit {
	case "-":
		return "0"
	case "=":
		return "-"
	// Note we purposely want to return "3" here when we're incrementing from "2"
	// because we're going from +ve to -ve, which requires the previous number to
	// also be incremented. Returning a "3" means the next iteration of the loop
	// will do this by "fixing" "3".
	default:
		d, _ := strconv.Atoi(digit)
		return strconv.Itoa(d + 1)
	}
}

func main() {
	fmt.Printf("Part 1: %q\n", toSNAFU(sumFuelReq()))
	// fmt.Printf("Part 2: %d\n", )
}
