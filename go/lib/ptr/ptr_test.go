package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/t-ash0410/stack-example/go/lib/ptr"
)

func TestPtr(t *testing.T) {
	t.Parallel()

	t.Run("It returns expected values", func(t *testing.T) {
		t.Parallel()

		var (
			want = "some string"
			got  = ptr.Ptr(want)
		)
		assert.Equal(t, want, *got)
	})
}

func TestValue(t *testing.T) {
	t.Parallel()

	t.Run("It returns expected values", func(t *testing.T) {
		t.Parallel()

		var (
			want = "some string"
			got  = ptr.Value(&want)
		)
		assert.Equal(t, want, got)
	})

	t.Run("It returns zero value", func(t *testing.T) {
		t.Parallel()

		var (
			in   *string
			want = ""
			got  = ptr.Value(in)
		)
		assert.Equal(t, want, got)
	})
}

func TestValueOrDefault(t *testing.T) {
	t.Parallel()

	t.Run("It returns expected values", func(t *testing.T) {
		t.Parallel()

		var (
			want = "some string"
			got  = ptr.ValueOrDefault(&want, "invalid")
		)
		assert.Equal(t, want, got)
	})

	t.Run("It returns fallback", func(t *testing.T) {
		t.Parallel()

		var (
			in   *string
			want = "expected"
			got  = ptr.ValueOrDefault(in, "expected")
		)
		assert.Equal(t, want, got)
	})
}
