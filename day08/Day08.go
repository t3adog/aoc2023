package day08

import (
	"fmt"
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
	//fmt.Println("currentKey: ", currentKey, "currentNode:", currentNode)
	i := 0
	for i < len(instructions) {
		result++
		//fmt.Println("instructuion is: ", string(instructions[i]))
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

	fmt.Println(currentKeys)
	i := 0
	for i < len(instructions) {
		result++
		for keyIndex, currentKey := range currentKeys {
			currentNode := nodes[currentKey]
			if string(instructions[i]) == "L" {
				currentKey = currentNode.Left
				currentKeys[keyIndex] = currentKey
			} else {
				currentKey = currentNode.Right
				currentKeys[keyIndex] = currentKey
			}
			//currentNode = nodes[currentKey]
		}
		if isAllKeysFinished(currentKeys) {
			return result
		}
		if i == len(instructions)-1 {
			i = 0
		} else {
			i++
		}
	}
	return
}

func isAllKeysFinished(keys []string) bool {
	//fmt.Println("Current keys: ", keys)
	for _, v := range keys {
		if !strings.HasSuffix(v, "Z") {
			return false
		}
	}
	fmt.Println("OH LOL TRUE", keys)
	return true
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
