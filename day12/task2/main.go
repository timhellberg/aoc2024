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

var directions = map[string][2]int{
	"RIGHT": {1, 0},
	"LEFT":  {-1, 0},
	"DOWN":  {0, 1},
	"UP":    {0, -1},
}

type pointAndDir struct {
	point Point
	dir   [2]int
}

var visitedPoints = make(map[Point]bool)

func getNeighbours(point Point, plant rune) []Point {
	var points = []Point{point}

	if visitedPoints[point] {
		return nil
	}

	visitedPoints[point] = true

	for _, dir := range directions {
		neighborPoint := Point{X: point.X + dir[0], Y: point.Y + dir[1]}

		if grid[neighborPoint] == plant {
			neighbourPoints := getNeighbours(neighborPoint, plant)
			points = append(points, neighbourPoints...)

		}
	}

	return points
}

func getBadNeighbours(point Point, plant rune) []Point {
	var points = make([]Point, 0)

	for _, dir := range directions {
		neighborPoint := Point{X: point.X + dir[0], Y: point.Y + dir[1]}

		if grid[neighborPoint] != plant {
			points = append(points, neighborPoint)

		}
	}

	return points
}

func contains(slice []Point, value Point) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func calcSides(points []Point, badPoints []Point) int {
	var sidesMap = make(map[pointAndDir]struct{})
	var sides int
	for _, point := range points {
		for _, dir := range directions {
			if _, exists := sidesMap[pointAndDir{point: point, dir: dir}]; exists {
				continue
			}
			sidesMap[pointAndDir{point: point, dir: dir}] = struct{}{}
			neighborPoint := Point{X: point.X + dir[0], Y: point.Y + dir[1]}
			if contains(badPoints, neighborPoint) {
				sides++
			}

			for _, dir2 := range directions {
				otherNeighbourPoint1 := Point{X: point.X + dir2[0], Y: point.Y + dir2[1]}
				for contains(points, otherNeighbourPoint1) {
					sameDirPoint := Point{X: otherNeighbourPoint1.X + dir[0], Y: otherNeighbourPoint1.Y + dir[1]}
					if contains(badPoints, sameDirPoint) {
						sidesMap[pointAndDir{point: otherNeighbourPoint1, dir: dir}] = struct{}{}
					} else {
						break
					}
					otherNeighbourPoint1 = Point{X: otherNeighbourPoint1.X + dir2[0], Y: otherNeighbourPoint1.Y + dir2[1]}
				}

				otherNeighbourPoint2 := Point{X: point.X - dir2[0], Y: point.Y - dir2[1]}
				for contains(points, otherNeighbourPoint2) {
					sameDirPoint := Point{X: otherNeighbourPoint2.X + dir[0], Y: otherNeighbourPoint2.Y + dir[1]}
					if contains(badPoints, sameDirPoint) {
						sidesMap[pointAndDir{point: otherNeighbourPoint2, dir: dir}] = struct{}{}
					} else {
						break
					}
					otherNeighbourPoint2 = Point{X: otherNeighbourPoint2.X - dir2[0], Y: otherNeighbourPoint2.Y - dir2[1]}
				}

			}

		}
	}

	fmt.Println(sides)

	return sides
}

func main() {
	file, err := os.Open("../testinput2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := createGrid(file)

	var regions = make(map[*Points]int)
	for point, value := range grid {
		if !visitedPoints[point] {
			points := getNeighbours(point, value)
			if len(points) > 0 {
				var badPoints []Point
				for _, point := range points {
					badPoints = append(badPoints, getBadNeighbours(point, value)...)
				}
				fmt.Println(badPoints)
				sides := calcSides(points, badPoints)
				regions[&Points{Points: points}] = sides
			}

		}

	}

	var sum int
	for key, value := range regions {
		area := len(key.Points)
		sum += area * value
	}

	fmt.Println(sum)
}
