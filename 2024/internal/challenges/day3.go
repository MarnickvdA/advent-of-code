package challenges

import (
	"regexp"
	"strconv"
	"strings"
)

type DayThree struct{}

func (DayThree) SolvePartOne(input []string) (string, error) {
	sum := 0

	// retrieve all matches for regex mul(X,Y) where X & Y are between 1 and 3 digits.
	r := regexp.MustCompile(`(mul\(([0-9]){1,3},([0-9]){1,3}\))`)

	// This challenge only contains one line.
	line := input[0]

	// minimum length of 1 entry is 7 characters, so max match array size is line length / 7
	matches := r.FindAllString(line, len(line)/7)

	for _, m := range matches {
		values := strings.Split(string(m[4:len(m)-1]), ",")

		a, err := strconv.Atoi(values[0])
		b, err := strconv.Atoi(values[1])

		if err != nil {
			return "", err
		}

		sum += a * b
	}

	return strconv.Itoa(sum), nil
}

func (DayThree) SolvePartTwo(input []string) (string, error) {
	sum := 0

	// retrieve all matches for regex mul(X,Y) where X & Y are between 1 and 3 digits.
	r := regexp.MustCompile(`mul\(([0-9]){1,3},([0-9]){1,3}\)`)
	dor := regexp.MustCompile(`do\(\)`)
	dontr := regexp.MustCompile(`don\'t\(\)`)

	// This challenge only contains one line.
	line := input[0]

	// retrieve all the 'non-execute' windows for every line
	// window is defined by ______don't()xxxxxxxxxxxdo()______don't()xxxxxxdo()___
	// so, if a mul() entry is in the nonexec window, we shouldn't do the multiplication.
	doIdx := dor.FindAllStringIndex(line, len(line)/4)
	dontIdx := dontr.FindAllStringIndex(line, len(line)/7)

	canSum := make([]int, len(line))

	currDo, currDont := 0, 0
	currVal := 1
	for i := range canSum {
		if currDo < len(doIdx) && i == doIdx[currDo][0] {
			currVal = 1
			currDo++
		}

		if currDont < len(dontIdx) && i == dontIdx[currDont][0] {
			currVal = 0
			currDont++
		}

		canSum[i] = currVal
	}

	// minimum length of 1 entry is 7 characters, so max match array size is line length / 7
	matches := r.FindAllString(line, len(line)/7)
	matchIdx := r.FindAllStringIndex(line, len(line)/7)

	for i, m := range matches {
		if canSum[matchIdx[i][0]] == 0 {
			continue
		}

		values := strings.Split(string(m[4:len(m)-1]), ",")

		a, err := strconv.Atoi(values[0])
		if err != nil {
			return "", err
		}

		b, err := strconv.Atoi(values[1])
		if err != nil {
			return "", err
		}

		sum += a * b
	}

	return strconv.Itoa(sum), nil
}
