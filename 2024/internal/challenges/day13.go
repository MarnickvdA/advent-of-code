package challenges

import (
	"regexp"
	"strconv"
)

type DayThirteen struct{}

func (DayThirteen) SolvePartOne(input []string) (int, error) {
	cs := parseClawContraptions(&input)

	m := 0

	for _, c := range cs {
		tokens := bruteforce(c.a, c.b, c.t)

		m += tokens
	}

	return m, nil
}

func (DayThirteen) SolvePartTwo(input []string) (int, error) {
	return 0, nil
}

type Button struct {
	dx int
	dy int
}

type Location struct {
	x int
	y int
}

type ClawContraption struct {
	a Button
	b Button
	t Location
}

func parseClawContraptions(input *[]string) []*ClawContraption {
	e := []*ClawContraption{}

	for i := range *input {
		if i%4 == 0 {
			e = append(e, &ClawContraption{})
		}

		if i%4 == 3 {
			continue
		}

		c := e[i/4]

		r := regexp.MustCompile(`([0-9])+`)
		s := r.FindAllString((*input)[i], 2)

		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		switch i % 4 {
		case 0:
			c.a = Button{dx: x, dy: y}
		case 1:
			c.b = Button{dx: x, dy: y}
		case 2:
			c.t = Location{x: x, y: y}
		}

	}

	return e
}

func dfs(a, b Button, l Location, x, y, ac, bc, sum int, sums *[]int) {
	if x+a.dx == l.x && y+a.dy == l.y && ac+1 <= 100 {
		*sums = append(*sums, sum+3)
		return
	}

	if x+b.dx == l.x && y+b.dy == l.y && bc+1 <= 100 {
		*sums = append(*sums, sum+1)
		return
	}

	if x+a.dx < l.x && y+a.dy < l.y && ac+1 < 100 {
		dfs(a, b, l, x+a.dx, y+a.dy, ac+1, bc, sum+3, sums)
	}

	if x+b.dx < l.x && y+b.dy < l.y && bc+1 < 100 {
		dfs(a, b, l, x+b.dx, y+b.dy, ac, bc+1, sum+1, sums)
	}
}

func bruteforce(a, b Button, l Location) int {
	cheapest := 401
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i*a.dx+j*b.dx == l.x && i*a.dy+j*b.dy == l.y && i*3+j < cheapest {
				cheapest = i*3 + j
			}
		}
	}

	if cheapest == 401 {
		return 0
	}

	return cheapest
}
