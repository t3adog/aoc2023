package day04

import (
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers map[int]int // whorst practies
	numbers        map[int]int // whorst practies
	scores         int
}

func PartOne(input []string) (result int) {

	cards := make([]Card, len(input))

	for _, line := range input {
		cards = append(cards, parseCard(line))
	}

	for _, card := range cards {
		points := 0
		for _, winningNumber := range card.winningNumbers {
			if card.numbers[winningNumber] == winningNumber {
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

		for x := 1; x <= card.scores; x++ {
			numOfCards[i+x] += copies
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
	card.winningNumbers = parseNumbers(strings.Trim(cards[0], " "))
	card.numbers = parseNumbers(strings.Trim(cards[1], " "))
	card.scores = calcScores(card)

	return
}

func calcScores(card Card) int {
	points := 0
	for _, num := range card.numbers {
		if card.winningNumbers[num] == num {
			points++
		}
	}
	return points
}

func parseNumbers(input string) (numbers map[int]int) {
	numbers = make(map[int]int)
	for _, num := range strings.Split(input, " ") {
		number, err := strconv.Atoi(string(num))
		if err == nil {
			numbers[number] = number
		}
	}
	return
}
