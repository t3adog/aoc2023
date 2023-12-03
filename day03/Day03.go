package day03

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/kptlr/aoc2023/utils"
)

var DigitsRegexp = regexp.MustCompile("\\.?[0-9]+\\.?")

type Point struct {
	y int
	x int
}

type Square struct {
	from Point
	to   Point
}

func PartOne(input []string) (result int) {
	for index, line := range input {
		digits := DigitsRegexp.FindAllString(line, -1)
		for _, digit := range digits {
			isDetail := isDetail(input, digit, index)
			input[index] = strings.Replace(input[index], digit, utils.GeneratePlug(len(digit), "."), 1) // Цифры могут повторяться. Так что глушим вхождения. Костыль, зато какой!
			if isDetail {
				d, _ := strconv.Atoi(strings.ReplaceAll(digit, ".", ""))
				result += d
			}
		}
	}
	return
}

func isDetail(input []string, digit string, index int) bool {
	searchSquare := parseSquare(input, digit, index)

	for y := searchSquare.from.y; y <= searchSquare.to.y; y++ {
		if len(strings.ReplaceAll(strings.ReplaceAll(input[y][searchSquare.from.x:searchSquare.to.x+1], digit, ""), ".", "")) > 0 {
			return true
		}
	}
	return false
}

func parseSquare(input []string, digit string, index int) (result Square) {
	maxX := len(input[index]) - 1
	maxY := len(input) - 1

	digitIndex := strings.Index(input[index], digit)

	// Ищем FROM
	if strings.HasPrefix(digit, ".") {
		result.from.x = digitIndex
	} else {
		if digitIndex-1 > 0 {
			result.from.x = digitIndex - 1
		} else {
			result.from.x = 0
		}
	}

	if index-1 >= 0 {
		result.from.y = index - 1
	} else {
		result.from.y = 0
	}

	// Ищем TO
	if strings.HasSuffix(digit, ".") {
		result.to.x = digitIndex + (len(digit) - 1)
	} else {
		if digitIndex+(len(digit)-1)+1 <= maxX {
			result.to.x = digitIndex + (len(digit) - 1) + 1
		} else {
			result.to.x = maxX
		}
	}

	if index+1 <= maxY {
		result.to.y = index + 1
	} else {
		result.to.y = maxY
	}
	return
}

func PartTwo(input []string) (result int) {
	return 1
}
