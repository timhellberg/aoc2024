package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func createList(file *os.File) []string {
	var list []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		list = strings.Split(line, "")
	}

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

func insertElements(slice []string, index int, newValue1 string, newValue2 string) []string {
	slice = append(slice, "")
	copy(slice[index+2:], slice[index+1:])
	slice[index] = newValue1
	slice[index+1] = newValue2

	return slice
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	diskMap := createList(file)

	highestId := len(diskMap) / 2
	blockedNumbers := make(map[string]struct{})

	var iteration int

	for bc := len(diskMap) - 1; bc > 1; bc-- {

		if _, exists := blockedNumbers[diskMap[bc]]; exists {
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
						blockedNumber := strconv.Itoa(filesize) + strconv.Itoa(highestId)
						diskMap = insertElements(diskMap, fw, blockedNumber, strconv.Itoa(difference))
						moved = true
						blockedNumbers[blockedNumber] = struct{}{}
						break
					}
				}
			}

			if !moved {
				blockedNumber := strconv.Itoa(filesize) + strconv.Itoa(highestId)
				diskMap[bc] = blockedNumber
				blockedNumbers[blockedNumber] = struct{}{}
			}

			iteration++
			highestId--
		}
	}

	blocks := constructBlocks(diskMap)

	var sum int
	for index, char := range blocks {
		if char != "." {
			number, _ := strconv.Atoi(char)
			sum += number * index
		}

	}

	fmt.Println(sum)
}
