package challenges

import (
	"fmt"
	"strings"
)

type DayEight struct{}

type Position struct {
	x int
	y int
}

func (p *Position) difference(with *Position) (int, int) {
	return p.x - with.x, p.y - with.y
}

func (p *Position) isValidForOne(key string, m [][]string) bool {
	return p.inBounds(m) && m[p.y][p.x] != key && m[p.y][p.x] != "#"
}

func (p *Position) inBounds(m [][]string) bool {
	return p.x >= 0 && p.x < len(m[0]) && p.y >= 0 && p.y < len(m)
}

func (p *Position) label(key string) string {
	return fmt.Sprintf("%s_(%d,%d)", key, p.x, p.y)
}

func getDataFromInput(input []string) ([][]string, map[string][]Position) {
	m := [][]string{}
	pos := map[string][]Position{}

	for i, l := range input {
		e := strings.Split(l, "")
		m = append(m, e)
		for j, v := range m[i] {
			if v != "." {
				if pos[v] == nil {
					pos[v] = []Position{}
				}
				pos[v] = append(pos[v], Position{x: j, y: i})
			}
		}
	}

	return m, pos
}

func (DayEight) SolvePartOne(input []string) (int, error) {
	m, pos := getDataFromInput(input)

	antinodes := 0
	for k, v := range pos {
		for i := range v {
			for j := i + 1; j < len(v); j++ {
				dx, dy := v[i].difference(&v[j])

				pi := Position{x: v[i].x + dx, y: v[i].y + dy}
				if pi.isValidForOne(k, m) {
					m[pi.y][pi.x] = "#"
					antinodes++
				}

				pj := Position{x: v[j].x + (-1 * dx), y: v[j].y + (-1 * dy)}
				if pj.isValidForOne(k, m) {
					m[pj.y][pj.x] = "#"
					antinodes++
				}
			}
		}
	}

	return antinodes, nil
}

func (DayEight) SolvePartTwo(input []string) (int, error) {
	m, pos := getDataFromInput(input)

	for _, v := range pos {
		for i, vi := range v {
			for j, vj := range v {
				if i == j {
					continue
				}

				m[vi.y][vi.x] = "#"

				dx, dy := vi.difference(&vj)
				pi := Position{x: vi.x + dx, y: vi.y + dy}

				for pi.inBounds(m) {
					m[pi.y][pi.x] = "#"

					pi.x = pi.x + dx
					pi.y = pi.y + dy
				}
			}
		}
	}

	antinodes := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] == "#" {
				antinodes++
			}
		}
	}

	return antinodes, nil
}
