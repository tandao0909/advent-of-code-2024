package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	text := []string{}
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	// fmt.Println(text)
	// // Part 1
	// fmt.Println(evaluate(text))
	// Part 2
	fmt.Println(evaluatePart2(text))
	// fmt.Println(checkDiagonal(text))
}

func evaluate(text []string) int {
	count := 0
	count += countHorizontal(text)
	count += countVertical(text)
	count += countDiagonal(text)
	count += countDiagonal2(text)
	return count
}

func evaluatePart2(text []string) int {
	count := 0

	for i := 0; i < len(text)-2; i++ {
		for j := 0; j < len(text[i])-2; j++ {
			// We parse the the 3x3 subtext
			subtext := []string{}
			for k := 0; k < 3; k++ {
				substring := ""
				for l := 0; l < 3; l++ {
					substring += string(text[i+k][j+l])
				}
				subtext = append(subtext, substring)
			}
			if checkDiagonal(subtext) {
				count++
			}
		}
	}
	return count
}

func checkDiagonal(text []string) bool {
	extendText := extendMatrix(text)
	flag := true

	for i := 0; i < len(text) - 2; i++ {
		for j := 0; j < len(text[i]) - 2; j++ {
			substring := ""
			for k := 0; k < 3; k++ {
				substring += string(extendText[i+k][j+k])
			}
			if substring != "MAS" && reverseString(substring) != "MAS" {
				flag = false
			}
		}
	}
	for i := 0; i < len(text) - 2; i++ {
		for j := 2; j < len(text[i]); j++ {
			substring := ""
			for k := 0; k < 3; k++ {
				substring += string(extendText[i+k][j-k])
			}
			if substring != "MAS" && reverseString(substring) != "MAS" {
				flag = false
			}
		}
	}
	return flag
}

func countHorizontal(text []string) int {
	count := 0
	extendText := extendMatrix(text)
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			substring := extendText[i][j : j+4]
			if substring == "XMAS" || reverseString(substring) == "XMAS" {
				count++
			}
		}
	}
	return count
}

func countVertical(text []string) int {
	count := 0
	extendText := extendMatrix(text)

	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			substring := ""
			for k := 0; k < 4; k++ {
				substring += string(extendText[i+k][j])
			}
			if substring == "XMAS" || reverseString(substring) == "XMAS" {
				count++
			}
		}
	}
	return count
}

func countDiagonal(text []string) int {
	count := 0
	extendText := extendMatrix(text)

	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			substring := ""
			for k := 0; k < 4; k++ {
				substring += string(extendText[i+k][j+k])
			}
			if substring == "XMAS" || reverseString(substring) == "XMAS" {
				count++
			}
		}
	}
	return count
}

func countDiagonal2(text []string) int {
	count := 0
	extendText := extendMatrix(text)

	for i := 0; i < len(text); i++ {
		for j := 3; j < len(text[i]); j++ {
			substring := ""
			for k := 0; k < 4; k++ {
				substring += string(extendText[i+k][j-k])
			}
			if substring == "XMAS" || reverseString(substring) == "XMAS" {
				count++
			}
		}
	}
	return count
}

func extendMatrix(text []string) []string {
	textCopy := make([]string, len(text))
	copy(textCopy, text)
	// Extend each row to the right
	right, bottom := 3, 3
	for i := range textCopy {
		textCopy[i] += string(make([]rune, right))
		for j := 0; j < right; j++ {
			textCopy[i] += "n"
		}
	}

	// Add new rows at the bottom
	rowLength := len(textCopy[0])
	for i := 0; i < bottom; i++ {
		newRow := ""
		for j := 0; j < rowLength; j++ {
			newRow += "n"
		}
		textCopy = append(textCopy, newRow)
	}

	return textCopy
}
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
