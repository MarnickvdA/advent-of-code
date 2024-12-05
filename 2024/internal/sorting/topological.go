package sorting

func TopologicalSort(graph map[int][]int) []int {
	visited := make([]bool, len(graph))
	var sorted []int
	var visitAll func([]int)
	var stack []int

	visitAll = func(nodes []int) {
		for _, n := range nodes {
			if !visited[n] {
				visited[n] = true
				visitAll(graph[n])
				stack = append(stack, n)
			}
		}
	}

	for n := range graph {
		if n != 0 {
			visitAll([]int{n})
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		sorted = append(sorted, stack[i])
	}

	return sorted
}
