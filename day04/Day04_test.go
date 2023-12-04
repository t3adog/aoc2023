package day04

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadLines("test", "day04")
	result := PartOne(input)
	assert.Equal(t, 13, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day04")
	result := PartTwo(input)
	assert.Equal(t, 30, result)
}
