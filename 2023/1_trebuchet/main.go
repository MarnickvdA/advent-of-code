package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	lines := readInput("input.txt")

	fmt.Println("Sum is ", getSum(lines))

}

func getSum(lines []string) int {
	sum := 0
	for x, line := range lines {
		firstNum := ""
		lastNum := ""

		for i, char := range line {
			switch {
			case unicode.IsDigit(char):
				add(&firstNum, &lastNum, char)
			case char == 'o':
				if stringNumMatcher(line, "one", i) {
					add(&firstNum, &lastNum, '1')
				}
			case char == 't':
				if stringNumMatcher(line, "two", i) {
					add(&firstNum, &lastNum, '2')
				}

				if stringNumMatcher(line, "three", i) {
					add(&firstNum, &lastNum, '3')
				}
			case char == 'f':
				if stringNumMatcher(line, "four", i) {
					add(&firstNum, &lastNum, '4')
				}

				if stringNumMatcher(line, "five", i) {
					add(&firstNum, &lastNum, '5')
				}
			case char == 's':
				if stringNumMatcher(line, "six", i) {
					add(&firstNum, &lastNum, '6')
				}
				if stringNumMatcher(line, "seven", i) {
					add(&firstNum, &lastNum, '7')
				}
			case char == 'e':
				if stringNumMatcher(line, "eight", i) {
					add(&firstNum, &lastNum, '8')
				}
			case char == 'n':
				if stringNumMatcher(line, "nine", i) {
					add(&firstNum, &lastNum, '9')
				}
			}
		}

		i, err := strconv.Atoi(firstNum + lastNum)

		if err != nil {
			panic(err)
		}

		fmt.Println(x, line, "should add", firstNum, "and", lastNum, "becoming", i)

		sum += i
	}

	return sum
}

func readInput(filename string) []string {
	bs, err := os.ReadFile(filename)

	if err != nil {
		// Option #1: log the error and call to newDeck()
		// Option #2: log the error and exit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(bs), "\n")

	return lines
}

func add(firstNum *string, lastNum *string, value rune) {
	if *firstNum == "" {
		*firstNum = string(value)
	}

	*lastNum = string(value)
}

func stringNumMatcher(curLine string, numString string, curIndex int) bool {
	lineLength := len(curLine)
	lineRunes := []rune(curLine)

	if lineLength < curIndex+len(numString) {
		return false
	}

	for i, c := range numString {
		if lineRunes[curIndex+i] != c {
			return false
		}
	}

	return true
}
