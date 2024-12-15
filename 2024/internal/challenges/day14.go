package challenges

import (
	"fmt"
	"regexp"
	"strconv"
)

type DayFourteen struct{}

type Robot struct {
	x, y   int
	dx, dy int
}

const WIDTH = 101
const HEIGHT = 103

func (DayFourteen) SolvePartOne(input []string) (int, error) {
	tiles := make([][]int, HEIGHT)
	for x := range tiles {
		tiles[x] = make([]int, WIDTH)
	}

	robots := parseRobots(&input)

	for ticks := range 10000 {
		for _, r := range robots {
			r.runSimulation(1)
			tiles[r.y][r.x] += 1
		}

		displayTiles(ticks, &tiles)

		for _, r := range robots {
			tiles[r.y][r.x] -= 1
		}
	}

	m1, m2, m3, m4 := 0, 0, 0, 0

	for y := range tiles {
		for x := range tiles[y] {
			if y < HEIGHT/2 {
				if x < WIDTH/2 {
					m1 += tiles[y][x]
				} else if x > WIDTH/2 {
					m2 += tiles[y][x]
				}
			} else if y > HEIGHT/2 {
				if x < WIDTH/2 {
					m3 += tiles[y][x]
				} else if x > WIDTH/2 {
					m4 += tiles[y][x]
				}
			}

		}
	}

	return m1 * m2 * m3 * m4, nil
}

func (DayFourteen) SolvePartTwo(input []string) (int, error) {
	return 0, nil
}

func displayTiles(tick int, tiles *[][]int) {
	print := ``
	potential := false
	for y := range *tiles {
		c := 0
		for x := range (*tiles)[y] {
			if (*tiles)[y][x] > 0 {
				c++
				print += "#"
			} else {
				print += "."
			}
		}
		print += "\n"

		if c > 15 {
			potential = true
		}
	}

	if potential {
		fmt.Printf("TILES FOR TICK: %d\n", tick)
		fmt.Println(print)
	}
}

func parseRobots(input *[]string) []*Robot {
	robots := []*Robot{}

	r := regexp.MustCompile(`-?([0-9])+`)

	for _, line := range *input {
		s := r.FindAllString(line, 4)

		if len(s) != 4 {
			continue
		}

		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		dx, _ := strconv.Atoi(s[2])
		dy, _ := strconv.Atoi(s[3])

		robot := Robot{
			x: x, y: y, dx: dx, dy: dy,
		}

		robots = append(robots, &robot)
	}

	return robots
}

func (r *Robot) runSimulation(ticks int) {
	for range ticks {
		r.x = ((r.x+r.dx)%WIDTH + WIDTH) % WIDTH
		r.y = ((r.y+r.dy)%HEIGHT + HEIGHT) % HEIGHT
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}
