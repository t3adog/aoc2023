package day06

import (
	"strconv"

	"github.com/kptlr/aoc2023/utils"
)

type Race struct {
	time int
	dist int
}

func PartOne(input []string) (result int) {
	races := parseRaces(input)
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
	return 1
}

func parseRaces(input []string) (races []Race) {
	timesStr := utils.DigitsRegexp.FindAllString(input[0], -1)
	distsStr := utils.DigitsRegexp.FindAllString(input[1], -1)

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
