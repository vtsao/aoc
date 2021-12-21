// This program implements the solution for
// https://adventofcode.com/2021/day/21.
//
// curl -b "$(cat .session)" -o day21_input.txt https://adventofcode.com/2021/day/21/input
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type dice struct {
	next int
}

func (d *dice) roll() int {
	d.next++
	if d.next == 101 {
		d.next = 1
	}
	return d.next
}

func move(p int, d *dice) int {
	for i := 0; i < 3; i++ {
		p += d.roll()
	}
	return p % 10
}

func game1(p1, p2 int, d *dice, endGame int) (int, int, int) {
	p1Turn := true
	p1Score, p2Score := 0, 0
	diceRolls := 0
	for {
		if p1Turn {
			p1 = move(p1, d)
			p1Score += p1 + 1
			diceRolls += 3
			if p1Score >= endGame {
				break
			}
			p1Turn = false
			continue
		}
		p2 = move(p2, d)
		p2Score += p2 + 1
		diceRolls += 3
		if p2Score >= endGame {
			break
		}
		p1Turn = true
	}

	return p1Score, p2Score, diceRolls
}

type key struct {
	p1, p2, p1Score, p2Score int
	p1Turn                   bool
}

type wins struct {
	p1, p2 int
}

func game2(p1, p2, p1Score, p2Score, endGame int, p1Turn bool, cache map[key]wins) (int, int) {
	if p1Score >= endGame {
		return 1, 0
	}
	if p2Score >= endGame {
		return 0, 1
	}
	k := key{p1: p1, p2: p2, p1Score: p1Score, p2Score: p2Score, p1Turn: p1Turn}
	if wins, ok := cache[k]; ok {
		return wins.p1, wins.p2
	}

	p1Wins, p2Wins := 0, 0
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				roll := i + j + k
				if p1Turn {
					newP1 := p1 + roll
					newP1 %= 10
					newP1Score := p1Score + newP1 + 1
					p1WinsR, p2WinsR := game2(newP1, p2, newP1Score, p2Score, endGame, false, cache)
					p1Wins += p1WinsR
					p2Wins += p2WinsR
					continue
				}
				newP2 := p2 + roll
				newP2 %= 10
				newP2Score := p2Score + newP2 + 1
				p1WinsR, p2WinsR := game2(p1, newP2, p1Score, newP2Score, endGame, true, cache)
				p1Wins += p1WinsR
				p2Wins += p2WinsR
			}
		}
	}

	cache[k] = wins{p1: p1Wins, p2: p2Wins}
	return p1Wins, p2Wins
}

func parseInput() (int, int) {
	file, _ := os.Open("day21_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	p1, _ := strconv.Atoi(strings.TrimPrefix(scanner.Text(), "Player 1 starting position: "))
	scanner.Scan()
	p2, _ := strconv.Atoi(strings.TrimPrefix(scanner.Text(), "Player 2 starting position: "))

	return p1, p2
}

func main() {
	p1, p2 := parseInput()
	p1--
	p2--

	_, p2Score, diceRolls := game1(p1, p2, &dice{}, 1000)
	fmt.Printf("Part 1: %d\n", p2Score*diceRolls)

	p1Wins, p2Wins := game2(p1, p2, 0, 0, 21, true, map[key]wins{})
	fmt.Printf("Part 2: %d\n", int(math.Max(float64(p1Wins), float64(p2Wins))))
}
