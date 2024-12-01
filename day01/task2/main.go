package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countOccurrences(slice []int, target int) int {
	count := 0
	for _, num := range slice {
		if num == target {
			count++
		}
	}
	return count
}

func main() {
	var leftNumbers, rightNumbers []int

	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)
		if len(numbers) == 2 {
			left, err1 := strconv.Atoi(numbers[0])
			right, err2 := strconv.Atoi(numbers[1])
			if err1 == nil && err2 == nil {
				leftNumbers = append(leftNumbers, left)
				rightNumbers = append(rightNumbers, right)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	var similarityScore int

	for _, number := range leftNumbers {
		occurences := countOccurrences(rightNumbers, number)
		similarityScore += number * occurences
	}

	fmt.Println(similarityScore)
}
