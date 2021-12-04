// This program implements the solution for https://adventofcode.com/2021/day/4.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type adjacent struct {
	nums   map[int]interface{}
	marked int
}

func newBoard() *board {
	var rows []*adjacent
	var cols []*adjacent
	for i := 0; i < 5; i++ {
		rows = append(rows, &adjacent{
			nums: map[int]interface{}{},
		})
		cols = append(cols, &adjacent{
			nums: map[int]interface{}{},
		})
	}
	return &board{
		rows: rows,
		cols: cols,
	}
}

type board struct {
	rows, cols []*adjacent
	sum        int
	finished   bool
}

func (b *board) mark(num int) bool {
	for _, row := range b.rows {
		if _, ok := row.nums[num]; ok {
			b.sum -= num
			row.marked++
			if row.marked == 5 {
				return true
			}
		}
	}
	for _, col := range b.cols {
		if _, ok := col.nums[num]; ok {
			col.marked++
			if col.marked == 5 {
				return true
			}
		}
	}

	return false
}

func main() {
	file, err := os.Open("day04_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nums := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	scanner.Scan()
	var boards []*board
	for {
		b := newBoard()
		for i := 0; i < 5; i++ {
			row := strings.Split(scanner.Text(), " ")
			j := 0
			for _, numStr := range row {
				if numStr == "" {
					continue
				}
				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatal(err)
				}
				b.rows[i].nums[num] = nil
				b.cols[j].nums[num] = nil
				b.sum += num
				j++
			}
			scanner.Scan()
		}
		boards = append(boards, b)
		if !scanner.Scan() {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var scores []int
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		for _, b := range boards {
			if b.finished {
				continue
			}
			if b.mark(num) {
				scores = append(scores, b.sum*num)
				b.finished = true
			}
		}
	}
	fmt.Printf("Part 1: %d\n", scores[0])
	fmt.Printf("Part 2: %d\n", scores[len(scores)-1])
}
