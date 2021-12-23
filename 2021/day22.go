// This program implements the solution for
// https://adventofcode.com/2021/day/22.
//
// curl -b "$(cat .session)" -o day22_input.txt https://adventofcode.com/2021/day/22/input
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type kuboid struct {
	xStart, xEnd, yStart, yEnd, zStart, zEnd int
}

type instruction struct {
	on   bool
	kube kuboid
}

func parseInput() []instruction {
	file, _ := os.Open("day22_input.txt")
	defer file.Close()

	var instructions []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructionParts := strings.Split(scanner.Text(), " ")
		on := false
		if instructionParts[0] == "on" {
			on = true
		}

		coordParts := strings.Split(instructionParts[1], ",")

		xCoordStr := strings.TrimPrefix(coordParts[0], "x=")
		xCoordParts := strings.Split(xCoordStr, "..")
		xStart, _ := strconv.Atoi(xCoordParts[0])
		xEnd, _ := strconv.Atoi(xCoordParts[1])

		yCoordStr := strings.TrimPrefix(coordParts[1], "y=")
		yCoordParts := strings.Split(yCoordStr, "..")
		yStart, _ := strconv.Atoi(yCoordParts[0])
		yEnd, _ := strconv.Atoi(yCoordParts[1])

		zCoordStr := strings.TrimPrefix(coordParts[2], "z=")
		zCoordParts := strings.Split(zCoordStr, "..")
		zStart, _ := strconv.Atoi(zCoordParts[0])
		zEnd, _ := strconv.Atoi(zCoordParts[1])

		instructions = append(instructions, instruction{
			on: on,
			kube: kuboid{
				xStart: xStart,
				xEnd:   xEnd,
				yStart: yStart,
				yEnd:   yEnd,
				zStart: zStart,
				zEnd:   zEnd,
			},
		})
	}

	return instructions
}

type incExc struct {
	kube    kuboid
	include bool
}

func intersect(k1, k2 kuboid) (kuboid, bool) {
	k := kuboid{}

	k.xStart = int(math.Max(float64(k1.xStart), float64(k2.xStart)))
	k.xEnd = int(math.Min(float64(k1.xEnd), float64(k2.xEnd)))
	k.yStart = int(math.Max(float64(k1.yStart), float64(k2.yStart)))
	k.yEnd = int(math.Min(float64(k1.yEnd), float64(k2.yEnd)))
	k.zStart = int(math.Max(float64(k1.zStart), float64(k2.zStart)))
	k.zEnd = int(math.Min(float64(k1.zEnd), float64(k2.zEnd)))

	if k.xStart <= k.xEnd && k.yStart <= k.yEnd && k.zStart <= k.zEnd {
		return k, true
	}

	return kuboid{}, false
}

func reboot(instructions []instruction, init bool) int {
	var incExcs []incExc

	for _, instruction := range instructions {
		kube := instruction.kube
		if init && (kube.xStart < -50 || kube.xEnd > 50 || kube.yStart < -50 || kube.yEnd > 50 || kube.zStart < -50 || kube.zEnd > 50) {
			continue
		}
		for _, ie := range incExcs {
			intersectKube, ok := intersect(kube, ie.kube)
			if !ok {
				continue
			}

			incExcs = append(incExcs, incExc{kube: intersectKube, include: !ie.include})
		}
		if instruction.on {
			incExcs = append(incExcs, incExc{kube: kube, include: true})
		}
	}

	onReactors := 0
	for _, ie := range incExcs {
		k := ie.kube
		reactors := (k.xEnd + 1 - k.xStart) * (k.yEnd + 1 - k.yStart) * (k.zEnd + 1 - k.zStart)
		if ie.include {
			onReactors += reactors
			continue
		}
		onReactors -= reactors
	}

	return onReactors
}

func main() {
	instructions := parseInput()
	fmt.Printf("Part 1: %d\n", reboot(instructions, true))
	fmt.Printf("Part 2: %d\n", reboot(instructions, false))
}
