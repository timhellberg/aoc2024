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

func numberSplitter(number string) []string {
	newNumber1 := number[0 : len(number)/2]
	newNumber2 := number[len(number)/2:]
	fixPotentialZeros, _ := strconv.Atoi(newNumber2)
	newNumber2 = strconv.Itoa(fixPotentialZeros)

	return []string{newNumber1, newNumber2}
}

var memo = make(map[Memory]int)

func Blink(number string, blink int) int {
	memory := Memory{Number: number, Blinks: blink}

	if memorisedList, exists := memo[memory]; exists {
		return memorisedList
	}

	if blink == 0 {
		memo[memory] = 1
		return memo[memory]
	}

	blink--
	if number == "0" {
		number = "1"
		memo[memory] = Blink(number, blink)
		return memo[memory]
	} else if len(number)%2 == 0 {
		sliceResult := numberSplitter(number)
		memo[memory] = Blink(sliceResult[0], blink) + Blink(sliceResult[1], blink)
		return memo[memory]
	} else {
		product, _ := strconv.Atoi(number)
		product *= 2024
		memo[memory] = Blink(strconv.Itoa(product), blink)
		return memo[memory]
	}

}

type Memory struct {
	Number string
	Blinks int
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
	blinks := 75

	var result int
	for _, number := range numbers {
		memory := Memory{Number: number, Blinks: blinks}
		memo[memory] = Blink(number, blinks)
		result += memo[memory]
	}

	elapsed := time.Since(start)
	fmt.Printf("Blinks: %v \n", blinks)
	fmt.Printf("Answer: %v \n", result)
	fmt.Printf("Time: %v \n", elapsed)

}
