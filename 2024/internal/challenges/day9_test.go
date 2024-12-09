package challenges

import "testing"

func TestDayNinePartOne(t *testing.T) {
	line := `2333133121414131402`

	ans, _ := DayNine{}.SolvePartOne([]string{line})

	if ans != 1928 {
		t.Errorf("Expected 1928, got %d", ans)
	}
}

func TestDayNinePartTwo(t *testing.T) {
	line := `2333133121414131402`

	ans, _ := DayNine{}.SolvePartTwo([]string{line})

	if ans != 2858 {
		t.Errorf("Expected 2858, got %d", ans)
	}
}
