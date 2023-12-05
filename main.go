package main

import (
	"fmt"

	"github.com/kptlr/aoc2023/day01"
	"github.com/kptlr/aoc2023/day02"
	"github.com/kptlr/aoc2023/day03"
	"github.com/kptlr/aoc2023/day04"
	"github.com/kptlr/aoc2023/day05"
	"github.com/kptlr/aoc2023/day06"
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
	fmt.Println("Day 03, Part01: ", day03.PartOne(utils.ReadLines(env, "day03")))
	fmt.Println("Day 03, Part02: ", day03.PartTwo(utils.ReadLines(env, "day03")))

	// Day04
	fmt.Println("Day 04, Part01: ", day04.PartOne(utils.ReadLines(env, "day04")))
	fmt.Println("Day 04, Part02: ", day04.PartTwo(utils.ReadLines(env, "day04")))

	// Day05
	fmt.Println("Day 05, Part01: ", day05.PartOne(utils.ReadLines(env, "day05")))
	fmt.Println("Day 05, Part02: ", day05.PartTwo(utils.ReadLines(env, "day05")))

	// Day06
	fmt.Println("Day 06, Part01: ", day06.PartOne(utils.ReadLines(env, "day06")))
	fmt.Println("Day 06, Part02: ", day06.PartTwo(utils.ReadLines(env, "day06")))

	// ...

}
