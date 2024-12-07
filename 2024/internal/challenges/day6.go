package challenges

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type DaySix struct{}

func getGrid(input []string) ([][]string, int, int) {
	m := make([][]string, len(input))
	x, y := 0, 0

	for fy := range m {
		m[fy] = make([]string, len(input[fy]))

		for fx := range m[fy] {
			if string(input[fy][fx]) == "^" {
				x = fx
				y = fy
			} else {
				m[fy][fx] = string(input[fy][fx])
			}
		}
	}

	return m, x, y
}

func (DaySix) SolvePartOne(input []string) (int, error) {
	m, x, y := getGrid(input)
	runSimulation(m, x, y, len(m))

	count := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] == "X" {
				count++
			}
		}
	}

	printMatrix(0, m)

	return count, nil
}

func (DaySix) SolvePartTwo(input []string) (int, error) {
	m := make([][]string, len(input))

	p1, x, y := getGrid(input)
	runSimulation(p1, x, y, len(m))

	count := 0
	startx, starty := 0, 0

	for i := range m {
		m[i] = make([]string, len(input[i]))

		for j := range input[i] {
			s := string(input[i][j])

			switch s {
			case ".":
				continue
			case "#":
				m[i][j] = "#"
				continue
			}

			starty = i
			startx = j
		}
	}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if p1[y][x] == "X" && !(x == startx && y == starty) {
				m[x][y] = "#"
				if !runSimulation(m, startx, starty, len(m)) {
					count++
				}
			}
		}
	}

	return count, nil
}

func runSimulation(m [][]string, sx, sy, mx int) bool {
	x, y, startx, starty := sx, sy, sx, sy

	steps, dir := 0, 0
	m[y][x] = "X"

	for x >= 0 && y >= 0 && x < mx && y < mx && steps < (mx*mx) {
		if m[y][x] == "#" {
			dir++
			dir %= 4
			x = startx
			y = starty
		} else {
			startx = x
			starty = y
			m[y][x] = "X"
			steps++
		}

		switch dir {
		case 0:
			y--
		case 1:
			x++
		case 2:
			y++
		case 3:
			x--
		}
	}

	return steps != (mx * mx)
}

func checkLoops(m [][]int, a, b, x, y, z int) bool {
	v := map[string]int{}

	directions := map[int]struct{ x, y int }{ // []int{Y,X}
		0:   {y: -1, x: 0},
		90:  {y: 0, x: 1},
		180: {y: 1, x: 0},
		270: {y: 0, x: -1},
	}

	for x >= 0 || y >= 0 {
		v[fmt.Sprintf("(%d,%d)", x, y)] = v[fmt.Sprintf("(%d,%d)", x, y)] + 1

		if v[fmt.Sprintf("(%d,%d)", x, y)] == 3 {
			log.Printf(
				"We are in a loop for custom obstacle at (%d,%d), detected at (x:%d, y:%dy)",
				a,
				b,
				x,
				y,
			)
			return true
		}

		d := directions[z]

		if y+d.y < 0 || y+d.y >= len(m) || x+d.x < 0 || x+d.x >= len(m[y]) {
			break
		}

		for m[y+d.y][x+d.x] == 1 {
			z = (z + 90) % 360
			d = directions[z]
		}

		y += d.y
		x += d.x
	}

	return false
}

func printMatrix(tick int, m [][]string) {

	fmt.Printf("MATRIX ITERATION #%d\n\n", tick)
	for y := range m {
		for x := range m[y] {
			fmt.Print(m[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func printMatrixWithObstacle(m [][]string, dy, dx int) {
	log.Printf("SET OBSTACLE AT y: %d, x: %d TO CREATE LOOP", dy, dx)

	m[dy][dx] = "O"

	for y := range m {
		fmt.Println(strings.Join(m[y], ""))
	}
	fmt.Println()

	time.Sleep(1000 * time.Millisecond)
}
