package day10

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne_01(t *testing.T) {
	input := utils.ReadLines("test", "day10_part01_01")
	result := PartOne(input)
	assert.Equal(t, 4, result)
}

func TestPartOne_02(t *testing.T) {
	input := utils.ReadLines("test", "day10_part01_02")
	result := PartOne(input)
	assert.Equal(t, 8, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day10_part01_01")
	result := PartTwo(input)
	assert.Equal(t, 1, result)
}
