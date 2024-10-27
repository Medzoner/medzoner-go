package tracer_test

import (
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
	t.Run("Unit: test Tracer failed", func(t *testing.T) {
		tr, err := tracer.NewHttpTracer(config.Config{
			TracerFile: "/xdf/trace.fail",
		})
		assert.Error(t, err, "failed to create trace output file: open /xdf/trace.fail: no such file or directory")
		assert.Assert(t, tr == nil)
	})
}
