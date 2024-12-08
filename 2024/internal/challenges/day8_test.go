package challenges

import (
	"strings"
	"testing"
)

func TestDayEightPartOne(t *testing.T) {
	input := strings.Split(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`, "\n")

	ans, _ := DayEight{}.SolvePartOne(input)

	if ans != 14 {
		t.Errorf("Expected 14, got %d", ans)
	}
}

func TestDayEightPartTwo(t *testing.T) {
	input := strings.Split(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`, "\n")

	ans, _ := DayEight{}.SolvePartTwo(input)

	if ans != 34 {
		t.Errorf("Expected 34, got %d", ans)
	}

	//	x := `
	//
	// ##....#....# 4
	// .#.#....#... 3
	// ..#.##....#. 4
	// ..##...#.... 3
	// ....#....#.. 2
	// .#...##....# 4
	// ...#..#..... 2
	// #....#.#.... 3
	// ..#.....#... 2
	// ....#....#.. 2
	// .#........#. 2
	// ...#......## 3
	//
	//	`
}
