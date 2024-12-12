package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createList(file *os.File) []string {
	var list []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		list = strings.Split(line, "")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return list
}

func constructBlocks(diskMap []string) []string {
	var result []string

	var id int
	for index, char := range diskMap {
		size, _ := strconv.Atoi(char)
		list := make([]string, size, size)
		if index%2 == 0 {
			for i := 0; i < size; i++ {
				list[i] = strconv.Itoa(id)
			}
			id++
		} else {
			for i := 0; i < size; i++ {
				list[i] = "."
			}
		}

		result = append(result, list...)
	}

	return result
}

func main() {
	file, err := os.Open("../testinput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	diskMap := createList(file)
	blocks := constructBlocks(diskMap)

	for bc := len(blocks) - 1; bc > 0; bc-- {
		for fw := 0; fw < bc; fw++ {
			if blocks[fw] == "." {
				blocks[fw] = blocks[bc]
				blocks[bc] = "."
			}
		}
	}

	var sum int
	for index, char := range blocks {
		if char == "." {
			break
		}

		number, _ := strconv.Atoi(char)
		sum += number * index
	}

	fmt.Println(sum)
}
