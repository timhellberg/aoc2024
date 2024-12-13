package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"gonum.org/v1/gonum/optimize"
)

func createSlice(file *os.File) []string {
	var slice []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		slice = append(slice, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return slice
}

func extractNumbers(input string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func hasDecimalPart(f float64) bool {
	return math.Abs(f-math.Round(f)) > 1e-30
}

func main() {

	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	lines := createSlice(file)

	var sum int
	for i := 0; i < len(lines); i += 3 {
		if i+2 >= len(lines) {
			break
		}

		buttonALine := lines[i]
		buttonBLine := lines[i+1]
		prizeLine := lines[i+2]

		buttonANumbers := extractNumbers(buttonALine)
		buttonBNumbers := extractNumbers(buttonBLine)
		prizeNumbers := extractNumbers(prizeLine)

		if len(buttonANumbers) < 2 || len(buttonBNumbers) < 2 || len(prizeNumbers) < 2 {
			fmt.Println("Error: Invalid input format")
			continue
		}

		taskValue := 10000000000000
		XA, YA := buttonANumbers[0], buttonANumbers[1]
		XB, YB := buttonBNumbers[0], buttonBNumbers[1]
		XTarget, YTarget := prizeNumbers[0]+taskValue, prizeNumbers[1]+taskValue

		problem := optimize.Problem{
			Func: func(x []float64) float64 {
				a, b := x[0], x[1]
				leftSideX := a*float64(XA) + b*float64(XB)
				leftSideY := a*float64(YA) + b*float64(YB)
				rightSideX := float64(XTarget)
				rightSideY := float64(YTarget)
				return (leftSideX-rightSideX)*(leftSideX-rightSideX) + (leftSideY-rightSideY)*(leftSideY-rightSideY)
			},
		}

		init := []float64{1, 1}
		result, err := optimize.Minimize(problem, init, nil, nil)
		if err != nil {
			fmt.Println("Error optimizing:", err)
			return
		}

		if hasDecimalPart(result.X[0]) || hasDecimalPart(result.X[1]) {
			continue
		}

		sum += 3*int(result.X[0]) + int(result.X[1])
	}

	fmt.Println(sum)

}
