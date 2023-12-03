package day02

import (
	"strconv"
	"strings"
)

type Color string

type ColorSets struct {
	colors map[Color]int
}

type Game struct {
	id   int
	sets []ColorSets
}

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

var colorsMap = map[string]Color{
	"red":   Red,
	"green": Green,
	"blue":  Blue,
}

var colorCapacity = map[Color]int{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func PartOne(input []string) (result int) {
	games := make([]Game, 0)
	for _, v := range input {
		games = append(games, parseGame(v))
	}

	for _, game := range games {
		correct := true
		for _, cs := range game.sets {
			for color, count := range cs.colors {
				if colorCapacity[color] < count {
					correct = false
					break
				}
			}
		}
		if correct {
			result += game.id
		}
	}
	return result
}

func PartTwo(input []string) (result int) {
	games := make([]Game, 0)
	for _, v := range input {
		games = append(games, parseGame(v))
	}
	for _, game := range games {
		red := 0
		green := 0
		blue := 0
		for _, colorSet := range game.sets {
			if colorSet.colors[Red] > red {
				red = colorSet.colors[Red]
			}
			if colorSet.colors[Green] > green {
				green = colorSet.colors[Green]
			}
			if colorSet.colors[Blue] > blue {
				blue = colorSet.colors[Blue]
			}
		}
		result += red * green * blue
	}
	return
}
func parseGame(input string) Game {
	gameId, _ := strconv.Atoi(strings.Split(strings.Split(input, ":")[0], " ")[1])
	colorSets := make([]ColorSets, 0)
	for _, set := range strings.Split(strings.Split(input, ":")[1], ";") {
		colorsSet := make(map[Color]int)
		for _, colors := range strings.Split(set, ",") {

			color := colorsMap[strings.Split(strings.Trim(colors, " "), " ")[1]]
			count, _ := strconv.Atoi(strings.Split(strings.Trim(colors, " "), " ")[0])
			colorsSet[color] = count
			colorSets = append(colorSets, ColorSets{colorsSet})
		}
	}
	return Game{gameId, colorSets}
}
