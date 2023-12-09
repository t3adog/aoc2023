package day08

import (
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func PartOne(input []string) (result int) {
	instructions := []rune(input[0])
	nodes := parseNodes(input[2:])

	currentKey := "AAA"
	currentNode := nodes[currentKey]
	i := 0
	for i < len(instructions) {
		result++
		if string(instructions[i]) == "L" {
			currentKey = currentNode.Left
		} else {
			currentKey = currentNode.Right
		}
		if currentKey == "ZZZ" {
			break
		}

		currentNode = nodes[currentKey]
		if i == len(instructions)-1 {
			i = 0
		} else {
			i++
		}
	}
	return
}

func PartTwo(input []string) (result int) {
	instructions := []rune(input[0])
	nodes := parseNodes(input[2:])
	currentKeys := make([]string, 0)

	for key, _ := range nodes {
		if strings.HasSuffix(key, "A") {
			currentKeys = append(currentKeys, key)
		}
	}
	loops := make([]int, 0)
	for _, k := range currentKeys {
		currentKey := k
		currentNode := nodes[currentKey]
		i := 0
		count := 0
		for i < len(instructions) {
			count++
			if string(instructions[i]) == "L" {
				currentKey = currentNode.Left
			} else {
				currentKey = currentNode.Right
			}
			if strings.HasSuffix(currentKey, "Z") {
				loops = append(loops, count)
				break
			}

			currentNode = nodes[currentKey]
			if i == len(instructions)-1 {
				i = 0
			} else {
				i++
			}
		}
	}
	return lcm(loops)
}

// https://ru.wikipedia.org/wiki/Наименьшее_общее_кратное
func lcm(nums []int) int {
	result := nums[0] * nums[1] / gcd(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		result = lcm([]int{result, nums[i]})
	}

	return result
}

// https://en.wikipedia.org/wiki/Greatest_common_divisor
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func parseNodes(input []string) map[string]Node {
	result := make(map[string]Node)
	for _, v := range input {
		key, node := parseNode(v)
		result[key] = node
	}
	return result
}

func parseNode(input string) (key string, node Node) {
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")
	key = strings.Trim(strings.Split(input, "=")[0], " ")

	nodes := strings.Split(input, "=")[1]
	node.Left = strings.Trim(strings.Split(nodes, ",")[0], " ")
	node.Right = strings.Trim(strings.Split(nodes, ",")[1], " ")

	return
}
