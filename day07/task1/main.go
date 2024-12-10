package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	TestValue string
	Numbers   []string
}

func createEquations(file *os.File) []Equation {
	scanner := bufio.NewScanner(file)
	var equations []Equation

	for scanner.Scan() {
		line := scanner.Text()

		splitString := strings.Split(line, ":")

		numbers := strings.Fields(splitString[1])

		equation := Equation{TestValue: splitString[0], Numbers: numbers}
		equations = append(equations, equation)

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return equations
}

func generateCombinations(n int) [][]rune {

	totalCombinations := 1 << (n - 1) // 2^(n-1)

	combinations := make([][]rune, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]rune, n-1)
		for j := 0; j < n-1; j++ {
			if (i>>j)&1 == 1 {
				combination[j] = '+' // Set '+' if bit is 1
			} else {
				combination[j] = '*' // Set '*' if bit is 0
			}
		}
		combinations[i] = combination
	}

	return combinations
}

func calcIfAnyCombinationIsValid(equation Equation, combinations [][]rune) bool {
	correctNumber, _ := strconv.Atoi(equation.TestValue)
	for _, combination := range combinations {
		result, _ := strconv.Atoi(equation.Numbers[0])

		for i := 1; i < len(equation.Numbers); i++ {
			number, _ := strconv.Atoi(equation.Numbers[i])
			if combination[i-1] == '*' {
				result = result * number
			} else if combination[i-1] == '+' {
				result = result + number
			}

			if result > correctNumber {
				break
			}
		}

		if result == correctNumber {
			return true
		}

	}

	return false
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	equations := createEquations(file)
	combinationMap := make(map[int][][]rune)

	var sum int
	for _, equation := range equations {
		lenNumbers := len(equation.Numbers)
		combinations, exists := combinationMap[lenNumbers]
		if !exists {
			combinations = generateCombinations(lenNumbers)
			combinationMap[lenNumbers] = combinations
		}

		if calcIfAnyCombinationIsValid(equation, combinations) {
			value, _ := strconv.Atoi(equation.TestValue)
			sum += value

		}
	}

	fmt.Println(sum)
}
