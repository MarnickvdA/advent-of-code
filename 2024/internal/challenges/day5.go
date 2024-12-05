package challenges

import (
	"strconv"
	"strings"

	"aoc.abitofsoftware.com/internal/sorting"
)

type DayFive struct{}

func getAdjacencyMatrix(input []string) [][]int {
	adj := make([][]int, 100)

	for _, line := range input {
		pairs := strings.Split(line, "|")
		if len(pairs) == 2 {
			x, _ := strconv.Atoi(pairs[0])
			y, _ := strconv.Atoi(pairs[1])

			if adj[x] == nil {
				adj[x] = []int{}
			}

			adj[x] = append(adj[x], y)
			continue
		}
	}

	return adj
}

func (DayFive) SolvePartOne(input []string) (int, error) {
	sum := 0
	var adj [][]int

	for i, line := range input {
		if adj == nil && line == "" {
			adj = getAdjacencyMatrix(input[:i])
		}

		order := strings.Split(line, ",")
		if len(order) > 1 {
			visited := map[int]bool{}
			addToSum := 0

		orderLoop:
			for i, val := range order {
				v, _ := strconv.Atoi(val)
				visited[v] = true

				if len(order)/2 == i {
					addToSum = v
				}

				for _, a := range adj[v] {
					if visited[a] {
						addToSum = 0
						break orderLoop
					}
				}
			}

			sum += addToSum
		}
	}

	return sum, nil
}

func (DayFive) SolvePartTwo(input []string) (int, error) {
	sum := 0
	var adj [][]int

	for i, line := range input {
		if adj == nil && line == "" {
			adj = getAdjacencyMatrix(input[:i])
		}

		order := strings.Split(line, ",")
		if len(order) > 1 {
			visited := map[int]bool{}
			addToSum := 0

		orderLoop:
			for _, val := range order {
				v, _ := strconv.Atoi(val)
				visited[v] = true

				for _, a := range adj[v] {
					if visited[a] {
						io := make([]int, len(order))
						for a, o := range order {
							v, _ := strconv.Atoi(o)
							io[a] = v
						}

						graph := map[int][]int{}
						exists := map[int]bool{}

						for _, o := range io {
							exists[o] = true
						}

						for _, o := range io {
							for _, a := range adj[o] {
								if exists[a] {
									if graph[o] == nil {
										graph[o] = []int{}
									}

									graph[o] = append(graph[o], a)
								}
							}
						}

						sorted := sorting.TopologicalSort(graph)
						addToSum = sorted[len(sorted)/2]

						break orderLoop
					}
				}
			}

			sum += addToSum
		}
	}

	return sum, nil
}
