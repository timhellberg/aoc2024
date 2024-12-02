package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkIfSafe(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		diff := int(math.Abs(float64(numbers[i] - numbers[i+1])))
		if diff > 3 || diff < 1 {
			return false
		}
		if i > 0 && ((numbers[i-1] < numbers[i] && numbers[i] > numbers[i+1]) || (numbers[i-1] > numbers[i] && numbers[i] < numbers[i+1])) {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var safeReports int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		numberstring := strings.Fields(line)
		var numbers []int
		for _, numStr := range numberstring {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				return
			}
			numbers = append(numbers, num)
		}

		if checkIfSafe(numbers) {
			safeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	fmt.Println(safeReports)
}
