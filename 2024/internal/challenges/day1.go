package challenges

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type DayOne struct{}

func (DayOne) SolvePartOne(input []string) (int, error) {
	var sum float64

	var left, right []int

	for i := 0; i < len(input); i++ {
		segs := strings.Split(input[i], "   ")

		if len(segs) != 2 {
			// log.Printf("invalid segments in line #%d with '%s'", i, input[i])
			continue
		}

		a, _ := strconv.Atoi(segs[0])
		b, _ := strconv.Atoi(segs[1])

		left = append(left, a)
		right = append(right, b)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		sum += math.Abs(float64(left[i] - right[i]))
	}

	return int(sum), nil
}

func (DayOne) SolvePartTwo(input []string) (int, error) {
	sim := 0
	rightCount := map[int]int{}

	for i := 0; i < len(input); i++ {
		segs := strings.Split(input[i], "   ")

		if len(segs) != 2 {
			// log.Printf("invalid segments in line #%d with '%s'", i, input[i])
			continue
		}

		b, _ := strconv.Atoi(segs[1])
		rightCount[b] = rightCount[b] + 1
	}

	for i := 0; i < len(input); i++ {
		segs := strings.Split(input[i], "   ")

		if len(segs) != 2 {
			// log.Printf("invalid segments in line #%d with '%s'", i, input[i])
			continue
		}

		a, _ := strconv.Atoi(segs[0])
		sim += a * rightCount[a]
	}

	return sim, nil
}
