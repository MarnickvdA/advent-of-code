package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readInput("input.txt")

	for _, line := range lines {
		fmt.Println(line)
	}
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
