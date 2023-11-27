package main

import (
	"fmt"

	"github.com/kptlr/aoc2023/day01"
	"github.com/kptlr/aoc2023/utils"
)

func main() {
	// Day 01
	fmt.Println(day01.PartOne())
	fmt.Println(day01.PartTwo())

	lines, _ := utils.ReadLines("test", "day01")
	lines2, _ := utils.ReadLines("prod", "day01")

	fmt.Println(lines)
	fmt.Println(lines2)

}
