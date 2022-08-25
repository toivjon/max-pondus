package sort

import (
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
)

func TestReverseWithNil(t *testing.T) {
	slice := ([]string)(nil)
	Reverse(slice)
	assert.Nil(t, slice)
}

func TestReverseWithEmptySlice(t *testing.T) {
	slice := []string{}
	Reverse(slice)
	assert.Len(t, slice, 0)
}

func TestReverseWithNonEmptySlice(t *testing.T) {
	slice := []string{"foo", "bar", "foobar"}
	Reverse(slice)
	assert.EqualSlice(t, []string{"foobar", "bar", "foo"}, slice)
}
