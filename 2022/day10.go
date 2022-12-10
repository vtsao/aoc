// This program implements the solution for
// https://adventofcode.com/2022/day/10.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	c := cpu{x: 1}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), " ")
		inst := lineParts[0]
		val := 0
		if len(lineParts) > 1 {
			v, err := strconv.Atoi(lineParts[1])
			if err != nil {
				log.Fatal(err)
			}
			val = v
		}
		c.exec(inst, val)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", c.sumSignalStrs)
	fmt.Println("Part 2:")
	c.displayCRT()
	fmt.Printf("FBURHZCH\n")
}

type cpu struct {
	x             int
	cycle         int
	signalStrs    []int
	sumSignalStrs int
	crt           []string
}

func (c *cpu) exec(inst string, val int) {
	switch inst {
	case "noop":
		c.drawPixel()
		c.cycle++
		c.emitSignalStr()
	case "addx":
		for i := 0; i < 2; i++ {
			c.drawPixel()
			c.cycle++
			c.emitSignalStr()
		}
		c.x += val
	}
}

func (c *cpu) emitSignalStr() {
	switch {
	case c.cycle == 20:
		fallthrough
	case c.cycle >= 60 && (c.cycle-20)%40 == 0:
		signalStr := c.x * c.cycle
		c.signalStrs = append(c.signalStrs, signalStr)
		c.sumSignalStrs += signalStr
	}
}

func (c *cpu) drawPixel() {
	row := c.cycle / 40
	if row >= len(c.crt) {
		c.crt = append(c.crt, "")
	}

	pos := c.cycle % 40
	if pos == c.x || pos == c.x-1 || pos == c.x+1 {
		c.crt[row] += "#"
		return
	}
	c.crt[row] += "."
}

func (c *cpu) displayCRT() {
	for _, row := range c.crt {
		fmt.Println(row)
	}
}
