package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) == 3 {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				sum += x * y
			}
		}
	}

	fmt.Println(sum)
}
