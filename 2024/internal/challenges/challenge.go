package challenges

type Challenge interface {
	SolvePartOne(input []string) (int, error)
	SolvePartTwo(input []string) (int, error)
}

// type DayXXX struct{}
//
// func (DayXXX) SolvePartOne(input []string) (int, error) {
// 	return 0, nil
// }
//
// func (DayXXX) SolvePartTwo(input []string) (int, error) {
// 	return 0, nil
// }
