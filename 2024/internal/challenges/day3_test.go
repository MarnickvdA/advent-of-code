package challenges

import "testing"

func TestSolveDayThreePartOne(t *testing.T) {
	data := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	lines := []string{data}

	ans, err := DayThree{}.SolvePartOne(lines)

	if err != nil {
		t.Fatal(err)
	}

	if ans != "161" {
		t.Errorf("Expected 161, got '%s'", ans)
	}
}

func TestSolveDayThreePartTwo(t *testing.T) {
	data := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	lines := []string{data}

	ans, err := DayThree{}.SolvePartTwo(lines)

	if err != nil {
		t.Fatal(err)
	}

	if ans != "48" {
		t.Errorf("Expected 48, got '%s'", ans)
	}
}
