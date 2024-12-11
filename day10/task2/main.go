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

type Position struct {
	X int
	Y int
}

func (p1 Position) Add(p2 Position) Position {
	return Position{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

var globalMap map[Position]int

func FindTrail(position Position, value int, topographicMap map[Position]int, directions map[string]Position) int {
	if globalMap == nil {
		globalMap = make(map[Position]int)
	}

	if value, exists := globalMap[position]; exists {
		return value
	}

	if topographicMap[position] == 9 {
		globalMap[position] = 1
		return globalMap[position]
	}

	if _, exists := globalMap[position]; !exists {
		globalMap[position] = 0
	}

	for _, direction := range directions {
		newPosition := position.Add(direction)
		if topographicMap[newPosition] == value+1 {
			globalMap[position] += FindTrail(newPosition, value+1, topographicMap, directions)

		}
	}

	return globalMap[position]
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := createGrid(file)
	topographicMap := make(map[Position]int)

	for y, row := range grid {
		for x, number := range row {
			topographicMap[Position{X: x, Y: y}] = int(number - '0')
		}
	}

	directions := map[string]Position{
		"RIGHT": {X: 1, Y: 0},
		"LEFT":  {X: -1, Y: 0},
		"DOWN":  {X: 0, Y: 1},
		"UP":    {X: 0, Y: -1},
	}

	var sum int
	for position, value := range topographicMap {
		if value == 0 {
			sum += FindTrail(position, value, topographicMap, directions)
		}
	}

	fmt.Println(sum)

}
