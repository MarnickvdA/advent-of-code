package challenges

import (
	"strings"
	"testing"
)

func TestDayFivePartOne(t *testing.T) {
	input := strings.Split(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`, "\n")

	ans, err := DayFive{}.SolvePartOne(input)

	if err != nil {
		t.Fatal(err)
	}

	if ans != 143 {
		t.Errorf("Expected 143, got %d", ans)
	}
}

func TestDayFivePartTwo(t *testing.T) {
	input := strings.Split(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`, "\n")

	ans, err := DayFive{}.SolvePartTwo(input)

	if err != nil {
		t.Fatal(err)
	}

	if ans != 123 {
		t.Errorf("Expected 123, got %d", ans)
	}
}
