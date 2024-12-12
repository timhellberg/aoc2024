package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func createSpecialSlice(file *os.File) []string {
	var list []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		list = strings.Split(line, "")
	}

	//Makes all free space numbers negative
	for i := 1; i < len(list); i += 2 {
		list[i] = "-" + list[i]
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return list
}

func constructBlocks(diskMap []string) []string {
	var result []string

	for index, entry := range diskMap {
		var size int
		if entry[0] == '-' && entry[1] != '0' {
			number, _ := strconv.Atoi(entry)
			size = int(math.Abs(float64(number)))
			list := make([]string, size, size)
			for i := 0; i < size; i++ {
				list[i] = "."
			}
			result = append(result, list...)
		} else if index == 0 {
			size, _ := strconv.Atoi(entry)
			list := make([]string, size, size)
			for i := 0; i < size; i++ {
				list[i] = "."
			}
			result = append(result, list...)
		} else if entry[0] != '-' && len(entry) >= 2 {
			var id int
			size := int(entry[0] - '0')

			if len(entry) == 2 {
				id = int(entry[1] - '0')
			} else if len(entry) > 2 {
				num, _ := strconv.Atoi(entry[1:])
				id = num
			}

			list := make([]string, size, size)
			for i := 0; i < size; i++ {
				list[i] = strconv.Itoa(id)
			}
			result = append(result, list...)
		}

	}

	return result
}

func getCheckSumFromBlock(blocks []string) int {
	var sum int
	for index, char := range blocks {
		if char != "." {
			number, _ := strconv.Atoi(char)
			sum += number * index
		}

	}

	return sum
}

// faster
func getCheckSumFromDiskMap(diskMap []string) int {
	var result int
	var currentIndex int
	for index, entry := range diskMap {
		if index == 0 {
			size := int(entry[0] - '0')
			currentIndex += size
			continue
		}

		if entry[0] == '-' {
			size := int(entry[1] - '0')
			currentIndex += size
			continue
		}

		size := int(entry[0] - '0')
		id, _ := strconv.Atoi(entry[1:])
		for i := currentIndex; i < currentIndex+size; i++ {
			result += id * i
		}
		currentIndex += size

	}
	return result
}

func insertElements(slice []string, index int, newValue1 string, newValue2 string) []string {
	slice = append(slice, "")
	copy(slice[index+2:], slice[index+1:])
	slice[index] = newValue1
	slice[index+1] = newValue2

	return slice
}

/*
A special slice is the input where all the free spaces are negative (every other number) is negative.
So for test input 2333133121414131402 a special slice is [2 -3 3 -3 1 -3 3 -1 2 -1 4 -1 4 -1 3 -1 4 -0 2]
The slice is then sorted by moving positive numbers from the right to the first negative number from the left that is big enough
to hold it.
Moving means:

	-inserting the positive number + id at the index of the negative number
	-appending the remaining difference (remaining free space) at index+1 of the negative number
	-Making the positive number negative at the index of the number that was moved

The special slice for the test input would sort like this:
No iterations: [2 -3 3 -3 1 -3 3 -1 2 -1 4 -1 4 -1 3 -1 4 -0 2]
Iteration 1: [2 29 -1 3 -3 1 -3 3 -1 2 -1 4 -1 4 -1 3 -1 4 -0 -2] 2 becomes 29 (size 2 id 9)  at index 1, remaining difference -1 at index 2
Iteration 2: [2 29 -1 3 -3 1 -3 3 -1 2 -1 4 -1 4 -1 3 -1 48 -0 -2] 4 dosn't move but becomes 48 (size 4 id 8)

Result: [2 29 12 31 37 -1 24 -1 33 -1 -2 -1 45 -1 46 -1 -3 -1 48 -0 -2] The first number does not get assigned id 0, the calculation logic handles this case instead
*/
func SortSpecialSlice(diskMap []string) []string {
	highestId := len(diskMap) / 2
	processedNumbers := make(map[string]struct{})

	for bc := len(diskMap) - 1; bc > 1; bc-- {
		if _, exists := processedNumbers[diskMap[bc]]; exists {
			continue
		}

		moved := false
		if diskMap[bc][0] != '-' && diskMap[bc][0] != '0' {
			filesize, _ := strconv.Atoi(diskMap[bc])
			for fw := 0; fw < bc; fw++ {
				if diskMap[fw][0] == '-' {
					freeSpace, _ := strconv.Atoi(diskMap[fw])
					difference := filesize + freeSpace
					if difference <= 0 {
						diskMap[bc] = "-" + strconv.Itoa(filesize)
						processedNumber := strconv.Itoa(filesize) + strconv.Itoa(highestId)
						if difference < 0 {
							diskMap = insertElements(diskMap, fw, processedNumber, strconv.Itoa(difference))
						} else if difference == 0 {
							diskMap[fw] = processedNumber
						}
						moved = true
						processedNumbers[processedNumber] = struct{}{}
						break
					}
				}
			}

			if !moved {
				processedNumber := strconv.Itoa(filesize) + strconv.Itoa(highestId)
				diskMap[bc] = processedNumber
				processedNumbers[processedNumber] = struct{}{}
			}

			highestId--
		}

	}
	return diskMap
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	start := time.Now()

	diskMap := createSpecialSlice(file)
	diskMap = SortSpecialSlice(diskMap)

	sum := getCheckSumFromDiskMap(diskMap)
	fmt.Println(sum)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
