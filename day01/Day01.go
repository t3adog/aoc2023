package day01

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var DigitsRegexp = regexp.MustCompile("[0-9]+")

var DigitsTemplate = map[string]string{
	"one":   "1",
	"1":     "1",
	"two":   "2",
	"2":     "2",
	"three": "3",
	"3":     "3",
	"four":  "4",
	"4":     "4",
	"five":  "5",
	"5":     "5",
	"six":   "6",
	"6":     "6",
	"seven": "7",
	"7":     "7",
	"eight": "8",
	"8":     "8",
	"nine":  "9",
	"9":     "9",
}

type Digit struct {
	Index int
	Digit string
}

func PartOne(input []string) (result int) {
	for _, value := range input {
		strDigits := []rune(strings.Join(DigitsRegexp.FindAllString(value, -1), ""))
		calibrationValue, _ := strconv.Atoi(string(strDigits[0]) + string(strDigits[len(strDigits)-1]))
		result += calibrationValue
	}
	return
}

func PartTwo(input []string) (result int) {
	for _, v := range input {
		result += parseDigit(v)
	}
	return
}

func parseDigit(input string) (calibrationValue int) {

	var foundDigits = []Digit{}
	for k := range DigitsTemplate {
		foundDigits = findDigitsBySubstring(k, input, foundDigits)
	}

	sort.Slice(foundDigits, func(i, j int) bool {
		return foundDigits[i].Index < foundDigits[j].Index
	})

	calibrationValue, _ = strconv.Atoi(foundDigits[0].Digit + foundDigits[len(foundDigits)-1].Digit)

	return
}

func findDigitsBySubstring(substring string, input string, result []Digit) []Digit {
	index := strings.Index(input, substring)
	if index > -1 {
		result = append(result, Digit{index, DigitsTemplate[substring]})
		return findDigitsBySubstring(substring, strings.Replace(input, substring, generatePlug(len(substring)), 1), result)
	} else {
		return result
	}
}

// Что бы не терять индекс
func generatePlug(len int) (result string) {
	for i := 0; i < len; i++ {
		result += "X"
	}
	return
}
