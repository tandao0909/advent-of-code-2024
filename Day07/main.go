package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./test.txt")
	check(err)

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test, numbers := parseInput(scanner.Text())
		if checkValue(test, numbers) {
			count += test
		}
	}
	fmt.Println(count)
}

func checkValue(test int, numbers []int) bool {
	return recursiveCheck(test, numbers, 0, 0)
}

// index is the index of number to check, value is the amount of evaluated equation before that index
func recursiveCheck(test int, numbers []int, index int, value int) bool {
	if index >= len(numbers) {
		return value == test
	}
	if value > test {
		return false
	}
	return recursiveCheck(test, numbers, index+1, value+numbers[index]) ||
		recursiveCheck(test, numbers, index+1, value*numbers[index]) ||
		recursiveCheck(test, numbers, index+1, concatNumbers(value, numbers[index]))
}

func concatNumbers(a, b int) int {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	concatStr := strA + strB

	concatInt, err := strconv.Atoi(concatStr)
	check(err)

	return concatInt
}

// func printDebug(test int, numbers []int, index int, value int, evaluateSymbols []rune) {
// 	fmt.Println("Test value", test)
// 	fmt.Println("Numbers", numbers)
// 	fmt.Println("Index", index)
// 	fmt.Println("Value", value)
// 	fmt.Println("Symbols", string(evaluateSymbols))
// 	fmt.Println()
// }

func parseInput(line string) (int, []int) {
	fields := strings.Fields(line)
	test := 0
	numbers := []int{}
	for i := 0; i < len(fields); i++ {
		if i == 0 {
			testString := fields[i]
			testNumber, err := strconv.Atoi(testString[:len(testString)-1])
			check(err)
			test = testNumber
		} else {
			number, err := strconv.Atoi(fields[i])
			check(err)
			numbers = append(numbers, number)
		}
	}
	return test, numbers
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
