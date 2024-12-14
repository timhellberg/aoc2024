package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type Points struct {
	Points []Point
}

var grid = make(map[Point]rune)

func createGrid(file *os.File) map[Point]rune {
	scanner := bufio.NewScanner(file)

	var y int
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)

		for x, rune := range row {
			grid[Point{X: x, Y: y}] = rune
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return grid
}

var visitedPoints = make(map[Point]bool)

var directions = map[string][2]int{
	"RIGHT": {1, 0},
	"LEFT":  {-1, 0},
	"DOWN":  {0, 1},
	"UP":    {0, -1},
}

func getNeighbours(point Point, plant rune) ([]Point, int) {
	var points = []Point{point}
	var sides int

	if visitedPoints[point] {
		return nil, 0
	}

	visitedPoints[point] = true

	for key, dir := range directions {
		neighborPoint := Point{X: point.X + dir[0], Y: point.Y + dir[1]}

		if grid[neighborPoint] != plant {
			var alreadyChecked bool
			for key2, dir2 := range directions {
				if key == key2 {
					continue
				}
				otherNeighbourPoint := Point{X: point.X + dir2[0], Y: point.Y + dir2[1]}
				if grid[otherNeighbourPoint] == plant && grid[neighborPoint] == grid[otherNeighbourPoint] {
					alreadyChecked = true
					break
				}
			}

			if !alreadyChecked {
				sides++
			}
		}

		if grid[neighborPoint] == plant {
			neighbourPoints, neighbourSides := getNeighbours(neighborPoint, plant)
			points = append(points, neighbourPoints...)
			sides += neighbourSides
		}
	}

	return points, sides
}

func main() {
	file, err := os.Open("../testinput3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := createGrid(file)

	var regions = make(map[*Points]int)
	for point, value := range grid {
		if !visitedPoints[point] {
			points, sides := getNeighbours(point, value)
			if len(points) > 0 {
				regions[&Points{Points: points}] = sides
			}
		}
	}

	var sum int
	for key, value := range regions {
		fmt.Println(key, value)
		area := len(key.Points)
		sum += area * value
	}

	fmt.Println(sum)
}
