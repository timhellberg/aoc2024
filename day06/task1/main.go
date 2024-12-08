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

func getPathGrid(startGrid [][]rune, startX, startY int, startDirection rune) [][]rune {

	directionMap := map[rune][]int{
		'>': {1, 0},  // Right
		'<': {-1, 0}, // Left
		'V': {0, 1},  // Down
		'^': {0, -1}, // Up
	}

	walkedOut := false
	currentDirection := startDirection
	currentX := startX
	currentY := startY
	pathGrid := startGrid
	rows := len(pathGrid)
	cols := len(pathGrid[0])

	for walkedOut == false {
		pathGrid[currentY][currentX] = 'X'
		newX := currentX + directionMap[currentDirection][0]
		newY := currentY + directionMap[currentDirection][1]

		if newX < 0 || newX >= cols || newY < 0 || newY >= rows {
			return pathGrid
		} else if pathGrid[newY][newX] == '#' {
			currentDirection = getNewDirection(currentDirection)
		} else {
			currentX, currentY = newX, newY
		}
	}

	return pathGrid
}

func getNewDirection(direction rune) rune {
	switch direction {
	case '>':
		return 'V'
	case 'V':
		return '<'
	case '<':
		return '^'
	case '^':
		return '>'
	}
	return direction
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := createGrid(file)
	pathGrid := grid

	for y, row := range grid {
		for x, char := range row {
			if char == '^' || char == '<' || char == '>' || char == 'V' {
				pathGrid = getPathGrid(pathGrid, x, y, char)

			}
		}
	}

	var sumX int
	for _, row := range pathGrid {
		for _, char := range row {
			if char == 'X' {
				sumX++
			}
		}
	}

	fmt.Println(sumX)
}
