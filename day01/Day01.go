package day01

import (
	"regexp"
	"strconv"
	"strings"
)

var DigitsRegexp = regexp.MustCompile("[0-9]+")

func PartOne(input []string) (result int) {
	for _, value := range input {
		calibrationValue := parseCalibrationValue(value)
		result += calibrationValue
	}
	return
}

func PartTwo(input []string) int {
	return 1
}

func parseCalibrationValue(input string) (result int) {
	strDigits := []rune(strings.Join(DigitsRegexp.FindAllString(input, -1), ""))
	result, _ = strconv.Atoi(string(strDigits[0]) + string(strDigits[len(strDigits)-1]))
	return
}
