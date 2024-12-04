package challenges

import (
	"strings"
	"testing"
)

func TestDayOnePartOne(t *testing.T) {
	data := `3   4
4   3
2   5
1   3
3   9
3   3
  `

	input := strings.Split(data, "\n")

	ans, err := DayOne{}.SolvePartOne(input)

	if err != nil {
		t.Fatal(err)
	}

	if ans != 11 {
		t.Errorf("Expected 11, got %d", ans)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	data := `3   4
4   3
2   5
1   3
3   9
3   3
  `

	input := strings.Split(data, "\n")

	ans, err := DayOne{}.SolvePartTwo(input)

	if err != nil {
		t.Error(err)
	}

	if ans != 31 {
		t.Errorf("Expected 31, got %d", ans)
	}
}
