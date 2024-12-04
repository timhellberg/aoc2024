package main

import (
	"bufio"
	"fmt"
	"os"
)

func createGrid(file *os.File) [][]rune {
	var grid [][]rune

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return grid
}

func getWordMatchesInAllDirections(grid [][]rune, startX, startY int, word string) int {

	directions := [8][2]int{
		{1, 0},   // Right
		{-1, 0},  // Left
		{0, 1},   // Down
		{0, -1},  // Up
		{1, 1},   // Diagonal Down-Right
		{-1, 1},  // Diagonal Down-Left
		{1, -1},  // Diagonal Up-Right
		{-1, -1}, // Diagonal Up-Left
	}

	wordLength := len(word)
	rows := len(grid)
	cols := len(grid[0])

	var matchCounter int

	for _, dir := range directions {
		dx, dy := dir[0], dir[1]

		matches := true
		for i := 0; i < wordLength; i++ {
			newX := startX + i*dx
			newY := startY + i*dy

			if newX < 0 || newX >= cols || newY < 0 || newY >= rows {
				matches = false
				break
			}

			if grid[newY][newX] != rune(word[i]) {
				matches = false
				break
			}
		}

		if matches {
			matchCounter++
		}
	}

	return matchCounter
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := createGrid(file)
	word := "XMAS"
	var wordCounter int

	for y, row := range grid {
		for x, char := range row {
			if char == 'X' {
				wordCounter += getWordMatchesInAllDirections(grid, x, y, word)
			}
		}
	}

	fmt.Println(wordCounter)
}
