package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kptlr/aoc2023/config"
)

func ReadLines(env string, day string) []string {
	filepath := filepath.Join(config.Resources, env, day+".txt")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Ошибка чтения файла: " + err.Error())
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func GeneratePlug(len int, symbols string) (result string) {
	for i := 0; i < len; i++ {
		result += symbols
	}
	return
}
