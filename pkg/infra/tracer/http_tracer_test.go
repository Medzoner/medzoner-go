package tracer_test

import (
	"context"
	"github.com/Medzoner/medzoner-go/pkg/infra/config"
	"github.com/Medzoner/medzoner-go/pkg/infra/tracer"
	"gotest.tools/assert"
	"testing"
)

func TestTracer(t *testing.T) {
	t.Run("Unit: test Tracer success", func(t *testing.T) {
		tr, err := tracer.NewHttpTracer(config.Config{
			TracerFile: "trace.out",
		})
		assert.NilError(t, err)
		assert.Assert(t, tr != nil)
	})
	t.Run("Unit: test Tracer success", func(t *testing.T) {
		tr, err := tracer.NewHttpTracer(config.Config{
			TracerFile: "trace.out",
		})
		err = tr.ShutdownLogger(context.Background())
		err = tr.ShutdownMeter(context.Background())
		err = tr.ShutdownTracer(context.Background())
		if err != nil {
			t.Error(err)
		}
		assert.NilError(t, err)
		assert.Assert(t, tr != nil)
	})
}
