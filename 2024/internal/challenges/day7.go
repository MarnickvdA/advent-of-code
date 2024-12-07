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
	solutions := map[int][][]string{}

	for _, line := range input {
		partials := strings.Split(line, ": ")
		sum, _ := strconv.Atoi(partials[0])
		parts := []int{}
		for _, s := range strings.Split(partials[1], " ") {
			val, _ := strconv.Atoi(s)
			parts = append(parts, val)
		}

		solutions[sum] = [][]string{}

		for k := range operations {
			findSum(sum, 0, 0, k, parts, []string{}, &solutions, operations)
		}
	}

	sum := 0
	for k, v := range solutions {
		if len(v) > 0 {
			sum += k
		}
	}

	return sum
}

func findSum(
	target, total, i int,
	op string,
	parts []int,
	ops []string,
	sols *map[int][][]string,
	operations map[string]Operation,
) {
	if i == len(parts) {
		if total == target {
			(*sols)[target] = append((*sols)[target], ops)
		}
		return
	}

	if total > target {
		return
	}

	total = operations[op](total, parts[i])
	ops = append(ops, op)

	for k := range operations {
		findSum(target, total, i+1, k, parts, ops, sols, operations)
	}
}
