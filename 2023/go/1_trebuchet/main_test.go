package main

import "testing"

func TestStringNumMatcher(t *testing.T) {
	if !stringNumMatcher("abcone5", "one", 3) {
		t.Fail()
	}
	if !stringNumMatcher("abctwo5", "two", 3) {
		t.Fail()
	}
	if !stringNumMatcher("abcthree3", "three", 3) {
		t.Fail()
	}
	if !stringNumMatcher("three3", "three", 0) {
		t.Fail()
	}
	// Should fail.
	if stringNumMatcher("abcone5", "one", 2) {
		t.Fail()
	}
}

func TestGetSum(t *testing.T) {
	lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	if getSum(lines) != 281 {
		t.Fail()
	}
}
