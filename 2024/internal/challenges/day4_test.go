package challenges

import (
	"strings"
	"testing"
)

func TestDayFourPartOne(t *testing.T) {
	data := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	input := strings.Split(data, "\n")

	ans, err := DayFour{}.SolvePartOne(input)

	if err != nil {
		t.Fatal(err)
	}

	if ans != 18 {
		t.Errorf("Expected 18, got '%d'", ans)
	}
}
