package day04

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var DigitsRegexp = regexp.MustCompile("[0-9]+")

type Card struct {
	id         int
	winNumbers []int
	numbers    []int
	scores     int
}

func PartOne(input []string) (result int) {

	cards := make([]Card, len(input))

	for _, line := range input {
		cards = append(cards, parseCard(line))
	}

	for _, card := range cards {
		points := 0
		for _, number := range card.numbers {
			if slices.Contains(card.winNumbers, number) {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			}
		}
		result += points
	}

	return
}

func PartTwo(input []string) (result int) {
	cards := make([]Card, len(input))
	for idx, line := range input {
		cards[idx] = parseCard(line)
	}

	result = calculateScore(cards)
	return
}

func calculateScore(cards []Card) (result int) {
	numOfCards := make([]int, len(cards))

	for i := range numOfCards {
		numOfCards[i] = 1
	}

	for i, card := range cards {
		copies := numOfCards[i]

		if card.scores == 0 {
			continue
		}

		for score := 1; score <= card.scores; score++ {
			numOfCards[i+score] += copies
		}
	}

	for _, numOfCard := range numOfCards {
		result += numOfCard
	}

	return
}

func parseCard(input string) (card Card) {
	cardId, _ := strconv.Atoi(strings.Split(strings.Split(input, ":")[0], " ")[1])
	card.id = cardId
	cards := strings.Split(strings.Split(input, ":")[1], "|")
	card.winNumbers = parseNumbers(strings.Trim(cards[0], " "))
	card.numbers = parseNumbers(strings.Trim(cards[1], " "))
	card.scores = calcScores(card)

	return
}

func calcScores(card Card) int {
	points := 0
	for _, num := range card.numbers {
		if slices.Contains(card.winNumbers, num) {
			points++
		}
	}
	return points
}

func parseNumbers(input string) (numbers []int) {
	strNums := DigitsRegexp.FindAllString(input, -1)
	for _, numStr := range strNums {
		number, _ := strconv.Atoi(numStr)
		numbers = append(numbers, number)
	}
	return
}
