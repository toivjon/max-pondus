package assert

import (
	"reflect"
	"regexp"
	"testing"
)

// Equal checks whether the given expected and actual values are equal.
func Equal[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

// EqualSlice checks whether the given expected and actual contain same elements.
func EqualSlice[T comparable](t *testing.T, expected []T, actual []T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
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

// Nil checks whether the given value is nil.
func Nil(t *testing.T, value any) {
	t.Helper()
	if value == nil {
		t.Errorf("Expected nil but was %v", value)
	}
}

// Len checks whether the given slice has the given length.
func Len[T any](t *testing.T, slice []T, expected int) {
	t.Helper()
	if len(slice) != expected {
		t.Errorf("Expected slice to have length %d but was %d", expected, len(slice))
	}
}
