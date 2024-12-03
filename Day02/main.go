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
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		report := splitLine(scanner.Text())
		// // Part 1
		// result += checkReport(report)
		// Part 2
		result += checkReportPart2(report)
	}
	fmt.Println(result)
}

func checkReport(report []int) int {
	status := "increase"
	if report[0] > report[1] {
		status = "decrease"
	}

	for i := 0; i < len(report)-1; i++ {
		if status == "increase" {
			if report[i] > report[i+1] || checkThreshold(report[i+1]-report[i]) {
				return 0
			}
		} else {
			if report[i+1] > report[i] || checkThreshold(report[i]-report[i+1]) {
				return 0
			}
		}
	}
	return 1
}

func checkReportPart2(report []int) int {
	if checkReport(report) == 1 {
		return 1
	}
	for i := 0; i < len(report); i++ {
		newReport := append([]int{}, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if checkReport(newReport) == 1 {
			fmt.Println(report)

			return 1
		}
	}
	return 0
}

// Check if threshold is out of the defined range, i.e., < 1 or > 3
func checkThreshold(threshold int) bool {
	if threshold < 1 || threshold > 3 {
		return true
	}
	return false
}

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
