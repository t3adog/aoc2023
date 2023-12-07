package day07

import (
	"slices"
	"sort"
	"strconv"
	"strings"
)

var combinationStrength = map[string]int{
	"highCard":    1,
	"onePair":     2,
	"twoPair":     3,
	"threeOfKind": 4,
	"fullHouse":   5,
	"fourOfKind":  6,
	"fiveOfKind":  7,
}

var cardStrength = map[string]int{
	"j": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type Hand struct {
	labels      string
	amount      int
	combination string
}

func PartOne(input []string) (result int) {
	hands := parseHands(input)

	sort.Slice(hands, func(i, j int) bool {
		return firstHandLessSecondHand(hands[i], hands[j])
	})

	return calcTotalWinnigs(hands)
}

func PartTwo(input []string) (result int) {
	hands := parseHandsWithJoker(input)

	sort.Slice(hands, func(i, j int) bool {
		return firstHandLessSecondHand(hands[i], hands[j])
	})

	return calcTotalWinnigs(hands)
}

func parseCombination(input string) string {
	combinationMap := map[rune]int{}
	for _, v := range input {
		value, ok := combinationMap[v]
		if !ok {
			combinationMap[v] = 1
		} else {
			combinationMap[v] = value + 1
		}
	}
	combinations := make([]int, 0)
	for _, v := range combinationMap {
		combinations = append(combinations, v)
	}
	slices.Sort(combinations)

	first := combinations[len(combinations)-1]

	if first == 5 {
		return "fiveOfKind"
	}

	if first == 4 {
		return "fourOfKind"
	}

	second := combinations[len(combinations)-2]

	if first == 3 && second == 2 {
		return "fullHouse"
	}
	if first == 3 {
		return "threeOfKind"
	}
	if first == 2 && second == 2 {
		return "twoPair"
	}
	if first == 2 {
		return "onePair"
	}

	return "highCard"
}

func parseCombinationWithJoker(input string) string {

	maxCombination := parseCombination(input)
	for _, v := range input {
		if string(v) == "j" {
			continue
		}
		newCombination := parseCombination(strings.ReplaceAll(input, "j", string(v)))
		if combinationStrength[maxCombination] < combinationStrength[newCombination] {
			maxCombination = newCombination
		}
	}
	return maxCombination
}

func parseHands(input []string) (hands []Hand) {
	for _, v := range input {
		hand := strings.Split(v, " ")[0]
		amount, _ := strconv.Atoi(strings.Split(v, " ")[1])
		hands = append(hands, Hand{labels: hand, amount: amount, combination: parseCombination(hand)})
	}
	return
}

func parseHandsWithJoker(input []string) (hands []Hand) {
	for _, v := range input {
		v = strings.ReplaceAll(v, "J", "j")
		hand := strings.Split(v, " ")[0]
		amount, _ := strconv.Atoi(strings.Split(v, " ")[1])
		hands = append(hands, Hand{labels: hand, amount: amount, combination: parseCombinationWithJoker(hand)})
	}
	return
}

func firstHandLessSecondHand(firstHand Hand, secondHand Hand) bool {

	if combinationStrength[firstHand.combination] != combinationStrength[secondHand.combination] {
		return combinationStrength[firstHand.combination] < combinationStrength[secondHand.combination]
	} else {
		for i, _ := range firstHand.labels {
			firstHandChar := string(firstHand.labels[i])
			secondHandChar := string(secondHand.labels[i])
			if cardStrength[firstHandChar] == cardStrength[secondHandChar] {
				continue
			}
			return cardStrength[firstHandChar] < cardStrength[secondHandChar]
		}
		return false
	}
}

func calcTotalWinnigs(hands []Hand) (result int) {
	for i, v := range hands {
		result = result + (v.amount * (i + 1))
	}
	return
}
