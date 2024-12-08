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

func getPathCoordinates(startGrid [][]rune, startX, startY int, startDirection rune) map[Coordinate]struct{} {

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

	coordinateMap := make(map[Coordinate]struct{})

	for walkedOut == false {
		coordinate := Coordinate{X: currentX, Y: currentY}

		if _, exists := coordinateMap[coordinate]; !exists {
			coordinateMap[coordinate] = struct{}{}
		}

		newX := currentX + directionMap[currentDirection][0]
		newY := currentY + directionMap[currentDirection][1]

		if newX < 0 || newX >= cols || newY < 0 || newY >= rows {
			return coordinateMap
		} else if pathGrid[newY][newX] == '#' {
			currentDirection = getNewDirection(currentDirection)
		} else {
			currentX, currentY = newX, newY
		}
	}

	return coordinateMap
}

func checkIfLooping(startGrid [][]rune, startX, startY int, startDirection rune) bool {

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
	visitedMap := make(map[VisitedCoordinate]struct{})

	for walkedOut == false {

		visitedCoordinate := VisitedCoordinate{Coordinate: Coordinate{X: currentX, Y: currentY}, Direction: currentDirection}

		if _, exists := visitedMap[visitedCoordinate]; exists {
			return true
		}

		visitedMap[visitedCoordinate] = struct{}{}

		newX := currentX + directionMap[currentDirection][0]
		newY := currentY + directionMap[currentDirection][1]

		if newX < 0 || newX >= cols || newY < 0 || newY >= rows {
			return false
		} else if pathGrid[newY][newX] == '#' {
			currentDirection = getNewDirection(currentDirection)
		} else {
			currentX, currentY = newX, newY
		}
	}

	return false
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

type VisitedCoordinate struct {
	Coordinate Coordinate
	Direction  rune
}

type Coordinate struct {
	X int
	Y int
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

	var startCoordinate Coordinate
	var startDirection rune

	var coordinates = make(map[Coordinate]struct{})
	for y, row := range grid {
		for x, char := range row {
			if char == '^' || char == '<' || char == '>' || char == 'V' {
				startCoordinate = Coordinate{X: x, Y: y}
				startDirection = char
				coordinates = getPathCoordinates(pathGrid, x, y, char)

			}
		}
	}

	var sum int
	for coordinate := range coordinates {
		if coordinate == startCoordinate {
			continue
		}

		grid[coordinate.Y][coordinate.X] = '#'

		if checkIfLooping(grid, startCoordinate.X, startCoordinate.Y, startDirection) {
			sum += 1
		}

		grid[coordinate.Y][coordinate.X] = '.'

	}

	println(sum)
}
