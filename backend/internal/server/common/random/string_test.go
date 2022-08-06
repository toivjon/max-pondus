package random

import (
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
)

func TestStringPanicsWithNegativeLength(t *testing.T) {
	defer func() { _ = recover() }()
	String(-1)
	t.Errorf("Expected a panic but was not issued")
}

func TestStringReturnsEmptyStringWithZeroLength(t *testing.T) {
	s := String(0)
	if len(s) != 0 {
		t.Errorf("Expected an empty string but was %q", s)
	}
}

func TestStringReturnsDiffentStringsOnConsecutiveCalls(t *testing.T) {
	a := String(10)
	b := String(10)
	assert.NotEqual(t, a, b)
}
