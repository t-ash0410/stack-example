package ctxtest_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/t-ash0410/stack-example/go/lib/ctxtest"
)

func TestCanceledContext(t *testing.T) {
	ctx := ctxtest.CanceledContext()
	assert.EqualError(t, ctx.Err(), context.Canceled.Error())
}
