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

var sidesMap = make(map[Point][2]int)
var directions = map[string][2]int{
	"RIGHT": {1, 0},
	"LEFT":  {-1, 0},
	"DOWN":  {0, 1},
	"UP":    {0, -1},
}

var sides int

func getNeighbours(point Point, plant rune) []Point {
	var points = []Point{point}

	if visitedPoints[point] {
		return nil
	}

	visitedPoints[point] = true

	for key, dir := range directions {
		neighborPoint := Point{X: point.X + dir[0], Y: point.Y + dir[1]}

		if grid[neighborPoint] != plant {
			sidesMap[point] = [2]int{dir[0], dir[1]}
			var alreadyChecked bool
			for key2, dir2 := range directions {
				if key == key2 {
					continue
				}
				otherNeighbourPoint := Point{X: point.X + dir2[0], Y: point.Y + dir2[1]}
				if grid[otherNeighbourPoint] == plant && sidesMap[otherNeighbourPoint] == dir {
					alreadyChecked = true
					break
				}
			}

			if !alreadyChecked {
				sides++
			}

		}

		if grid[neighborPoint] == plant {
			points = append(points, getNeighbours(neighborPoint, plant)...)
		}
	}

	return points
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
		points := getNeighbours(point, value)

		regions[&Points{Points: points}] = sides
		sides = 0
	}

	var sum int
	for key, value := range regions {
		area := len(key.Points)
		sum += area * value

	}

	fmt.Println(sum)
}
