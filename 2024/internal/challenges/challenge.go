package challenges

type Challenge interface {
	SolvePartOne(input []string) (string, error)
	SolvePartTwo(input []string) (string, error)
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
