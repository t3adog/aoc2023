package day01

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadLines("test", "day01_part01")
	result := PartOne(input)
	assert.Equal(t, 142, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day01_part02")
	result := PartTwo(input)
	assert.Equal(t, 281, result)
}

func Test_parseDigits(t *testing.T) {
	assert.Equal(t, 29, parseDigit("two1nine"))
	assert.Equal(t, 83, parseDigit("eightwothree"))
	assert.Equal(t, 13, parseDigit("abcone2threexyz"))
	assert.Equal(t, 15, parseDigit("one5foursevenfivezpgvjmdhl"))
	assert.Equal(t, 29, parseDigit("two1nine"))
	assert.Equal(t, 83, parseDigit("eightwothree"))
	assert.Equal(t, 13, parseDigit("abcone2threexyz"))
	assert.Equal(t, 24, parseDigit("xtwone3four"))
	assert.Equal(t, 42, parseDigit("4nineeightseven2"))
	assert.Equal(t, 14, parseDigit("zoneight234"))
	assert.Equal(t, 76, parseDigit("7pqrstsixteen"))
	assert.Equal(t, 11, parseDigit("kc1"))
	assert.Equal(t, 83, parseDigit("eighthree"))
	assert.Equal(t, 79, parseDigit("sevenine"))
	assert.Equal(t, 66, parseDigit("6seven96"))
}
