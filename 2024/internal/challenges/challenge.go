package challenges

type Challenge interface {
	SolvePartOne(input []string) (int, error)
	SolvePartTwo(input []string) (int, error)
}

// type DayXXX struct{}
//
// func (DayXXX) SolvePartOne(input []string) (string, error) {
// 	return "", nil
// }
//
// func (DayXXX) SolvePartTwo(input []string) (string, error) {
// 	return "", nil
// }
