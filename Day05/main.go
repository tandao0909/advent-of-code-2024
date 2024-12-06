package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	check(err)
	defer file.Close()

	updates := [][]int{}
	orders := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ContainsRune(text, '|') {
			order := processOrdering(text)
			orders[order[0]] = append(orders[order[0]], order[1])

		} else if strings.ContainsRune(text, ',') {
			update := processUpdate(text)
			updates = append(updates, update)
		}
	}

	count := 0
	countPart2 := 0
	for _, update := range updates {
		if checkUpdate(update, orders) {
			count += update[len(update)/2]
		} else {
			countPart2 += middleOfCorrectedUpdate(update, orders)
		}
	}

	fmt.Println(count)
	fmt.Println(countPart2)
}

func reverseOfCorrectUpdate(update []int, orders map[int][]int) []int {
	newUpdate := make([]int, len(update))
	for i := 0; i < len(update); i++ {
		firstPage := update[i]
		// The think is all number behind a number in a update is in its order
		// Hence, the index of a number in the reverse of a correct update is the amount of numbers of the update
		// that in its orders
		index := 0
		index += countElementsInSlice(update[:i], orders[firstPage])
		index += countElementsInSlice(update[i+1:], orders[firstPage])
		newUpdate[index] = firstPage
	}
	return newUpdate
}

func middleOfCorrectedUpdate(update []int, orders map[int][]int) int {
	return reverseOfCorrectUpdate(update, orders)[len(update)/2]
}

// Count how many elements of a is in b
func countElementsInSlice(a, b []int) int {
	count := 0
	for _, elem := range a {
		if contains(b, elem) {
			count++
		}
	}
	return count
}

func checkUpdate(update []int, orders map[int][]int) bool {
	for i := 0; i < len(update); i++ {
		firstPage := update[i]
		for j := i + 1; j < len(update); j++ {
			if !contains(orders[firstPage], update[j]) {
				return false
			}
		}
	}
	return true
}

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

func processOrdering(text string) []int {
	numbers := []int{}
	for _, string := range strings.Split(text, "|") {
		number, err := strconv.Atoi(string)
		check(err)
		numbers = append(numbers, number)
	}
	return numbers
}

func processUpdate(text string) []int {
	numbers := []int{}
	for _, string := range strings.Split(text, ",") {
		number, err := strconv.Atoi(string)
		check(err)
		numbers = append(numbers, number)
	}
	return numbers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
