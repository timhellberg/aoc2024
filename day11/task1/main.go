package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func createSlice(file *os.File) []string {
	var slice []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		slice = strings.Fields(line)

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return slice
}

func numberSplitter(slice []string, index int) []string {
	number := slice[index]
	newNumber1 := number[0 : len(number)/2]
	newNumber2 := number[len(number)/2:]
	fixPotentialZeros, _ := strconv.Atoi(newNumber2)
	newNumber2 = strconv.Itoa(fixPotentialZeros)

	slice = append(slice, "")
	copy(slice[index+2:], slice[index+1:])
	slice[index] = newNumber1
	slice[index+1] = newNumber2
	return slice
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	numbers := createSlice(file)

	start := time.Now()
	var blinks int
	index := 0
	for blinks < 25 {

		for index < len(numbers) {
			number := numbers[index]

			if number == "0" {
				numbers[index] = "1"
			} else if len(number)%2 == 0 {
				numbers = numberSplitter(numbers, index)
				index++
			} else {
				product, _ := strconv.Atoi(number)
				product *= 2024
				numbers[index] = strconv.Itoa(product)
			}

			index++

		}

		blinks++
		index = 0

	}

	elapsed := time.Since(start)
	fmt.Printf("Blinks: %v \n", blinks)
	fmt.Printf("Answer: %v \n", len(numbers))
	fmt.Printf("Time: %v \n", elapsed)

}
