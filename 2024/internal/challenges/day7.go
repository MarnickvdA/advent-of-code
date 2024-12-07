package challenges

import (
	"fmt"
	"strconv"
	"strings"
)

type DaySeven struct{}

type Operation func(a, b int) int

func (DaySeven) SolvePartOne(input []string) (int, error) {
	operations := map[string]Operation{
		"+": func(a, b int) int {
			return a + b
		},
		"*": func(a, b int) int {
			return a * b
		},
	}

	sum := solveWithOperations(input, operations)

	return sum, nil
}

func (DaySeven) SolvePartTwo(input []string) (int, error) {
	operations := map[string]Operation{
		"+": func(a, b int) int {
			return a + b
		},
		"*": func(a, b int) int {
			return a * b
		},
		"||": func(a, b int) int {
			concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
			return concat
		},
	}

	sum := solveWithOperations(input, operations)

	return sum, nil
}

func solveWithOperations(input []string, operations map[string]Operation) int {
	total := 0

	for _, line := range input {
		partials := strings.Split(line, ": ")
		sum, _ := strconv.Atoi(partials[0])

		parts := []int{}
		for _, s := range strings.Split(partials[1], " ") {
			val, _ := strconv.Atoi(s)
			parts = append(parts, val)
		}

		for k := range operations {
			ans := findSum(sum, 0, 0, k, &parts, &operations)
			if ans == sum {
				total += ans
				break
			}
		}
	}

	return total
}

func findSum(
	target, total, i int,
	op string,
	parts *[]int,
	operations *map[string]Operation,
) int {
	if i == len(*parts) {
		if total == target {
			return total
		}
		return 0
	}

	if total > target {
		return 0
	}

	total = (*operations)[op](total, (*parts)[i])

	for k := range *operations {
		ans := findSum(target, total, i+1, k, parts, operations)
		if ans == target {
			return ans
		}
	}

	return 0
}
