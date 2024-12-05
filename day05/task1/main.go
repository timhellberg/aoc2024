package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkIfCorrect(row []string, rulesmap map[string][]string) bool {
	isCorrect := true
	var listlen = len(row)
	for index, key := range row {
		numbersThatShouldBeAfter := rulesmap[key]
		numbersAfterIndex := row[index+1 : listlen]

		for i := 0; i < len(numbersAfterIndex); i++ {
			if !slices.Contains(numbersThatShouldBeAfter, numbersAfterIndex[i]) {
				isCorrect = false
			}
		}
	}

	return isCorrect
}

func parseInputData(filePath string) (map[string][]string, [][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var rulesMap = make(map[string][]string)
	var updates [][]string
	scanningRules := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			scanningRules = false
		}

		if scanningRules {
			rule := strings.Split(line, "|")
			rulesMap[rule[0]] = append(rulesMap[rule[0]], rule[1])
		} else if !scanningRules && line != "" {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error scanning file: %w", err)
	}

	return rulesMap, updates, nil
}

func main() {
	rulesMap, updates, err := parseInputData("../input.txt")
	if err != nil {
		fmt.Println("Failed to parse input data:", err)
		return
	}

	var sum int
	for _, row := range updates {
		isCorrect := checkIfCorrect(row, rulesMap)

		if isCorrect {
			number, _ := strconv.Atoi(row[(len(row) / 2)])
			sum += number
		}
	}

	fmt.Println(sum)
}
