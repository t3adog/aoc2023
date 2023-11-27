package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

var root = "resources"

func ReadLines(env string, dayNumber string) ([]string, error) {
	filepath := filepath.Join(root, env, dayNumber+".txt")
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
