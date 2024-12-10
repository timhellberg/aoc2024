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

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	grid := createGrid(file)
	antinodes := make(map[complex128]struct{})
	for y1, row := range grid {
		for x1, char := range row {
			if char != '.' {
				lowerCharPosition := complex(float64(x1), float64(y1))

				for y2 := y1; y2 < len(grid); y2++ {
					for x2 := 0; x2 < len(row); x2++ {

						if y2 == y1 && x2 <= x1 {
							continue
						}

						if grid[y2][x2] == char {
							higherCharPosition := complex(float64(x2), float64(y2))
							distance := higherCharPosition - lowerCharPosition

							lowerAntiNode := lowerCharPosition
							for real(lowerAntiNode) >= 0 && real(lowerAntiNode) < float64(len(row)) && imag(lowerAntiNode) >= 0 {
								antinodes[lowerAntiNode] = struct{}{}
								lowerAntiNode -= distance
							}

							higherAntiNode := higherCharPosition
							for real(higherAntiNode) >= 0 && real(higherAntiNode) < float64(len(row)) && imag(higherAntiNode) < float64(len(grid)) {
								antinodes[higherAntiNode] = struct{}{}
								higherAntiNode += distance
							}

						}
					}
				}

			}
		}
	}

	sum := len(antinodes)
	fmt.Println(sum)

}
