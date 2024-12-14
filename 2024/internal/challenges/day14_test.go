package challenges

import (
	"strings"
	"testing"
)

func TestDayFourteen(t *testing.T) {
	t.Skip("Only enable when width and height constants are smaller.")

	input := strings.Split(`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`, "\n")

	unit := DayFourteen{}

	tests := []struct {
		id  string
		in  []string
		ans int
		fn  ChallengeFn
	}{
		{
			id:  "1",
			in:  input,
			ans: 12,
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
