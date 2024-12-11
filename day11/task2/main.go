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

func insertNumberSplitIntoArray(slice []string, index int) []string {
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

func numberSplitter(number string) []string {
	newNumber1 := number[0 : len(number)/2]
	newNumber2 := number[len(number)/2:]
	fixPotentialZeros, _ := strconv.Atoi(newNumber2)
	newNumber2 = strconv.Itoa(fixPotentialZeros)

	return []string{newNumber1, newNumber2}
}

var memo = make(map[Memory][]string)

func Blink(number string, blink int) []string {
	memory := Memory{Number: number, Blinks: blink}

	if memorisedList, exists := memo[memory]; exists {
		return memorisedList
	}

	if blink == 0 {
		memo[memory] = []string{number}
		return memo[memory]
	}

	blink--
	if number == "0" {
		number = "1"
		return Blink(number, blink)
	} else if len(number)%2 == 0 {
		sliceResult := numberSplitter(number)
		return append(Blink(sliceResult[0], blink), Blink(sliceResult[1], blink)...)
	} else {
		product, _ := strconv.Atoi(number)
		product *= 2024
		return Blink(strconv.Itoa(product), blink)
	}

}

type Memory struct {
	Number string
	Blinks int
}

func main() {
	file, err := os.Open("../testinput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	numbers := createSlice(file)

	start := time.Now()
	blinks := 25

	var result []string
	for _, number := range numbers {
		memory := Memory{Number: number, Blinks: blinks}
		memo[memory] = Blink(number, blinks)
		result = append(result, memo[memory]...)
	}

	elapsed := time.Since(start)
	fmt.Printf("Blinks: %v \n", blinks)
	fmt.Printf("Answer: %v \n", len(result))
	fmt.Printf("Time: %v \n", elapsed)

}
