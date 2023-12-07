package day07

import (
	"testing"

	"github.com/kptlr/aoc2023/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := utils.ReadLines("test", "day07")
	result := PartOne(input)
	assert.Equal(t, 6440, result)
}

func TestPartTwo(t *testing.T) {
	input := utils.ReadLines("test", "day07")
	result := PartTwo(input)
	assert.Equal(t, 1, result)
}

func TestParseCombination(t *testing.T) {
	assert.Equal(t, "highCard", parseCombination("23456"))
	assert.Equal(t, "onePair", parseCombination("A23A4"))
	assert.Equal(t, "twoPair", parseCombination("23432"))
	assert.Equal(t, "threeOfKind", parseCombination("TTT98"))
	assert.Equal(t, "fullHouse", parseCombination("23332"))
	assert.Equal(t, "fourOfKind", parseCombination("AA8AA"))
	assert.Equal(t, "fiveOfKind", parseCombination("AAAAA"))
}

func TestFirstHandLessSecondHand(t *testing.T) {
	assert.Equal(t, false, firstHandLessSecondHand(Hand{combination: "fourOfKind", labels: "33332"}, Hand{combination: "fourOfKind", labels: "2AAAA"}))
	assert.Equal(t, false, firstHandLessSecondHand(Hand{combination: "fullHouse", labels: "77888"}, Hand{combination: "fullHouse", labels: "77788"}))
}
