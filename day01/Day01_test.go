package day01

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
)

var input = utils.ReadLines("test", "day01")

func TestPartOne(t *testing.T) {
	result := PartOne(input)
	if result != 142 {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	result := PartTwo(input)
	if result != 1 {
		t.Fail()
	}
}
