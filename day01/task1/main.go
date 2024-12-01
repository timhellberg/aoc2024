package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	var distance int
	for index, _ := range leftNumbers {
		distance += int(math.Abs(float64(leftNumbers[index] - rightNumbers[index])))
	}

	fmt.Println(distance)
}
