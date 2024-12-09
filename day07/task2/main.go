package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	operators := []rune{'*', '+', '|'}
	totalCombinations := 1
	for i := 0; i < n-1; i++ {
		totalCombinations *= len(operators) // 3^n-1
	}

	combinations := make([][]rune, totalCombinations)
	for i := 0; i < totalCombinations; i++ {
		combination := make([]rune, n-1)
		num := i
		for j := 0; j < n-1; j++ {
			combination[j] = operators[num%len(operators)] // Select operator based on remainder
			num /= len(operators)                          // Update number for next position
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
			if combination[i-1] == '|' {
				stringResult := strconv.Itoa(result)
				stringResult = stringResult + equation.Numbers[i]
				result, _ = strconv.Atoi(stringResult)
			} else {
				number, _ := strconv.Atoi(equation.Numbers[i])
				if combination[i-1] == '*' {
					result = result * number
				} else if combination[i-1] == '+' {
					result = result + number
				}
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

	start := time.Now()

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
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
