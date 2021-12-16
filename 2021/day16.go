// This program implements the solution for
// https://adventofcode.com/2021/day/16.
//
// curl -b "$(cat .session)" -o day16_input.txt https://adventofcode.com/2021/day/16/input
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInput() string {
	file, _ := os.Open("day16_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	binaryStr := ""
	for _, hexNum := range scanner.Text() {
		num, _ := strconv.ParseInt(string(hexNum), 16, 64)
		binaryStr += fmt.Sprintf("%04b", num)
	}

	return binaryStr
}

func parseBits(binaryStr string) (string, int, int, int) {
	log.Printf("Parsing binary string %q\n", binaryStr)
	packetVer, _ := strconv.ParseInt(binaryStr[:3], 2, 64)
	packetVerSum := int(packetVer)
	packetTypeID, _ := strconv.ParseInt(binaryStr[3:6], 2, 64)

	if packetTypeID == 4 {
		log.Printf("Packet type 4...\n")
		literalValStr := ""
		parsedBits := 6
		for start := 6; ; start += 5 {
			chunk := binaryStr[start : start+5]
			literalValStr += chunk[1:]
			parsedBits += 5
			if chunk[0] == '0' {
				break
			}
		}

		literalVal, _ := strconv.ParseInt(literalValStr, 2, 64)
		log.Printf("literal value %d\n", literalVal)
		return binaryStr[parsedBits:], parsedBits, packetVerSum, int(literalVal)
	}

	log.Printf("Packet type operator...\n")
	remaining := ""
	totalParsedBits := 7
	var subPacketVals []int
	switch binaryStr[6] {
	case '0':
		subPacketLenBinaryStr := binaryStr[7:22]
		totalParsedBits += 15
		subPacketLenInt64, _ := strconv.ParseInt(subPacketLenBinaryStr, 2, 64)
		subPacketLen := int(subPacketLenInt64)
		totalParsedBits += subPacketLen
		log.Printf("Operator type 0, total sub-packets of %d length\n", subPacketLen)

		parsedBits := 0
		for remaining = binaryStr[22:]; parsedBits < subPacketLen; {
			newRemaining, newParsedBits, newPacketVerSum, newPacketVal := parseBits(remaining)
			remaining = newRemaining
			parsedBits += newParsedBits
			packetVerSum += newPacketVerSum
			subPacketVals = append(subPacketVals, newPacketVal)
			log.Printf("Parsed %d packets, want %d\n", parsedBits, subPacketLen)
		}
	case '1':
		numSubPacketsBinaryStr := binaryStr[7:18]
		totalParsedBits += 11
		numSubPacketsInt64, _ := strconv.ParseInt(numSubPacketsBinaryStr, 2, 64)
		numSubPackets := int(numSubPacketsInt64)
		log.Printf("Operator type 0, total %d sub-packets\n", numSubPackets)

		remaining = binaryStr[18:]
		for i := 0; i < numSubPackets; i++ {
			newRemaining, parsedBits, newPacketVerSum, newPacketVal := parseBits(remaining)
			remaining = newRemaining
			totalParsedBits += parsedBits
			packetVerSum += newPacketVerSum
			subPacketVals = append(subPacketVals, newPacketVal)
			log.Printf("Parsed %d sub-packets, want %d\n", i, numSubPackets)
		}
	}

	packetVal := 0
	switch packetTypeID {
	case 0:
		for _, val := range subPacketVals {
			packetVal += val
		}
	case 1:
		packetVal = subPacketVals[0]
		for i := 1; i < len(subPacketVals); i++ {
			packetVal *= subPacketVals[i]
		}
	case 2:
		packetVal = subPacketVals[0]
		for i := 1; i < len(subPacketVals); i++ {
			if val := subPacketVals[i]; val < packetVal {
				packetVal = val
			}
		}
	case 3:
		packetVal = subPacketVals[0]
		for i := 1; i < len(subPacketVals); i++ {
			if val := subPacketVals[i]; val > packetVal {
				packetVal = val
			}
		}
	case 5:
		if subPacketVals[0] > subPacketVals[1] {
			packetVal = 1
		}
	case 6:
		if subPacketVals[0] < subPacketVals[1] {
			packetVal = 1
		}
	case 7:
		if subPacketVals[0] == subPacketVals[1] {
			packetVal = 1
		}
	}

	return remaining, totalParsedBits, packetVerSum, packetVal
}

func main() {
	binaryStr := parseInput()

	_, _, packetVerSum, packetVal := parseBits(binaryStr)
	fmt.Printf("Part 1: %d\n", packetVerSum)
	fmt.Printf("Part 2: %d\n", packetVal)
}
