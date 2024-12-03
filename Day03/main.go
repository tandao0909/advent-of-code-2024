package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		text := scanner.Text()
		// // // Part 1
		// result += checkMemory(text)
		//  Part 2
		result += checkMemoryWithCondition(text)
	}
	fmt.Println(result)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func checkMemory(text string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	result := 0

	matches := re.FindAllString(text, -1)

	for _, match := range matches {
		reNum := regexp.MustCompile(`\d+`)
		numberStrings := reNum.FindAllString(match, -1)
		multiplication_result := 1

		for _, num := range numberStrings {
			number, err := strconv.Atoi(num)
			check(err)
			multiplication_result *= number
		}

		result += multiplication_result
	}
	return result
}

func checkMemoryWithCondition(text string) int {
	// We want to have the index of all the match mul(), do(), don't(), and have some switch.
	// Since the regex find exactly the index of "do()" and "don't()" string,
	// we will extract the first index into a list.
	fmt.Println(text)

	reDo := regexp.MustCompile(`do\(\)`)
	doMatches := reDo.FindAllStringIndex(text, -1)
	doMatchIndices := []int{}
	for _, value := range doMatches {
		doMatchIndices = append(doMatchIndices, value[0])
	}
	fmt.Println("Do match indices", doMatchIndices)

	reDont := regexp.MustCompile(`don't\(\)`)
	dontMatches := reDont.FindAllStringIndex(text, -1)
	dontMatchIndices := []int{}
	for _, value := range dontMatches {
		dontMatchIndices = append(dontMatchIndices, value[0])
	}
	fmt.Println("Don't match indices", dontMatchIndices)

	reMul := regexp.MustCompile(`mul\(\d+,\d+\)`)
	mulMatches := reMul.FindAllStringIndex(text, -1)
	result := 0
	for _, value := range mulMatches {
		if checkExecute(value[0], doMatchIndices, dontMatchIndices) {
			reNum := regexp.MustCompile(`\d+`)
			executionText := text[value[0]: value[1]]
			numberStrings := reNum.FindAllString(executionText, -1)
			multiplication_result := 1

			for _, num := range numberStrings {
				number, err := strconv.Atoi(num)
				check(err)
				multiplication_result *= number
			}

			result += multiplication_result
		}
	}

	return result
}

func checkExecute(mulIndex int, doIndices []int, dontIndices []int) bool {
	nearestDo := largestSmallerThan(doIndices, mulIndex)
	nearestDont := largestSmallerThan(dontIndices, mulIndex)
	//  If there is no don't before it, just execute
	if nearestDont == -1 {
		return true
	} else if nearestDont > nearestDo { // We just don't execute if there's a don't right before us
		return false
	} else {
		return true
	}
}

func largestSmallerThan(numbers []int, target int) int {
	largest := -1
	for _, num := range numbers {
		if num < target && num > largest {
			largest = num
		}
	}
	return largest
}
