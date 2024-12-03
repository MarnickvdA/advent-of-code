package challenges

import (
	"math"
	"strconv"
	"strings"
)

type DayTwo struct{}

func (DayTwo) SolvePartOne(input []string) (string, error) {
	ans := 0

	for i := 0; i < len(input); i++ {
		line := strings.TrimSpace(input[i])
		if len(line) == 0 {
			continue
		}

		levels := strings.Split(line, " ")
		if checkRapport(levels, nil) {
			ans++
		}
	}

	return strconv.Itoa(ans), nil
}

func checkRapport(levels []string, pI *int) bool {
	prevDiff := 0

	for x := 1; x < len(levels); x++ {
		curr, err := strconv.Atoi(levels[x])
		prev, err := strconv.Atoi(levels[x-1])

		if err != nil {
			return false
		}

		diff := curr - prev

		// Levels are allowed to increase / decrease between 1 and 3.
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			if pI != nil {
				*pI = x
			}
			break
		}

		// Rapports should be only increasing or only decreasing.
		if (prevDiff < 0 && diff > 0) || (prevDiff > 0 && diff < 0) {
			if pI != nil {
				*pI = x
			}
			break
		}

		if x == len(levels)-1 {
			return true
		}

		// Update state
		prevDiff = diff
	}

	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func (DayTwo) SolvePartTwo(input []string) (string, error) {
	ans := 0

	for i := 0; i < len(input); i++ {
		line := strings.TrimSpace(input[i])
		if len(line) == 0 {
			continue
		}

		levels := strings.Split(line, " ")

		x := 0

		if checkRapport(levels, &x) {
			ans++
			continue
		}

		l1 := strings.Split(line, " ")
		l2 := strings.Split(line, " ")

		alt1 := remove(l1, x-1)
		alt2 := remove(l2, x)

		if checkRapport(alt1, nil) || checkRapport(alt2, nil) {
			ans++
			continue
		}

		if x == 2 && checkRapport(levels[1:], nil) {
			ans++
			continue
		}
	}

	return strconv.Itoa(ans), nil
}

func isValid(curr, next, prevDiff int) bool {
	diff := next - curr

	return math.Abs(float64(diff)) >= 1 && math.Abs(float64(diff)) <= 3 &&
		((prevDiff <= 0 && diff < 0) || (prevDiff >= 0 && diff > 0))
}
