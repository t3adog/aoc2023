package day05

import (
	"strconv"
	"strings"

	"github.com/kptlr/aoc2023/utils"
)

type Seed int
type AnotherSeed struct {
	from   int
	length int
}
type SeedMap struct {
	to     int
	from   int
	length int
}

func PartOne(input []string) (result int) {
	seeds := parseSeeds(input[0])
	seedToSoil := make([]SeedMap, 0)
	soilToFertilizer := make([]SeedMap, 0)
	fertilizerToWater := make([]SeedMap, 0)
	waterToLight := make([]SeedMap, 0)
	lightToTemp := make([]SeedMap, 0)
	temToHum := make([]SeedMap, 0)
	humToLocation := make([]SeedMap, 0)

	gap := 0
	for _, value := range input[2:] {
		if strings.Contains(value, "map") {
			continue
		}
		if value == "" {
			gap++
			continue
		}
		if gap == 0 {
			seedToSoil = append(seedToSoil, parseMap(value))
		}
		if gap == 1 {
			soilToFertilizer = append(soilToFertilizer, parseMap(value))
		}
		if gap == 2 {
			fertilizerToWater = append(fertilizerToWater, parseMap(value))
		}
		if gap == 3 {
			waterToLight = append(waterToLight, parseMap(value))
		}
		if gap == 4 {
			lightToTemp = append(lightToTemp, parseMap(value))
		}
		if gap == 5 {
			temToHum = append(temToHum, parseMap(value))
		}
		if gap == 6 {
			humToLocation = append(humToLocation, parseMap(value))
		}
	}

	locations := make([]int, 0)

	for _, seed := range seeds {
		location := getFromMap(int(seed), seedToSoil)
		location = getFromMap(location, soilToFertilizer)
		location = getFromMap(location, fertilizerToWater)
		location = getFromMap(location, waterToLight)
		location = getFromMap(location, lightToTemp)
		location = getFromMap(location, temToHum)
		location = getFromMap(location, humToLocation)
		locations = append(locations, location)
	}
	return findMin(locations)
}

func getFromMap(oldValue int, maps []SeedMap) int {
	for _, v := range maps {
		if oldValue >= v.from && oldValue <= v.from+(v.length-1) {
			if v.to > v.from {
				diff := v.to - v.from
				return oldValue + diff
			} else {
				diff := v.from - v.to
				return oldValue - diff
			}
		}
	}
	return oldValue
}

func findMin(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func parseSeeds(input string) (seeds []Seed) {
	numDigits := utils.DigitsRegexp.FindAllString(input, -1)
	for _, numStr := range numDigits {
		num, _ := strconv.Atoi(numStr)
		seeds = append(seeds, Seed(num))
	}
	return
}

func parseAnotherSeeds(input string) (seeds []AnotherSeed) {
	numDigits := utils.DigitsRegexp.FindAllString(input, -1)
	for i := 1; i < len(numDigits); i += 2 {
		seedFrom, _ := strconv.Atoi(numDigits[i-1])
		lenght, _ := strconv.Atoi(numDigits[i])
		seeds = append(seeds, AnotherSeed{from: seedFrom, length: lenght})
	}
	return
}

func parseMap(input string) (result SeedMap) {
	numsStr := utils.DigitsRegexp.FindAllString(input, -1)
	to, _ := strconv.Atoi(numsStr[0])
	from, _ := strconv.Atoi(numsStr[1])
	length, _ := strconv.Atoi(numsStr[2])

	result.to = to
	result.from = from
	result.length = length
	return result
}

func PartTwo(input []string) (result int) { // work slow, but work
	seeds := parseAnotherSeeds(input[0])
	seedToSoil := make([]SeedMap, 0)
	soilToFertilizer := make([]SeedMap, 0)
	fertilizerToWater := make([]SeedMap, 0)
	waterToLight := make([]SeedMap, 0)
	lightToTemp := make([]SeedMap, 0)
	temToHum := make([]SeedMap, 0)
	humToLocation := make([]SeedMap, 0)

	gap := 0
	for _, value := range input[2:] {
		if strings.Contains(value, "map") {
			continue
		}
		if value == "" {
			gap++
			continue
		}
		if gap == 0 {
			seedToSoil = append(seedToSoil, parseMap(value))
		}
		if gap == 1 {
			soilToFertilizer = append(soilToFertilizer, parseMap(value))
		}
		if gap == 2 {
			fertilizerToWater = append(fertilizerToWater, parseMap(value))
		}
		if gap == 3 {
			waterToLight = append(waterToLight, parseMap(value))
		}
		if gap == 4 {
			lightToTemp = append(lightToTemp, parseMap(value))
		}
		if gap == 5 {
			temToHum = append(temToHum, parseMap(value))
		}
		if gap == 6 {
			humToLocation = append(humToLocation, parseMap(value))
		}
	}

	min := -1
	for _, seed := range seeds {
		for i := 0; i <= seed.length; i++ {
			location := getFromMap(seed.from+i, seedToSoil)
			location = getFromMap(location, soilToFertilizer)
			location = getFromMap(location, fertilizerToWater)
			location = getFromMap(location, waterToLight)
			location = getFromMap(location, lightToTemp)
			location = getFromMap(location, temToHum)
			location = getFromMap(location, humToLocation)

			if min < 0 {
				min = location
			}
			if location < min {
				min = location
			}

		}
	}
	return min
}
