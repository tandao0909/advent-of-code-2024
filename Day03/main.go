package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)


func main() {

	file, err := os.Open("data.txt")
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
	// The elegant is in the fact that we can check the combine regex, and alternate the execution accordingly
	// Inspired by https://www.reddit.com/r/adventofcode/comments/1h5obsr/2024_day_3_regular_expressions_go_brrr/
	// If your code isn't working, try to make the input file on the same line
	// My original code works, but it's very ugly, you can see it in the commit history
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	mulRe := regexp.MustCompile(`mul\(\d+,\d+\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)
	execute := true
	result := 0

	matches := re.FindAllStringIndex(text, -1)
	for _, match := range matches {
		substr := text[match[0]:match[1]]
		if doRe.MatchString(substr) {
			execute = true
		} else if dontRe.MatchString(substr) {
			execute = false
		} else if mulRe.MatchString(substr) {
			if execute {
				reNum := regexp.MustCompile(`\d+`)
				numberStrings := reNum.FindAllString(substr, -1)
				multiplication_result := 1
		
				for _, num := range numberStrings {
					number, err := strconv.Atoi(num)
					check(err)
					multiplication_result *= number
				}
		
				result += multiplication_result
			}
		}
	}

	return result
}
