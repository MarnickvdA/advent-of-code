package challenges

import (
	"strings"
	"testing"
)

func TestDayTwoPartOne(t *testing.T) {
	data := `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
  `

	lines := strings.Split(data, "\n")

	ans, err := DayTwo{}.SolvePartOne(lines)

	if err != nil {
		t.Error(err)
	}

	if ans != "2" {
		t.Errorf("Expected 2, got '%s'", ans)
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	data := `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

	lines := strings.Split(data, "\n")

	ans, err := DayTwo{}.SolvePartTwo(lines)

	if err != nil {
		t.Error(err)
	}

	if ans != "4" {
		t.Errorf("Expected 4, got '%s'", ans)
	}
}

func TestDayTwoPartTwoExtended(t *testing.T) {
	data := `
48 46 47 49 51 54 56
1 1 2 3 4 5
1 2 3 4 5 5
5 1 2 3 4 5
1 4 3 2 1
1 6 7 8 9
1 2 3 4 3
9 8 7 6 7`

	lines := strings.Split(data, "\n")

	ans, err := DayTwo{}.SolvePartTwo(lines)

	if err != nil {
		t.Error(err)
	}

	if ans != "8" {
		t.Errorf("Expected 8, got '%s'", ans)
	}
}
