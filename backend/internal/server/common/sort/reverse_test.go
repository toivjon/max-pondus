package sort_test

import (
	"testing"

	"github.com/toivjon/max-pondus/backend/internal/server/common/assert"
	"github.com/toivjon/max-pondus/backend/internal/server/common/sort"
)

func TestReverseWithNil(t *testing.T) {
	t.Parallel()
	slice := ([]string)(nil)
	sort.Reverse(slice)
	assert.Nil(t, slice)
}

func TestReverseWithEmptySlice(t *testing.T) {
	t.Parallel()
	slice := []string{}
	sort.Reverse(slice)
	assert.Len(t, slice, 0)
}

func TestReverseWithNonEmptySlice(t *testing.T) {
	t.Parallel()
	slice := []string{"foo", "bar", "foobar"}
	sort.Reverse(slice)
	assert.EqualSlice(t, []string{"foobar", "bar", "foo"}, slice)
}
