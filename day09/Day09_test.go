package day09

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadLines("test", "day09")
	result := PartOne(input)
	assert.Equal(t, 114, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day09")
	result := PartTwo(input)
	assert.Equal(t, 2, result)
}
