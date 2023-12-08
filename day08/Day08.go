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
			//fmt.Println("ХОБА")
			i = 0
		} else {
			i++
		}
		//fmt.Println(result)
	}
	return
}

func PartTwo(input []string) (result int) {
	return 1
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
