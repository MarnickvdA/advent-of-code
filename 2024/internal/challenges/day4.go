package challenges

import (
	"strconv"
)

type DayFour struct{}

func (DayFour) SolvePartOne(input []string) (string, error) {
	m := make([][]rune, len(input))
	for i := range m {
		m[i] = make([]rune, len(input[i]))

		for x, rune := range input[i] {
			m[i][x] = rune
		}
	}

	// traverse the matrix
	// for every possibility in bounds, check the following 8 possible directions:
	//
	//  8   1   2
	//   S..S..S
	//   .A.A.A.
	//   ..MMM..
	// 7 SAMXMAS 3
	//   ..MMM..
	//   .A.A.A.
	//   S..S..S
	//  6   5   4

	checkNorth := func(x, y int) bool {
		return y-3 >= 0 && m[y-1][x] == 'M' && m[y-2][x] == 'A' && m[y-3][x] == 'S'
	}

	checkNorthEast := func(x, y int) bool {
		return y-3 >= 0 && x+3 < len(m[y]) && m[y-1][x+1] == 'M' && m[y-2][x+2] == 'A' &&
			m[y-3][x+3] == 'S'
	}

	checkEast := func(x, y int) bool {
		return x+3 < len(m[y]) && m[y][x+1] == 'M' && m[y][x+2] == 'A' && m[y][x+3] == 'S'
	}

	checkSouthEast := func(x, y int) bool {
		return y+3 < len(m) && x+3 < len(m[y+3]) && m[y+1][x+1] == 'M' && m[y+2][x+2] == 'A' &&
			m[y+3][x+3] == 'S'
	}

	checkSouth := func(x, y int) bool {
		return y+3 < len(m) && m[y+1][x] == 'M' && m[y+2][x] == 'A' && m[y+3][x] == 'S'
	}

	checkSouthWest := func(x, y int) bool {
		return y+3 < len(m) && x-3 >= 0 && m[y+1][x-1] == 'M' && m[y+2][x-2] == 'A' &&
			m[y+3][x-3] == 'S'
	}

	checkWest := func(x, y int) bool {
		return x-3 >= 0 && m[y][x-1] == 'M' && m[y][x-2] == 'A' && m[y][x-3] == 'S'
	}

	checkNorthWest := func(x, y int) bool {
		return y-3 >= 0 && x-3 >= 0 && m[y-1][x-1] == 'M' && m[y-2][x-2] == 'A' &&
			m[y-3][x-3] == 'S'
	}

	directions := []func(int, int) bool{
		checkNorth,     // #1
		checkNorthEast, // #2
		checkEast,      // #3
		checkSouthEast, // #4
		checkSouth,     // #5
		checkSouthWest, // #6
		checkWest,      // #7
		checkNorthWest, // #8
	}

	sum := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] != 'X' {
				continue
			}

			for _, direction := range directions {
				if direction(x, y) {
					sum++
				}
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func (DayFour) SolvePartTwo(input []string) (string, error) {
	m := make([][]rune, len(input))
	for i := range m {
		m[i] = make([]rune, len(input[i]))

		for x, rune := range input[i] {
			m[i][x] = rune
		}
	}

	// Instead of X, we are going to traverse based on the 'hit' of letter A.
	sum := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] != 'A' {
				continue
			}

			// Don't iterate over the outer bounding values, since they won't satisfy anyways.
			if y == 0 || x == 0 || y == len(m)-1 || x == len(m[y])-1 {
				continue
			}

			allowed := map[rune]bool{
				'M': true,
				'S': true,
			}
			if !allowed[m[y-1][x-1]] || !allowed[m[y-1][x+1]] || !allowed[m[y+1][x-1]] ||
				!allowed[m[y+1][x+1]] {
				continue
			}

			if m[y-1][x-1] != m[y+1][x+1] && m[y-1][x+1] != m[y+1][x-1] {
				sum++
			}
		}
	}

	return strconv.Itoa(sum), nil
}
