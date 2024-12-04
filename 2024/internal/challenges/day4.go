package challenges

type DayFour struct{}

func (DayFour) SolvePartOne(input []string) (int, error) {
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
	ds := []struct {
		dx, dy int
	}{
		{dx: 0, dy: -1},  // North
		{dx: 1, dy: -1},  // North East
		{dx: 1, dy: 0},   // East
		{dx: 1, dy: 1},   // South East
		{dx: 0, dy: 1},   // South
		{dx: -1, dy: 1},  // South West
		{dx: -1, dy: 0},  // West
		{dx: -1, dy: -1}, // North West
	}

	sum := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] != 'X' {
				continue
			}

			for _, d := range ds {
				if y+(d.dy*3) < 0 || y+(d.dy*3) >= len(m) ||
					x+(d.dx*3) < 0 || x+(d.dx*3) >= len(m[0]) {
					continue
				}

				if m[y+(d.dy*1)][x+(d.dx*1)] == 'M' &&
					m[y+(d.dy*2)][x+(d.dx*2)] == 'A' &&
					m[y+(d.dy*3)][x+(d.dx*3)] == 'S' {
					sum++
				}
			}
		}
	}

	return sum, nil
}

func (DayFour) SolvePartTwo(input []string) (int, error) {
	m := make([][]rune, len(input))
	for i := range m {
		m[i] = make([]rune, len(input[i]))

		for x, rune := range input[i] {
			m[i][x] = rune
		}
	}

	allowed := map[rune]bool{
		'M': true,
		'S': true,
	}

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

			if !allowed[m[y-1][x-1]] || !allowed[m[y-1][x+1]] || !allowed[m[y+1][x-1]] ||
				!allowed[m[y+1][x+1]] {
				continue
			}

			if m[y-1][x-1] != m[y+1][x+1] && m[y-1][x+1] != m[y+1][x-1] {
				sum++
			}
		}
	}

	return sum, nil
}
