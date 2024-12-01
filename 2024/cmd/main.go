package main

import (
	"flag"
	"log"

	"aoc.abitofsoftware.com/internal/data"
)

var flags struct {
	day  int
	part int
}

type Challenge interface {
	solvePartOne(input []string) (string, error)
	solvePartTwo(input []string) (string, error)
}

var challenges = map[int]Challenge{
	1: DayOne{},
}

func main() {
	flag.IntVar(&flags.day, "day", 0, "Challenge day to execute")
	flag.IntVar(&flags.part, "part", 1, "Part of the challenge to execute")

	flag.Parse()

	if flags.day <= 0 {
		log.Fatalln("-day flag must be positive")
	}

	if flags.part != 1 && flags.part != 2 {
		log.Fatalln("-part flag must be 1 or 2")
	}

	input, err := data.ReadInput(flags.day)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d lines of input\n", len(input))

	c := challenges[flags.day]
	if c == nil {
		log.Fatal("Challenge does not exist")
	}

	var ans string

	if flags.part == 1 {
		ans, err = c.solvePartOne(input)
	}

	if flags.part == 2 {
		ans, err = c.solvePartTwo(input)
	}

	log.Printf("Solution to day %d part %d is '%s'", flags.day, flags.part, ans)
}
