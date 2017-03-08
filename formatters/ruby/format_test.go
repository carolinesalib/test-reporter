package ruby

import (
	"testing"

	"github.com/codeclimate/test-reporter/env"
	"github.com/stretchr/testify/require"
)

func Test_Format(t *testing.T) {
	r := require.New(t)

	rb := New("./ruby-example.json")
	err := rb.Parse()
	r.NoError(err)

	rep, err := rb.Format()
	r.NoError(err)

	r.InDelta(97.95, rep.CoveredPercent, 1)

	env, err := env.New()
	r.NoError(err)
	r.Equal(env, rep.CIService)
	r.Len(rep.SourceFiles, len(rb.Tests[0].SourceFiles))

	lc := rep.LineCounts
	r.Equal(lc.Covered, 56)
	r.Equal(lc.Missed, 1)
	r.Equal(lc.Total, 57)
}
