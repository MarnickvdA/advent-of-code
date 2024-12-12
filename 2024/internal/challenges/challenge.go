package challenges

type ChallengeFn func([]string) (int, error)

type Challenge interface {
	SolvePartOne(input []string) (int, error)
	SolvePartTwo(input []string) (int, error)
}

// type DayX struct{}
//
// func (DayX) SolvePartOne(input []string) (int, error) {
// 	return 0, nil
// }
//
// func (DayX) SolvePartTwo(input []string) (int, error) {
// 	return 0, nil
// }
