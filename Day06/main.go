package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var direction_symbols = [4]rune{'^', '>', 'v', '<'}
var direction_map = map[rune][]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func main() {
	file, err := os.Open("data.txt")
	check(err)
	defer file.Close()

	guardMap := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		guardMap = append(guardMap, scanner.Text())
	}
	// // Part 1
	// walk(guardMap)
	// fmt.Println(countXGuardMap(guardMap))
	// Part 2
	count := 0
	start := time.Now()
	var mu sync.Mutex
	var wg sync.WaitGroup

	for row_index, row := range guardMap {
		for column_index := range row {
			wg.Add(1)
			go func(row_index, column_index int) {
				defer wg.Done()
				if checkPosition(guardMap, row_index, column_index) {
					mu.Lock()
					count++
					mu.Unlock()
				}
			}(row_index, column_index)
		}
	}

	wg.Wait()
	fmt.Println(count)
	// Took 9m25.144748479s
	fmt.Printf("Took %s\n", time.Since(start))
}

// Return if we can place 'O' at this position
func checkPosition(guardMap []string, row_index int, column_index int) bool {
	startPosition, _ := startingPosition(guardMap)
	guardMapCopy := make([]string, len(guardMap))
	copy(guardMapCopy, guardMap)
	if row_index == startPosition[0] && column_index == startPosition[1] {
		return false
	} else if guardMap[row_index][column_index] == '#' {
		return false
	} else if checkWalk(guardMapCopy, row_index, column_index) {
		return true
	} else {
		return false
	}
}

// Return if we place 'O' on this position, we will stuck in a infinite loop
func checkWalk(guardMap []string, row_index int, column_index int) bool {
	guardIndex, guardDirectionIndex := startingPosition(guardMap)
	rowNumber, columnNumber := len(guardMap), len(guardMap[0])
	// If we are not stuck in a loop, then in each cell, there are only be at most 4 steps: Up right left down. 
	// If you repeat the same action again, you must have stuck.
	maxSteps := rowNumber * columnNumber * 4
	numSteps := 0

	changeGuardMap(guardMap, [2]int{row_index, column_index}, '#')
	for (numSteps < maxSteps){
		symbol := direction_symbols[guardDirectionIndex]
		direction := direction_map[symbol]
		nextIndex := [2]int{guardIndex[0] + direction[0], guardIndex[1] + direction[1]}
		if nextIndex[0] < 0 || nextIndex[0] >= rowNumber || nextIndex[1] < 0 || nextIndex[1] >= columnNumber {
			changeGuardMap(guardMap, guardIndex, 'X')
			return false
		}
		rowOfNextIndex := guardMap[nextIndex[0]]
		symbolNextStep := rowOfNextIndex[nextIndex[1]]
		if symbolNextStep == '#' {
			guardDirectionIndex = (guardDirectionIndex + 1) % len(direction_symbols)
			numSteps++
		} else {
			changeGuardMap(guardMap, guardIndex, 'X')
			guardIndex = nextIndex
		}
	}
	return true
}

func walk(guardMap []string) {
	guardIndex, guardDirectionIndex := startingPosition(guardMap)
	rowNumber, columnNumber := len(guardMap), len(guardMap[0])
	for {
		symbol := direction_symbols[guardDirectionIndex]
		direction := direction_map[symbol]
		nextIndex := [2]int{guardIndex[0] + direction[0], guardIndex[1] + direction[1]}
		if nextIndex[0] < 0 || nextIndex[0] >= rowNumber || nextIndex[1] < 0 || nextIndex[1] >= columnNumber {
			changeGuardMap(guardMap, guardIndex, 'X')
			break
		}
		rowOfNextIndex := guardMap[nextIndex[0]]
		symbolNextStep := rowOfNextIndex[nextIndex[1]]
		if symbolNextStep == '#' {
			guardDirectionIndex = (guardDirectionIndex + 1) % len(direction_symbols)
		} else {
			changeGuardMap(guardMap, guardIndex, 'X')
			guardIndex = nextIndex
		}
	}
}

func changeGuardMap(guardMap []string, index [2]int, char rune) {
	row_of_guard := []rune(guardMap[index[0]])
	row_of_guard[index[1]] = char
	guardMap[index[0]] = string(row_of_guard)
}

func printMap(guardMap []string) {
	fmt.Println()
	for _, row := range guardMap {
		fmt.Println(row)
	}
}

func countXGuardMap(guardMap []string) int {
	count := 0
	for _, row := range guardMap {
		for _, char := range row {
			if char == 'X' {
				count ++
			}
		}
	}
	return count
}

// Returns starting position and starting direction index in the directions array
func startingPosition(guardMap []string) (start_index [2]int, start_direction_index int) {
	for row_index, row := range guardMap {
		for column_index, char := range row {
			for i := 0; i < len(direction_symbols); i++ {
				if direction_symbols[i] == char {
					return [2]int{row_index, column_index}, i
				}
			}
		}
	}
	fmt.Println("The guard map does not have a starting symbol")
	return [2]int{0, 0}, 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
