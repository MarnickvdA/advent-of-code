package challenges

import (
	"strings"
	"testing"
)

func TestDayThirteen(t *testing.T) {
	inputA := strings.Split(`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`, "\n")

	unit := DayThirteen{}

	tests := []struct {
		id  string
		in  []string
		ans int
		fn  ChallengeFn
	}{
		{
			id:  "1.A",
			in:  inputA,
			ans: 480,
			fn:  unit.SolvePartOne,
		},
	}

	for _, test := range tests {
		ans, err := test.fn(test.in)

		if err != nil {
			t.Error(err)
			continue
		}

		if ans != test.ans {
			t.Errorf("%s | Expected %d, got %d", test.id, test.ans, ans)
		}
	}
}
