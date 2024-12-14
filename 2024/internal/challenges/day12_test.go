package challenges

import (
	"strings"
	"testing"
)

func TestDayTwelve(t *testing.T) {
	inputA := strings.Split(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`, "\n")

	inputB := strings.Split(`AAAA
BBCD
BBCC
EEEC`, "\n")

	inputC := strings.Split(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`, "\n")

	inputD := strings.Split(`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`, "\n")

	inputE := strings.Split(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`, "\n")

	unit := DayTwelve{}

	tests := []struct {
		id  string
		in  []string
		ans int
		fn  ChallengeFn
	}{
		{
			id:  "1.A",
			in:  inputA,
			ans: 1930,
			fn:  unit.SolvePartOne,
		},
		{
			id:  "1.B",
			in:  inputB,
			ans: 140,
			fn:  unit.SolvePartOne,
		},
		{
			id:  "1.C",
			in:  inputC,
			ans: 772,
			fn:  unit.SolvePartOne,
		},
		{
			id:  "2.A",
			in:  inputA,
			ans: 1206,
			fn:  unit.SolvePartTwo,
		},
		{
			id:  "2.B",
			in:  inputB,
			ans: 80,
			fn:  unit.SolvePartTwo,
		},
		{
			id:  "2.C",
			in:  inputC,
			ans: 436,
			fn:  unit.SolvePartTwo,
		},
		{
			id:  "2.D",
			in:  inputD,
			ans: 236,
			fn:  unit.SolvePartTwo,
		},
		{
			id:  "2.E",
			in:  inputE,
			ans: 368,
			fn:  unit.SolvePartTwo,
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
