// This program implements the solution for https://adventofcode.com/2020/day/4.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string

func validate(p passport) bool {
	byr, err := strconv.Atoi(p["byr"])
	if err != nil {
		return false
	}
	if !(byr >= 1920 && byr <= 2002) {
		return false
	}

	iyr, err := strconv.Atoi(p["iyr"])
	if err != nil {
		return false
	}
	if !(iyr >= 2010 && iyr <= 2020) {
		return false
	}

	eyr, err := strconv.Atoi(p["eyr"])
	if err != nil {
		return false
	}
	if !(eyr >= 2020 && eyr <= 2030) {
		return false
	}

	hgt := p["hgt"]
	switch {
	case strings.HasSuffix(hgt, "cm"):
		hgt = strings.TrimRight(hgt, "cm")
		h, err := strconv.Atoi(hgt)
		if err != nil {
			return false
		}
		if !(h >= 150 && h <= 193) {
			return false
		}
	case strings.HasSuffix(hgt, "in"):
		hgt = strings.TrimRight(hgt, "in")
		h, err := strconv.Atoi(hgt)
		if err != nil {
			return false
		}
		if !(h >= 59 && h <= 76) {
			return false
		}
	default:
		return false
	}

	var re = regexp.MustCompile(`^#[0-9|a-f]{6}$`)
	if !re.MatchString(p["hcl"]) {
		return false
	}

	if ecl := p["ecl"]; ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		return false
	}

	if len(p["pid"]) != 9 {
		return false
	}
	if _, err := strconv.Atoi(p["pid"]); err != nil {
		return false
	}

	return true
}

func main() {
	file, err := os.Open("day04_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valid1 := 0
	valid2 := 0
	p := passport{}
	scanner := bufio.NewScanner(file)
	// For this to work we need to cheat a bit by adding an extra newline at the
	// end of the input so we process the last passport.
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if _, ok := p["cid"]; len(p) == 7 && !ok || len(p) == 8 {
				valid1++
				if validate(p) {
					valid2++
				}
			}
			p = passport{}
			continue
		}

		fvPairs := strings.Split(line, " ")
		for _, fvPair := range fvPairs {
			parts := strings.Split(fvPair, ":")
			p[parts[0]] = parts[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d valid passports\n", valid1)
	fmt.Printf("Part 2: %d valid passports\n", valid2)
}
