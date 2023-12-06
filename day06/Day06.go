package day06

import (
	"strconv"
	"strings"

	"github.com/kptlr/aoc2023/utils"
)

type Race struct {
	time int
	dist int
}

func PartOne(input []string) (result int) {
	races := parseRacesV1(input)
	winOptions := make([][]int, 0)

	for _, v := range races {
		winOptions = append(winOptions, findWinOptions(v))
	}

	result = 1

	for _, v := range winOptions {
		result = result * len(v)
	}
	return result
}

func PartTwo(input []string) (result int) {
	races := parseRacesV2(input)
	winOptions := make([][]int, 0)

	for _, v := range races {
		winOptions = append(winOptions, findWinOptions(v))
	}

	result = 1

	for _, v := range winOptions {
		result = result * len(v)
	}
	return result
}

func parseRacesV1(input []string) (races []Race) {
	timesStr := utils.DigitsRegexp.FindAllString(input[0], -1)
	distsStr := utils.DigitsRegexp.FindAllString(input[1], -1)

	for i, _ := range timesStr {
		time, _ := strconv.Atoi(timesStr[i])
		dist, _ := strconv.Atoi(distsStr[i])
		races = append(races, Race{time, dist})
	}

	return
}

func parseRacesV2(input []string) (races []Race) {

	timesStr := utils.DigitsRegexp.FindAllString(strings.ReplaceAll(input[0], " ", ""), -1)
	distsStr := utils.DigitsRegexp.FindAllString(strings.ReplaceAll(input[1], " ", ""), -1)

	for i, _ := range timesStr {
		time, _ := strconv.Atoi(timesStr[i])
		dist, _ := strconv.Atoi(distsStr[i])
		races = append(races, Race{time, dist})
	}

	return
}

func findWinOptions(race Race) (winConfigurations []int) {
	for i := 1; i <= race.time; i++ {
		z := i * (race.time - i)
		if z > race.dist {
			winConfigurations = append(winConfigurations, i)
		}
	}
	return winConfigurations
}
