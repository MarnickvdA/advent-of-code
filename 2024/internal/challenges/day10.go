package challenges

import (
	"fmt"
	"strconv"
)

type DayTen struct{}

type Map [][]int
type Point [2]int
type Bounds [2]int

func (p *Point) validPoint(b *Bounds) bool {
	return p[0] >= 0 && p[1] >= 0 && p[0] < b[0] && p[1] < b[1]
}

func (p *Point) toString() string {
	return fmt.Sprintf("(%d,%d)", p[0], p[1])
}

func (m *Map) at(p Point) int {
	return (*m)[p[0]][p[1]]
}

func parseTopographicMap(input *[]string, m *Map, s *[]Point) {
	*m = make(Map, len(*input))
	*s = []Point{}

	for y, line := range *input {
		for x, c := range line {
			v, _ := strconv.Atoi(string(c))
			if (*m)[y] == nil {
				(*m)[y] = make([]int, len(line))
			}

			(*m)[y][x] = v

			if v == 0 {
				*s = append(*s, Point{y, x})
			}
		}
	}
}

var dirs = [][]int{
	{-1, 0}, {0, 1}, {1, 0}, {0, -1},
}

func getRouteCount(m *Map, p Point, prevAlt int, ends *map[string]int, b Bounds) {
	if !p.validPoint(&b) || m.at(p) != prevAlt+1 {
		return
	}

	if m.at(p) == 9 {
		(*ends)[p.toString()] += 1
		return
	}

	for _, dir := range dirs {
		getRouteCount(m, Point{p[0] + dir[0], p[1] + dir[1]}, prevAlt+1, ends, b)
	}
}

func (DayTen) SolvePartOne(input []string) (int, error) {
	var m Map
	var s []Point

	parseTopographicMap(&input, &m, &s)

	sum := 0
	for _, sp := range s {
		ends := map[string]int{}
		getRouteCount(&m, sp, -1, &ends, Bounds{len(m), len(m[0])})
		sum += len(ends)
	}

	return sum, nil
}

func (DayTen) SolvePartTwo(input []string) (int, error) {
	var m Map
	var s []Point

	parseTopographicMap(&input, &m, &s)

	sum := 0
	for _, sp := range s {
		ends := map[string]int{}
		getRouteCount(&m, sp, -1, &ends, Bounds{len(m), len(m[0])})

		for _, v := range ends {
			sum += v
		}
	}

	return sum, nil
}
