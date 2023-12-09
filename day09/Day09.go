package day09

import (
	"strconv"
	"strings"
)

func PartOne(input []string) (result int) {

	for _, v := range input {
		numbers := parseNumbes(v)
		result += extrapolationsSum(numbers)
	}

	return
}

func PartTwo(input []string) (result int) {

	for _, v := range input {
		numbers := parseNumbes(v)
		result += extrapolationsBackwardsSum(numbers)
	}
	return
}

func parseNumbes(input string) (numbers []int) {
	for _, numberStr := range strings.Split(input, " ") {
		number, _ := strconv.Atoi(numberStr)
		numbers = append(numbers, number)
	}
	return
}

func extrapolationsSum(numbers []int) int {
	if len(numbers) <= 1 {
		return 0
	}
	for i := 0; i < len(numbers)-1; i++ {
		numbers[i] = numbers[i+1] - numbers[i]
	}
	return numbers[len(numbers)-1] + extrapolationsSum(numbers[:len(numbers)-1])
}

func extrapolationsBackwardsSum(numbers []int) int {
	if len(numbers) <= 1 {
		return 0
	}
	for i := len(numbers) - 1; i > 0; i-- {
		numbers[i] = numbers[i] - numbers[i-1]
	}
	return numbers[0] - extrapolationsBackwardsSum(numbers[1:])
}
