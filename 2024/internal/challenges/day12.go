package challenges

import (
	"slices"
	"strings"
)

type DayTwelve struct{}

type Garden map[int]map[Position]bool

func parseInput(input *[]string) (Garden, map[int]int) {
	g := Garden{}
	cs := map[int]int{}

	id := 0
	cur := ""
	neighbors := []Position{{x: 0, y: 0}}
	visited := map[Position]bool{}

	inBounds := func(x, y int) bool {
		return y >= 0 && y < len(*input) && x >= 0 && x < len((*input)[0])
	}

	var checkSurrounding func(int, int)
	checkSurrounding = func(x, y int) {
		pos := Position{x: x, y: y}

		if visited[pos] { // we don't visited positions multiple times.
			return
		}

		if g[id] == nil {
			g[id] = map[Position]bool{}
		}

		g[id][pos] = true
		visited[pos] = true

		// t, l, d, r
		dirs := [][]int{
			{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1},
		}

		mdirs := make([]int, 8)
		for i, d := range dirs {
			if !inBounds(x+d[1], y+d[0]) {
				continue
			}

			if string(strings.Split((*input)[y+d[0]], "")[x+d[1]]) == cur {
				mdirs[i] = 1

				if i%2 == 0 {
					checkSurrounding(x+d[1], y+d[0])
				}
			} else if i%2 == 0 {
				neighbors = append(neighbors, Position{x: x + d[1], y: y + d[0]})
			}
		}

		corners := 0
		for i := 0; i < 8; i += 2 {
			if mdirs[i%8] == 0 && mdirs[(i+1)%8] == 0 && mdirs[(i+2)%8] == 0 {
				corners++
			}

			if mdirs[i%8] == 1 && mdirs[(i+1)%8] == 0 && mdirs[(i+2)%8] == 1 {
				corners++
			}

			if mdirs[i%8] == 0 && mdirs[(i+1)%8] == 1 && mdirs[(i+2)%8] == 0 {
				corners++
			}
		}

		cs[id] += corners
	}

	for len(neighbors) > 0 {
		pos := neighbors[0:1][0]
		neighbors = slices.Delete(neighbors, 0, 1)

		if !visited[pos] {
			cur = string((*input)[pos.y][pos.x : pos.x+1])

			checkSurrounding(pos.x, pos.y)

			id++
		}
	}

	return g, cs
}

func (DayTwelve) SolvePartOne(input []string) (int, error) {
	garden, _ := parseInput(&input)
	cost := 0

	for _, v := range garden {
		area := len(v)

		perimeter := 0

		for a := range v {
			for _, d := range dirs {
				if !v[Position{y: a.y + d[0], x: a.x + d[1]}] {
					perimeter++
				}
			}
		}

		cost += area * perimeter
	}

	return cost, nil
}

func (DayTwelve) SolvePartTwo(input []string) (int, error) {
	garden, corners := parseInput(&input)
	cost := 0

	for k, v := range garden {
		area := len(v)

		perimeter := corners[k]

		cost += area * perimeter
	}

	return cost, nil
}
