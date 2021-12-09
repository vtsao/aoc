// This program implements the solution for https://adventofcode.com/2021/day/8.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entry struct {
	nums    map[string]map[rune]interface{}
	outputs []string
}

func main() {
	file, err := os.Open("day08_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var entries []*entry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entryParts := strings.Split(scanner.Text(), " | ")

		// For each scrambled number string, sort it by its characters to get a key
		// so we have a way of referring to it (they are all unique).
		//
		// E.g., if the scrambled number string is "ffcc", then its key is "ccff".
		nums := map[string]map[rune]interface{}{}
		for _, rawNum := range strings.Split(entryParts[0], " ") {
			rawNumParts := strings.Split(rawNum, "")
			sort.Strings(rawNumParts)
			numKey := strings.Join(rawNumParts, "")

			// Store each sequence character in a hash set so we can look them up
			// quickly during decoding later on.
			nums[numKey] = map[rune]interface{}{}
			for _, r := range numKey {
				nums[numKey][r] = nil
			}
		}

		// Do the same thing to get the keys for the output numbers. We have to do
		// this key process again because the output numbers may be scrambled in a
		// different order than their counterpart in the scrambled numbers section.
		var outputs []string
		for _, rawNum := range strings.Split(entryParts[1], " ") {
			rawNumParts := strings.Split(rawNum, "")
			sort.Strings(rawNumParts)
			num := strings.Join(rawNumParts, "")
			outputs = append(outputs, num)
		}

		entries = append(entries, &entry{nums: nums, outputs: outputs})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Decode the scrambled number strings.
	easyDigitCnt := 0
	totalOutput := 0
	for _, entry := range entries {
		numKeyToDecodedNum := map[string]int{}
		decodedNumToSeqSet := map[int]map[rune]interface{}{}
		decoded := map[string]interface{}{}

		// Here we're going to decode each of the scrambled number strings using
		// deduction based on the constraints of the problem. Probably not the most
		// optimal way of doing this.

		// Find the "easy" numbers first - the ones with unique number of sequences.

		// Find 1.
		for numKey, encodedNum := range entry.nums {
			if len(encodedNum) != 2 {
				continue
			}
			numKeyToDecodedNum[numKey] = 1
			decodedNumToSeqSet[1] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}
		// Find 4.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 4 {
				continue
			}
			numKeyToDecodedNum[numKey] = 4
			decodedNumToSeqSet[4] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}
		// Find 7.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 3 {
				continue
			}
			numKeyToDecodedNum[numKey] = 7
			decodedNumToSeqSet[7] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}
		// Find 8.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 7 {
				continue
			}
			numKeyToDecodedNum[numKey] = 8
			decodedNumToSeqSet[8] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}

		// Now find the harder numbers.

		// Find 3. 3 is the only 5 segment number that has 2 of its segments appear
		// in 1.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 5 {
				continue
			}
			overlapping := 0
			for r := range encodedNum {
				if _, ok := decodedNumToSeqSet[1][r]; ok {
					overlapping++
				}
			}
			if overlapping != 2 {
				continue
			}
			numKeyToDecodedNum[numKey] = 3
			decodedNumToSeqSet[3] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}

		// Find 5. 5 is the only 5 segment number left at this point that has 3 of
		// its segments appear in 4.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 5 {
				continue
			}
			overlapping := 0
			for r := range encodedNum {
				if _, ok := decodedNumToSeqSet[4][r]; ok {
					overlapping++
				}
			}
			if overlapping != 3 {
				continue
			}
			numKeyToDecodedNum[numKey] = 5
			decodedNumToSeqSet[5] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}

		// Find 2. 2 is now the only 5 segment number left.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 5 {
				continue
			}
			numKeyToDecodedNum[numKey] = 2
			decodedNumToSeqSet[2] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}

		// Find 9. 9 is the only 6 segment number that has 4 of its segments appear
		// in 4.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 6 {
				continue
			}
			overlapping := 0
			for r := range encodedNum {
				if _, ok := decodedNumToSeqSet[4][r]; ok {
					overlapping++
				}
			}
			if overlapping != 4 {
				continue
			}
			numKeyToDecodedNum[numKey] = 9
			decodedNumToSeqSet[9] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}

		// Find 6. 6 is the only 6 segment number left at this point that has 5 of
		// its segments appear in 5.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 6 {
				continue
			}
			overlapping := 0
			for r := range encodedNum {
				if _, ok := decodedNumToSeqSet[5][r]; ok {
					overlapping++
				}
			}
			if overlapping != 5 {
				continue
			}
			numKeyToDecodedNum[numKey] = 6
			decodedNumToSeqSet[6] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}

		// Find 0. 0 is now the only 6 segment number left.
		for numKey, encodedNum := range entry.nums {
			if _, ok := decoded[numKey]; ok {
				continue
			}
			if len(encodedNum) != 6 {
				continue
			}
			numKeyToDecodedNum[numKey] = 0
			decodedNumToSeqSet[0] = entry.nums[numKey]
			decoded[numKey] = nil
			break
		}

		// Find out what the outputs are.
		outputStr := ""
		for _, numKey := range entry.outputs {
			// Part 1.
			if num, ok := numKeyToDecodedNum[numKey]; ok {
				if num == 1 || num == 4 || num == 7 || num == 8 {
					easyDigitCnt++
				}
			}

			// Part 2.
			outputStr += fmt.Sprint(numKeyToDecodedNum[numKey])
		}
		output, err := strconv.Atoi(outputStr)
		if err != nil {
			log.Fatal(err)
		}
		totalOutput += output
	}

	fmt.Printf("Part 1: %d\n", easyDigitCnt)
	fmt.Printf("Part 2: %d\n", totalOutput)
}
