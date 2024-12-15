package utils

import (
	"testing"
)

func TestCase(t *testing.T, fn func([]string) (int, error), input []string, ans int) {
	t.Helper()

	a, err := fn(input)

	if err != nil {
		t.Error(err)
	}

	if ans != a {
		t.Errorf("Expected %d, got %d", ans, a)
	}
}
