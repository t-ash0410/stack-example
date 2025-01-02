package testctx_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/t-ash0410/stack-example/go/lib/testctx"
)

func TestCanceledContext(t *testing.T) {
	ctx := testctx.CanceledContext()
	assert.EqualError(t, ctx.Err(), context.Canceled.Error())
}
