// This program implements the solution for
// https://adventofcode.com/2022/day/19.
package main

import (
	"fmt"
	"math"
	"strings"

	_ "embed"
)

//go:embed day19_input.txt
var input string

type blueprint struct {
	id                     int
	oreRobotOreCost        int
	clayRobotOreCost       int
	obsidianRobotOreCost   int
	obsidianRobotClayCost  int
	geodeRobotOreCost      int
	geodeRobotObsidianCost int
}

// parse parses a string like
// "Blueprint 1: Each ore robot costs 2 ore. Each clay robot costs 2 ore. Each obsidian robot costs 2 ore and 17 clay. Each geode robot costs 2 ore and 10 obsidian."
// into a blueprint struct.
func parse(blueprintStr string) blueprint {
	b := blueprint{}
	_, _ = fmt.Sscanf(
		blueprintStr,
		"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
		&b.id,
		&b.oreRobotOreCost,
		&b.clayRobotOreCost,
		&b.obsidianRobotOreCost,
		&b.obsidianRobotClayCost,
		&b.geodeRobotOreCost,
		&b.geodeRobotObsidianCost,
	)
	return b
}

func main() {
	sumQualities := 0
	for i, blueprintStr := range strings.Split(input, "\n") {
		maxGeode := 0
		dfs(24, resources{oreRobot: 1}, parse(blueprintStr), &maxGeode)
		sumQualities += maxGeode * (i + 1)
	}
	fmt.Printf("Part 1: %d\n", sumQualities)

	productGeodes := 1
	for i, blueprintStr := range strings.Split(input, "\n") {
		if i == 3 {
			break
		}
		maxGeode := 0
		dfs(32, resources{oreRobot: 1}, parse(blueprintStr), &maxGeode)
		productGeodes *= maxGeode
	}
	fmt.Printf("Part 2: %d\n", productGeodes)
}

type resources struct {
	ore, clay, obsidian, geode, oreRobot, clayRobot, obsidianRobot, geodeRobot int
}

// update simulates what happens at each minute to the resources. First you
// spend the specified minerals based on what robot you've decided to build.
// Then you get minerals that are gathered by your existing robots. Then at the
// end of the minute, the robot you've built this turn finish building and you
// get an additional count of that robot.
func (r resources) update(dOre, dClay, dObsidian int, robot string) resources {
	// Spend minerals to build a robot.
	r.ore += dOre
	r.clay += dClay
	r.obsidian += dObsidian

	// Existing robots gather minerals.
	r.ore += r.oreRobot
	r.clay += r.clayRobot
	r.obsidian += r.obsidianRobot
	r.geode += r.geodeRobot

	// You get more of the robot you built this turn.
	switch robot {
	case "ore":
		r.oreRobot++
	case "clay":
		r.clayRobot++
	case "obsidian":
		r.obsidianRobot++
	case "geode":
		r.geodeRobot++
	}

	return r
}

// dfs searches the decision graph of actions we can take per minute until there
// is no time left. It uses lots of pruning to allow this to finish in a
// reasonable time and prevent the DFS branching from exploding to be too large.
func dfs(timeLeft int, res resources, bp blueprint, maxGeode *int) {
	if timeLeft == 0 {
		if res.geode > *maxGeode {
			*maxGeode = res.geode
		}
		return
	}

	// We use a heuristic to help prune and decide if we should bother going down
	// this DFS branch/path. We can calculate the optimum geode we can get from
	// this state, which is if we were to build a geode robot for each minute of
	// the remaining time left plus our current geode and how much geode our
	// current geode robots can make in the time left. If this optimum geode is
	// less than our current best/max geode, there is no point exploring this
	// path.
	maxPossibleGeode := res.geode + res.geodeRobot*timeLeft
	for g := timeLeft; g > 0; g-- {
		maxPossibleGeode += g
	}
	if maxPossibleGeode < *maxGeode {
		return
	}

	// If we can build a geode robot, always build one, there's no reason not to.
	// And then there is no need to explore the other options for this turn.
	if res.ore >= bp.geodeRobotOreCost && res.obsidian >= bp.geodeRobotObsidianCost {
		dfs(timeLeft-1, res.update(-bp.geodeRobotOreCost, 0, -bp.geodeRobotObsidianCost, "geode"), bp, maxGeode)
		return
	}

	// If we can't build a geode robot, we need to explore building one of the
	// other robots, or saving. We can prune by only building up to a certain
	// number of each robot. We don't need more of each robot than the max cost of
	// each resource across the robots. We also can obviously only build a robot
	// if we have enough minerals for it. Lastly, we only want to save if we want
	// to build a type of robot (i.e., we haven't reached the max for that robot
	// type yet), but we can't afford it this turn - then we explore saving.
	// Because if we can afford to build each type of robot we want (haven't
	// reached the max yet), then there is no point in saving.
	save := false
	if res.obsidianRobot < bp.geodeRobotObsidianCost {
		if !(res.ore >= bp.obsidianRobotOreCost && res.clay >= bp.obsidianRobotClayCost) {
			if res.clayRobot > 0 {
				save = true
			}
		} else {
			dfs(timeLeft-1, res.update(-bp.obsidianRobotOreCost, -bp.obsidianRobotClayCost, 0, "obsidian"), bp, maxGeode)
		}
	}
	if res.clayRobot < bp.obsidianRobotClayCost {
		if !(res.ore >= bp.clayRobotOreCost) {
			save = true
		} else {
			dfs(timeLeft-1, res.update(-bp.clayRobotOreCost, 0, 0, "clay"), bp, maxGeode)
		}
	}
	if res.oreRobot < maxOreRobots(bp) {
		if !(res.ore >= bp.oreRobotOreCost) {
			save = true
		} else {
			dfs(timeLeft-1, res.update(-bp.oreRobotOreCost, 0, 0, "ore"), bp, maxGeode)
		}
	}

	if save {
		dfs(timeLeft-1, res.update(0, 0, 0, ""), bp, maxGeode)
	}
}

func maxOreRobots(bp blueprint) int {
	return int(math.Max(math.Max(float64(bp.clayRobotOreCost), float64(bp.obsidianRobotOreCost)), float64(bp.geodeRobotOreCost)))
}
