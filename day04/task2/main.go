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

func getWordMatchesInAllDirections(grid [][]rune, startX, startY int) bool {

	directions := [8][2]int{
		{1, 1},   // Diagonal Down-Right
		{-1, -1}, // Diagonal Up-Left
		{-1, 1},  // Diagonal Down-Left
		{1, -1},  // Diagonal Up-Right
	}

	downrightx, downrighty := startX+directions[0][0], startY+directions[0][1]
	upleftx, uplefty := startX+directions[1][0], startY+directions[1][1]

	downleftx, downlefty := startX+directions[2][0], startY+directions[2][1]
	uprightx, uprighty := startX+directions[3][0], startY+directions[3][1]

	matches := false
	firstDiagonalMatches := false
	secondDiagnoalMatches := false

	if grid[downrighty][downrightx] == 'M' && grid[uplefty][upleftx] == 'S' {
		firstDiagonalMatches = true
	} else if grid[downrighty][downrightx] == 'S' && grid[uplefty][upleftx] == 'M' {
		firstDiagonalMatches = true
	}

	if grid[downlefty][downleftx] == 'M' && grid[uprighty][uprightx] == 'S' {
		secondDiagnoalMatches = true
	} else if grid[downlefty][downleftx] == 'S' && grid[uprighty][uprightx] == 'M' {
		secondDiagnoalMatches = true
	}

	if firstDiagonalMatches && secondDiagnoalMatches {
		matches = true
	}

	return matches
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := createGrid(file)
	var wordCounter int

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			char := grid[y][x]
			if char == 'A' {
				if getWordMatchesInAllDirections(grid, x, y) {
					wordCounter++
				}
			}
		}
	}

	fmt.Println(wordCounter)
}
