package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findCanContainBag(bagsToParents map[string]map[string]interface{}, bag string, canContain map[string]interface{}) {
	parents := bagsToParents[bag]
	for p := range parents {
		canContain[p] = nil
		findCanContainBag(bagsToParents, p, canContain)
	}
}

func findBagContains(bagsToContains map[string]map[string]int, bag string) int {
	contains := bagsToContains[bag]
	total := 0
	for contain, val := range contains {
		total += val
		total += val * findBagContains(bagsToContains, contain)
	}
	return total
}

func main() {
	file, err := os.Open("day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bagsToParents := map[string]map[string]interface{}{}
	bagsToContains := map[string]map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rule := strings.TrimSuffix(scanner.Text(), ".")
		ruleParts := strings.Split(rule, " contain ")
		bag := strings.TrimSuffix(ruleParts[0], " bags")
		if _, ok := bagsToContains[bag]; !ok {
			bagsToContains[bag] = map[string]int{}
		}
		if ruleParts[1] == "no other bags" {
			continue
		}

		contains := strings.Split(ruleParts[1], ",")
		for _, contain := range contains {
			contain := strings.TrimSpace(contain)
			containParts := strings.Split(contain, " ")
			containVal, err := strconv.Atoi(containParts[0])
			if err != nil {
				log.Fatal(err)
			}
			containBag := strings.Join(containParts[1:len(containParts)-1], " ")
			if _, ok := bagsToParents[containBag]; !ok {
				bagsToParents[containBag] = map[string]interface{}{}
			}
			bagsToParents[containBag][bag] = nil
			bagsToContains[bag][containBag] = containVal
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	canContainShinyGold := map[string]interface{}{}
	findCanContainBag(bagsToParents, "shiny gold", canContainShinyGold)
	fmt.Printf("Part 1: %d bags can eventually contain at least one \"shiny gold bag\"\n", len(canContainShinyGold))

	shinyGoldContains := findBagContains(bagsToContains, "shiny gold")
	fmt.Printf("Part 2: \"shiny gold bags\" contain %d other bags\n", shinyGoldContains)
}
