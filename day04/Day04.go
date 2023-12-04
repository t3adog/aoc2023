package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers map[int]int
	numbers        map[int]int
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
	cards := make(map[int]Card)
	queue := make([]Card, 0)
	for _, line := range input {
		card := parseCard(line)
		cards[card.id] = card
		queue = append(queue, card)
	}

	fmt.Println("Size: ", len(queue))
	processCards := process(cards, 0, queue)

	fmt.Println(len(processCards))

	result = len(processCards)
	return
}

// Не 2391905

func process(cards map[int]Card, index int, queue []Card) []Card {
	fmt.Println("Index: ", index, "Queue: ", len(queue))
	for x := index; x < len(queue); x++ {
		fmt.Println(len(queue))
		card := queue[x]
		points := 0
		for _, winningNumber := range card.winningNumbers {
			if card.numbers[winningNumber] == winningNumber {
				points++
			}
		}
		for id := 1; id <= points; id++ {
			queue = append(queue, cards[card.id+id])
		}
	}
	return queue
}

func parseCard(input string) (card Card) {
	cardId, _ := strconv.Atoi(strings.Split(strings.Split(input, ":")[0], " ")[1])
	card.id = cardId
	cards := strings.Split(strings.Split(input, ":")[1], "|")
	card.winningNumbers = parseNumbers(strings.Trim(cards[0], " "))
	card.numbers = parseNumbers(strings.Trim(cards[1], " "))

	return
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
