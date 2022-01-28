package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapIntsToFloatDoubled(t *testing.T) {
	assert.Equal(
		t,
		Map(
			[]int{1, 2, 3},
			func(v int) float64 {
				return 2 * float64(v)
			},
		),
		[]float64{2, 4, 6},
	)
}
