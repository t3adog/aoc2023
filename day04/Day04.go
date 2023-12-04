package day04

import (
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers map[int]int
	numbers        map[int]int
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
	cards := make([]Card, 1)
	for _, line := range input {
		cards = append(cards, parseCard(line))
	}

	_, total_score := calculateScore(cards)
	result = total_score
	return
}

// Не 2391905

func calculateScore(cards []Card) ([]int, int) {
	number_of_cards := make([]int, len(cards))

	for i := range number_of_cards {
		number_of_cards[i] = 1
	}

	for i, v := range cards {
		me_copies := number_of_cards[i]

		if v.scores == 0 {
			continue
		}

		for x := 1; x <= v.scores; x++ {
			number_of_cards[i+x] += me_copies
		}
	}

	total_score := 0

	for _, v := range number_of_cards {
		total_score += v
	}

	return number_of_cards, total_score
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
