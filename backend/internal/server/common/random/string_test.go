package random_test

import (
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/random"
)

func TestStringPanicsWithNegativeLength(t *testing.T) {
	t.Parallel()
	defer func() { _ = recover() }()
	random.String(-1)
	t.Errorf("Expected a panic but was not issued")
}

func TestStringReturnsEmptyStringWithZeroLength(t *testing.T) {
	t.Parallel()
	s := random.String(0)
	if len(s) != 0 {
		t.Errorf("Expected an empty string but was %q", s)
	}
}

func TestStringReturnsDiffentStringsOnConsecutiveCalls(t *testing.T) {
	t.Parallel()
	a := random.String(10)
	b := random.String(10)
	assert.NotEqual(t, a, b)
}
