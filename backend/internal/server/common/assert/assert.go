package assert

import (
	"regexp"
	"testing"
)

// Equal checks whether the given expected and actual values are equal.
func Equal(t *testing.T, expected any, actual any) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

// NotEqual checks whether the given values are not equal.
func NotEqual(t *testing.T, lhs any, rhs any) {
	t.Helper()
	if lhs == rhs {
		t.Errorf("Expected different values but were %v and %v", lhs, rhs)
	}
}

// Match checks whether the given string matches with the provided pattern.
func Match(t *testing.T, pattern, value string) {
	t.Helper()
	result, err := regexp.MatchString(pattern, value)
	if err != nil {
		panic(err)
	}
	if !result {
		t.Errorf("Expected pattern %q to match %q", pattern, value)
	}
}
