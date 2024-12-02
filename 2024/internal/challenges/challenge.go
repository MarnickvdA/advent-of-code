package challenges

type Challenge interface {
	SolvePartOne(input []string) (string, error)
	SolvePartTwo(input []string) (string, error)
}
