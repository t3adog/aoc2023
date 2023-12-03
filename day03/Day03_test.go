package day03

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadLines("test", "day03")
	result := PartOne(input)
	assert.Equal(t, 4361, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day03")
	result := PartTwo(input)
	assert.Equal(t, 467835, result)
}

func TestParseSquare(t *testing.T) {
	input := utils.ReadLines("test", "day03")
	sq := Square{}
	sq.from.x = 0
	sq.from.y = 0
	sq.to.x = 3
	sq.to.y = 1
	assert.Equal(t, sq, parseSquare(input, "467.", 0))

	sq1 := Square{}
	sq1.from.x = 4
	sq1.from.y = 8
	sq1.to.x = 8
	sq1.to.y = 9
	assert.Equal(t, sq1, parseSquare(input, ".598.", 9))

	sq3 := Square{}
	sq3.from.x = 1
	sq3.from.y = 1
	sq3.to.x = 4
	sq3.to.y = 3
	assert.Equal(t, sq3, parseSquare(input, ".35.", 2))
}
