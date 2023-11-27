package main

import (
	"fmt"

	"github.com/kptlr/aoc2023/day01"
	"github.com/kptlr/aoc2023/utils"
)

var env = "prod"

func main() {
	// Day 1
	fmt.Print("Day 01, Part01: ")
	fmt.Println(day01.PartOne(utils.ReadLines(env, "day01")))
	fmt.Print("Day 01, Part02: ")
	fmt.Println(day01.PartTwo(utils.ReadLines(env, "day01")))
	//...
}
