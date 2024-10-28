package tracer_test

import (
	"context"
	"testing"

	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"

	"gotest.tools/assert"
)

func TestTracer(t *testing.T) {
	t.Run("Unit: test Tracer success", func(t *testing.T) {
		tr, err := tracer.NewHttpTracer(config.Config{})
		assert.NilError(t, err)
		assert.Assert(t, tr != nil)
	})
	t.Run("Unit: test Tracer success", func(t *testing.T) {
		tr, err := tracer.NewHttpTracer(config.Config{})
		assert.NilError(t, err)
		_ = tr.ShutdownLogger(context.Background())
		_ = tr.ShutdownMeter(context.Background())
		_ = tr.ShutdownTracer(context.Background())
		assert.Assert(t, tr != nil)
	})
}
