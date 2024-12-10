package challenges

import (
	"strings"
	"testing"
)

func TestDayTenPartOne(t *testing.T) {
	input := strings.Split(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n")

	ans, _ := DayTen{}.SolvePartOne(input)

	if ans != 36 {
		t.Errorf("Expected 36, got %d", ans)
	}
}

func TestDayTenPartTwo(t *testing.T) {
	input := strings.Split(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n")

	ans, _ := DayTen{}.SolvePartTwo(input)

	if ans != 81 {
		t.Errorf("Expected 81, got %d", ans)
	}
}
