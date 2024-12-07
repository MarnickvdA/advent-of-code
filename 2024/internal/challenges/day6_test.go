package challenges

import (
	"strings"
	"testing"
)

func TestDaySixPartTwo(t *testing.T) {
	t.Skip()

	input := strings.Split(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`, "\n")

	ans, _ := DaySix{}.SolvePartTwo(input)

	if ans != 6 {
		t.Errorf("Expected 6, got %d", ans)
	}
}
