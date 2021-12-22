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

type coord struct {
	x, y, z int
}

func reboot1(instructions []instruction) int {
	onReactors := map[coord]struct{}{}
	for _, instruction := range instructions {
		kube := instruction.kube
		if kube.xStart < -50 || kube.xEnd > 50 || kube.yStart < -50 || kube.yEnd > 50 || kube.zStart < -50 || kube.zEnd > 50 {
			continue
		}

		for i := kube.xStart; i <= kube.xEnd; i++ {
			for j := kube.yStart; j <= kube.yEnd; j++ {
				for k := kube.zStart; k <= kube.zEnd; k++ {
					c := coord{x: i, y: j, z: k}
					if instruction.on {
						onReactors[c] = struct{}{}
						continue
					}
					delete(onReactors, c)
				}
			}
		}
	}

	return len(onReactors)
}

type incExc struct {
	kube    kuboid
	include bool
}

func intersect(k1, k2 kuboid) (kuboid, bool) {
	k := kuboid{}

	if k1.xStart >= k2.xStart && k1.xStart <= k2.xEnd {
		k.xStart = k1.xStart
		k.xEnd = int(math.Min(float64(k1.xEnd), float64(k2.xEnd)))
	} else if k1.xEnd >= k2.xStart && k1.xEnd <= k2.xEnd {
		k.xStart = k2.xStart
		k.xEnd = int(math.Min(float64(k1.xEnd), float64(k2.xEnd)))
	} else {
		return kuboid{}, false
	}

	if k1.yStart >= k2.yStart && k1.yStart <= k2.yEnd {
		k.yStart = k1.yStart
		k.yEnd = int(math.Min(float64(k1.yEnd), float64(k2.yEnd)))
	} else if k1.yEnd >= k2.yStart && k1.yEnd <= k2.yEnd {
		k.yStart = k2.yStart
		k.yEnd = int(math.Min(float64(k1.yEnd), float64(k2.yEnd)))
	} else {
		return kuboid{}, false
	}

	if k1.zStart >= k2.zStart && k1.zStart <= k2.zEnd {
		k.zStart = k1.zStart
		k.zEnd = int(math.Min(float64(k1.zEnd), float64(k2.zEnd)))
	} else if k1.zEnd >= k2.zStart && k1.zEnd <= k2.zEnd {
		k.zStart = k2.zStart
		k.zEnd = int(math.Min(float64(k1.zEnd), float64(k2.zEnd)))
	} else {
		return kuboid{}, false
	}

	return k, true
}

func reboot2(instructions []instruction) int {
	var incExcs []incExc

	for _, instruction := range instructions {
		for _, ie := range incExcs {
			intersectKube, ok := intersect(instruction.kube, ie.kube)
			if !ok {
				continue
			}
			if instruction.on {
				if ie.include {
					incExcs = append(incExcs, incExc{kube: intersectKube, include: false})
					continue
				}
				incExcs = append(incExcs, incExc{kube: intersectKube, include: true})
				continue
			}
			if ie.include {
				incExcs = append(incExcs, incExc{kube: intersectKube, include: false})
				continue
			}
			incExcs = append(incExcs, incExc{kube: intersectKube, include: false})
		}
		if instruction.on {
			incExcs = append(incExcs, incExc{kube: instruction.kube, include: true})
		}
	}

	onReactors := 0
	for _, ie := range incExcs {
		k := ie.kube
		reactors := (k.xEnd - k.xStart) * (k.yEnd - k.yStart) * (k.zEnd - k.zStart)
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

	fmt.Printf("Part 1: %d\n", reboot1(instructions))
	fmt.Printf("Part 2: %d\n", reboot2(instructions))
}
