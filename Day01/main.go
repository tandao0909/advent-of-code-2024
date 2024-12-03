package main

import (
	"bufio"
	"fmt"
	"os"
	// "sort"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func splitLine(line string) []int {
	fields := strings.Fields(line)
	numList := []int{}
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		check(err)
		numList = append(numList, num)
	}
	return numList
}

func distance(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var numList1, numList2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := splitLine(scanner.Text())
		numList1 = append(numList1, numbers[0])
		numList2 = append(numList2, numbers[1])
	}

	// // Part 1
	// result := 0
	// sort.Ints(numList1)
	// sort.Ints(numList2)

	// for i := 0; i < len(numList1); i++ {
	// 	result += distance(numList1[i], numList2[i])
	// }
	// fmt.Println(result)

	// Part 2
	result := 0
	occurrences := make(map[int]int)
	for _, value := range numList2 {
		occurrences[value]++
	}
	for _, value := range numList1 {
		result += value * occurrences[value]
	}
	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
