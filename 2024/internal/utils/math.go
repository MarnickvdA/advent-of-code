package utils

func Mod(a, b int) int {
	return (a%b + b) % b
}