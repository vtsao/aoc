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

func parseBits(binaryStr string) (string, int, int) {
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
		return binaryStr[parsedBits:], parsedBits, packetVerSum
	}

	log.Printf("Packet type operator...\n")
	remaining := ""
	totalParsedBits := 7
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
			newRemaining, newParsedBits, verSum := parseBits(remaining)
			remaining = newRemaining
			parsedBits += newParsedBits
			packetVerSum += verSum
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
			newRemaining, parsedBits, verSum := parseBits(remaining)
			remaining = newRemaining
			totalParsedBits += parsedBits
			packetVerSum += verSum
			log.Printf("Parsed %d sub-packets, want %d\n", i, numSubPackets)
		}
	}

	return remaining, totalParsedBits, packetVerSum
}

func main() {
	binaryStr := parseInput()

	_, _, packetVerSum := parseBits(binaryStr)
	fmt.Printf("Part 1: %d\n", packetVerSum)
	// fmt.Printf("Part 2: %d\n", )
}
