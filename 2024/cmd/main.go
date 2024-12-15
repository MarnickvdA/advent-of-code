package main

import (
	"flag"
	"log"
	"time"

	c "aoc.abitofsoftware.com/internal/challenges"
	"aoc.abitofsoftware.com/internal/data"
)

var flags struct {
	day  int
	part int
}

var challenges = map[int]c.Challenge{
	1:  c.DayOne{},
	2:  c.DayTwo{},
	3:  c.DayThree{},
	4:  c.DayFour{},
	5:  c.DayFive{},
	6:  c.DaySix{},
	7:  c.DaySeven{},
	8:  c.DayEight{},
	9:  c.DayNine{},
	10: c.DayTen{},
	11: c.DayEleven{},
	12: c.DayTwelve{},
	13: c.DayThirteen{},
	14: c.DayFourteen{},
	15: c.DayFifteen{},
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
		log.Fatal("Challenge does not exist, did you forget to add it?")
	}

	var ans int
	var elapsed time.Duration

	if flags.part == 1 {
		start := time.Now()
		ans, err = c.SolvePartOne(input)
		elapsed = time.Since(start)
	}

	if flags.part == 2 {
		start := time.Now()
		ans, err = c.SolvePartTwo(input)
		elapsed = time.Since(start)
	}

	if err != nil {
		log.Fatal("Challenge failed to be solved due to an error", err)
	} else {
		log.Printf("Solution to day %d part %d is '%d' (in %s)", flags.day, flags.part, ans, elapsed)
	}
}
