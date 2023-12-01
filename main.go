package main

import (
	"fmt"

	"github.com/kptlr/aoc2023/day01"
	"github.com/kptlr/aoc2023/day02"
	"github.com/kptlr/aoc2023/day03"
	"github.com/kptlr/aoc2023/utils"
)

var env = "prod"

func main() {

	// Day 01
	fmt.Println("Day 01, Part01: ", day01.PartOne(utils.ReadLines(env, "day01")))
	fmt.Println("Day 01, Part02: ", day01.PartTwo(utils.ReadLines(env, "day01")))

	// Day02
	fmt.Println("Day 02, Part01: ", day02.PartOne(utils.ReadLines(env, "day02")))
	fmt.Println("Day 02, Part02: ", day02.PartTwo(utils.ReadLines(env, "day02")))

	// Day03
	fmt.Println("Day 03, Part03: ", day03.PartOne(utils.ReadLines(env, "day03")))
	fmt.Println("Day 03, Part03: ", day03.PartTwo(utils.ReadLines(env, "day03")))

	// ...

}
